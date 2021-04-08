package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	_ "strings"
)

type BenefitUsageDetail struct {
	IncurredAmount float64
	ApprovedAmount float64
	ClaimedAmount  float64
	BenefitId      string
	BenefitUsageId uint64
}

type BalancedBenefitUsageDetailAggregate struct {
	BenefitUsageDetails []BenefitUsageDetail
	SumIncurred         float64
	SumApproved         float64
}

func (bud *BenefitUsageDetail) String() string {
	return fmt.Sprintf("{\"incurred_amount\": %v, \"approved_amount\": %v, \"claimed_amount\": %v, \"benefit_id\": %v, \"benefit_usage_id\": %v}", bud.IncurredAmount, bud.ApprovedAmount, bud.ClaimedAmount, bud.BenefitId, bud.BenefitUsageId)
}

func balanceApprovedAmounts(benefitUsageId uint64, db *sql.DB) {
	var benefitUsageDetails []BenefitUsageDetail
	results, err := db.Query(fmt.Sprintf("SELECT incurred_amount, approved_amount, claimed_amount, benefit_id, benefit_usage_id FROM benefit_usage_details WHERE benefit_usage_id=%v", benefitUsageId))
	if err != nil {
		log.Fatalf("An error %v occurred while querying the database", err)
	}

	budGrp := make(map[string]BalancedBenefitUsageDetailAggregate)

	for results.Next() {
		var bud BenefitUsageDetail
		err = results.Scan(&bud.IncurredAmount, &bud.ApprovedAmount, &bud.ClaimedAmount, &bud.BenefitId, &bud.BenefitUsageId, )
		if err != nil {
			log.Fatalf("An error %v occurred while converting db row to struct", err)
		}
		if eAggregate, ok := budGrp[bud.BenefitId]; !ok {
			var sl []BenefitUsageDetail
			sl = append(sl, bud)
			sumIncurred := bud.IncurredAmount
			sumApproved := bud.ApprovedAmount

			budGrp[bud.BenefitId] = BalancedBenefitUsageDetailAggregate{BenefitUsageDetails: sl, SumApproved: sumApproved, SumIncurred: sumIncurred}
		} else {
			eAggregate.BenefitUsageDetails = append(eAggregate.BenefitUsageDetails, bud)
			eAggregate.SumIncurred += bud.IncurredAmount
			eAggregate.SumApproved += bud.ApprovedAmount
		}
		benefitUsageDetails = append(benefitUsageDetails, bud)
	}

	for benefitId, buds := range budGrp {
		fmt.Printf("%v: %v\n", benefitId, buds)
		for _, benefitUsageDetail := range buds.BenefitUsageDetails {
			uAppr := benefitUsageDetail.IncurredAmount * buds.SumIncurred / buds.SumApproved
			fmt.Println(fmt.Sprintf("updated incurred amount %v  updated approved amount %v", benefitUsageDetail.IncurredAmount, uAppr))
		}
	}

}

func main() {
	benefitUsageIds := []uint64{548323,
		495413,
		496474,
		539543,
		496437,
		493270,
		494926,
		494694,
		492535,
		494618,
		496063,
		493380,
		495990,
		493520,
		494686,
		494548,
		496710,
		536394,
		533023,
		493629,
		546018,
		495991,
		494516,
		539076,
		532663,
		494697,
		548206,
		549156,
		496148,
		495419,
		538284,
		496915,
		493310,
		493958,}
	db, err := sql.Open("mysql", "member4bi:member4bi@tcp(membership-replica.clsqkkbd9zef.ap-southeast-1.rds.amazonaws.com:3306)/benefit_service")
	defer db.Close()
	if err != nil {
		log.Fatalf("An error %v occurred while opening a connection to db", err)
	}

	for _, benefitUsageId := range benefitUsageIds {
		balanceApprovedAmounts(benefitUsageId, db)
	}
}

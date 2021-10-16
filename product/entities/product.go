package entities

type ProductRider struct {
    Id					string	`json:"id"`
    Slug				string	`json:"slug"`
    Name				string	`json:"name"`
    IsActive			bool	`json:"is_active"`
    ParentId			string	`json:"parent_id"`
    Summary				string	`json:"summary"`
    Description			string	`json:"description"`
    IconSvg				string	`json:"icon_svg"`
    CoveragePeriod		string	`json:"coverage_period"`
}

type ProductBenefitGroup struct {
    Id					string	`json:"id"`
    TooltipText			string	`json:"tooltip_text"`
    Name				string	`json:"name"`
    Order				int	`json:"order"`
	ProductBenefit		[]ProductBenefit `json:"benefits"`
}

type ProductBenefit struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
    IconSvg				string	`json:"icon_svg"`
    IconEtc				string	`json:"icon_etc"`
    ProductId			string	`json:"product_id"`
    ProductBenefitGroupId	string	`json:"product_benefit_group_id"`
    TooltipText			string	`json:"tooltip_text"`
    TooltipTextDescription	string	`json:"tooltip_text_description"`
    Order				int	`json:"order"`
}

type ProductCategory struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
}

type ProductInsuranceType struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
}
type Product struct {
    Id						string	`json:"id"`
    Slug					string	`json:"slug"`
    Name					string	`json:"name"`
    DocName					string	`json:"doc_name"`
    ParentId				string	`json:"parent_id"`
    IsActive		 		bool	`json:"is_active"`
    Featured		 		bool	`json:"featured"`
    BundlingWithRider		bool	`json:"bundling_with_rider"`
    Subheading				string	`json:"subheading"`
    Summary					string	`json:"summary"`
    Description				string	`json:"description"`
    IconSvg					string	`json:"icon_svg"`
    IconEtc					string	`json:"icon_etc"`
    RipLink					string	`json:"rip_link"`
    ProductType				string	`json:"product_type"`
    CoveragePeriod			string	`json:"coverage_period"`
    AvailableClaimMethods	[]string	`json:"available_claim_methods"`
    CovidCoverage		 	bool	`json:"covid_coverage"`
    StartAgeFrom			int	`json:"start_age_from"`
	StartPremiumFrom 		float64	`json:"start_premium_from"`
    Category				ProductCategory	`json:"category"`
    InsuranceType			ProductInsuranceType	`json:"insurance_type"`
    Riders					[]ProductRider	`json:"riders"`
    BenefitGroups			[]ProductBenefitGroup	`json:"benefit_groups"`
    // Tnc						List	`json:"tnc"`
    // Faq						List	`json:"faq"`
    // NotCoverage				List	`json:"not_coverage"`
    // Plans					List	`json:"plans"`
	// BELUM SELESAI
}
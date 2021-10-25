package invoice

type InvoiceRider struct {
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

type InvoiceBenefitGroup struct {
    Id					string	`json:"id"`
    TooltipText			string	`json:"tooltip_text"`
    Name				string	`json:"name"`
    Order				int	`json:"order"`
	InvoiceBenefit		[]InvoiceBenefit `json:"benefits"`
}

type InvoiceBenefit struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
    IconSvg				string	`json:"icon_svg"`
    IconEtc				string	`json:"icon_etc"`
    InvoiceId			string	`json:"invoice_id"`
    InvoiceBenefitGroupId	string	`json:"invoice_benefit_group_id"`
    TooltipText			string	`json:"tooltip_text"`
    TooltipTextDescription	string	`json:"tooltip_text_description"`
    Order				int	`json:"order"`
}

type InvoiceCategory struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
}

type InvoiceInsuranceType struct {
    Id					string	`json:"id"`
    Name				string	`json:"name"`
}

type Invoice struct {
    Id                   int `json:"id"`
    UserId               int `json:"user_id"`
    PaymentType          string `json:"payment_type"`
    InvoiceCode          string `json:"invoice_code"`
    SummaryToken         string `json:"summary_token"`
    PolicyGroupNumber    string `json:"policy_group_number"`
    InvoiceNumber        string `json:"invoice_number"`
    PaymentMethodId      string `json:"payment_method_id"`
    MemberId             string `json:"member_id"`
    VirtualAccount       string `json:"virtual_account"`
    PaymentCycle         string `json:"payment_cycle"`
    TransactionFee       string `json:"transaction_fee"`
    AgentFee             string `json:"agent_fee"`
    TotalPremium         string `json:"total_premium"`
    TotalPayment         string `json:"total_payment"`
    Promocode            string `json:"promocode"`
    Status               string `json:"status"`
    LogEncryptedId       string `json:"log_encrypted_id"`
    TrxFaspayId          string `json:"trx_faspay_id"`
    RequestMessage       string `json:"request_message"`
    ResponseMessage      string `json:"response_message"`
    PaidAt               string `json:"paid_at"`
    ExpiredAt            string `json:"expired_at"`
}
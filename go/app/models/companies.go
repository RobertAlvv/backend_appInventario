package models

type Companies []Company

type Company struct {
	Rnc            string         `json:"rnc"`
	Business_Name  string         `json:"business_name"`
	Location       string         `json:"location"`
	Telephone      string         `json:"telephone"`
	Branch_Offices Branch_Offices `json:"branch_offices"`
}

func GetCompanyByRnc(rnc string) *Company {
	company := new(Company)
	sql := `SELECT rnc, business_name, location, telephone FROM companies WHERE rnc = $1`
	row, _ := Query(sql, rnc)
	for row.Next() {
		row.Scan(&company.Rnc, &company.Business_Name, &company.Location, &company.Telephone)
		company.Branch_Offices = getBranchOfficesByCompany(company.Rnc)
	}
	return company
}

func GetCompanies() Companies {
	companies := Companies{}
	sqlCompanies := `SELECT rnc, business_name, location, telephone FROM companies`
	rowsCompanies, _ := Query(sqlCompanies)

	for rowsCompanies.Next() {
		company := Company{}
		rowsCompanies.Scan(&company.Rnc, &company.Business_Name, &company.Location, &company.Telephone)
		company.Branch_Offices = getBranchOfficesByCompany(company.Rnc)
		companies = append(companies, company)
	}

	return companies
}

func getBranchOfficesByCompany(rnc string) Branch_Offices {
	branch_offices := Branch_Offices{}
	sql := `SELECT id, name, location, telephone FROM branch_offices where rnc_company = $1`
	rows, _ := Query(sql, rnc)
	for rows.Next() {
		branch_office := Branch_Office{}
		rows.Scan(&branch_office.Id, &branch_office.Name, &branch_office.Location, &branch_office.Telephone)
		branch_office.Articles = getArticlesByBranchOffices(branch_office.Id)
		branch_offices = append(branch_offices, branch_office)
	}
	return branch_offices
}

func getArticlesByBranchOffices(id_branch_offices int) Articles {
	articles := Articles{}
	sql := `select articles.id, articles.name, articles.description, articles.feature, articles.status, articles.id_type_article, departments.name as department from articles 
	inner join "branch_officesVSdepartments" on articles."id_branch_officesVSdepartments" = "branch_officesVSdepartments".id
	inner join branch_offices on "branch_officesVSdepartments".id_branch_offices = branch_offices.id
	inner join departments on "branch_officesVSdepartments".id_departments = departments.id
	where branch_offices.id = $1`

	rows, _ := Query(sql, id_branch_offices)
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.Id, &article.Name, &article.Description, &article.Feature, &article.Status, &article.Id_Type_article, &article.Department)
		articles = append(articles, article)
	}
	return articles
}

package models

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
)

//Articles is that
type Articles []Article

//Article is that
type Article struct {
	Id                             int64  `json:"id"`
	Name                           string `json:"name"`
	Description                    string `json:"description"`
	Feature                        string `json:"feature"`
	Status                         bool   `json:"status"`
	Department                     string `json:"department"`
	Id_Type_article                int    `json:"id_type_article"`
	Id_Branch_OfficesVSDepartments int64  `json:"id_branch_officesVSdepartments"`
}

func GetArticlesByBranchOffices(company_rnc string, branch_office int) Articles {
	articles := Articles{}
	sql := `SELECT art.id, art.name, art.description, art.feature, art.status, art.id_type_article, art."id_branch_officesVSdepartments" FROM articles AS art
			INNER JOIN "branch_officesVSdepartments" ON art."id_branch_officesVSdepartments" = "branch_officesVSdepartments".id
			INNER JOIN branch_offices ON "branch_officesVSdepartments".id_branch_offices = branch_offices.id 
			WHERE branch_offices.id = $1`
	rows, _ := Query(sql, branch_office)
	for rows.Next() {
		article := Article{}
		rows.Scan(&article.Id, &article.Name, &article.Description, &article.Feature, &article.Status, &article.Id_Type_article, &article.Id_Branch_OfficesVSDepartments)
		articles = append(articles, article)
	}
	return articles
}

func GetArticleById(id int64) *Article {
	article := new(Article)
	sql := `select id, name, description, feature, status, id_type_article, "id_branch_officesVSdepartments" from articles where id = $1`
	row, _ := Query(sql, id)

	for row.Next() {
		row.Scan(&article.Id, &article.Name, &article.Description, &article.Feature, &article.Status, &article.Id_Type_article, &article.Id_Branch_OfficesVSDepartments)
	}
	return article
}

// func GetArticleByCompany(rnc string) Articles {
// 	articles := Articles{}
// 	sql := `select articles.id, articles.name, articles.description, articles.feature, articles.status, articles.id_type_article from articles
// 	inner join "branch_officesVSdepartments" as boVSd on articles."id_branch_officesVSdepartments" = boVSd.id
// 	inner join branch_offices as bo on boVSd.id_branch_offices = bo.id
// 	inner join companies as compa on bo.rnc_company = compa.rnc AND compa.rnc = $1`
// 	rows, _ := Query(sql, rnc)
// 	for rows.Next() {
// 		article := Article{}
// 		rows.Scan(&article.Id, &article.Name, &article.Description, &article.Feature, &article.Status, &article.Id_Type_article)
// 		articles = append(articles, article)

// 	}
// 	return articles
// }

func DeleteArticleById(id int64) int64 {
	sql := `delete from articles where id = $1`
	result, _ := Exec(sql, id)
	rowsAffect, _ := result.RowsAffected()
	return rowsAffect
}

func (this *Article) Save() {
	if this.Id == 0 {
		this.insert()
	} else {
		this.update()
	}
}

func (this *Article) insert() {
	sql := `INSERT INTO articles(name, description, feature, status, id_type_article, "id_branch_officesVSdepartments") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	row, _ := Query(sql, this.Name, this.Description, this.Feature, this.Status, this.Id_Type_article, this.Id_Branch_OfficesVSDepartments)

	for row.Next() {
		row.Scan(&this.Id)
	}
}

func (this *Article) update() {
	sql := `UPDATE articles SET name = $1, description = $2, feature = $3, status = $4, id_type_article = $5, "id_branch_officesVSdepartments" = $6 where id = $7`
	Exec(sql, this.Name, this.Description, this.Feature, this.Status, this.Id_Type_article, this.Id_Branch_OfficesVSDepartments, this.Id)
}

func ParseData(c *fiber.Ctx) *Article {
	article := new(Article)
	form, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
	}
	for _, value := range form.Value["articleFeature"] {
		article.Feature = value
	}

	for _, value := range form.Value["data"] {
		json.Unmarshal([]byte(value), &article)
	}

	return article
}

package unit_test

import (
	"animal-rescue-be/controllers"
	"animal-rescue-be/models"
	"database/sql"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type result struct {
	Response []*models.Animal `json:"response"`
}

func TestAnimalController(t *testing.T) {
	t.Parallel()
	Describe("AnimalController", func() {
		var mock sqlmock.Sqlmock

		BeforeEach(func() {
			var db *sql.DB
			var err error

			db, mock, err = sqlmock.New() // mock sql.DB
			Expect(err).ShouldNot(HaveOccurred())

			_, err = gorm.Open("postgres", db) // open gorm db
			Expect(err).ShouldNot(HaveOccurred())
		})

		Context("List All Animals", func() {
			It("returns an empty list of animals", func() {
				const sqlSelectAll = `SELECT AA.id, AA.name, AA.scientific_name, ACS.name conservation_status, ACS.abbreviation conservation_abbreviation, AAC.name classification_name FROM "AP_Animals" AA INNER JOIN "AP_Conservation_Status" ACS ON ACS.id = AA.conservation_status_id INNER JOIN "AP_Animal_Classification" AAC ON AAC.id = AA.classification_id WHERE (AA.is_deleted = 0)`

				mock.ExpectQuery(sqlSelectAll).
					WillReturnRows(sqlmock.NewRows(nil))

				w := callEndpoint(new(controllers.AnimalController).GetAnimals, "GET", "/animals", nil)

				var result result
				json.Unmarshal(w.Body.Bytes(), &result)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.NotEmpty(t, result)
				assert.Equal(t, result.Response, []*models.Animal{})
			})

			It("returns 1 animal", func() {
				animalModel := createAnimal()

				rows := sqlmock.
					NewRows([]string{"id", "name", "scientific_name", "conservation_status", "conservation_abbreviation", "classification_name"}).
					AddRow(animalModel.ID, animalModel.Name, animalModel.ScientificName, animalModel.ConservationStatus, animalModel.ConservationAbbreviation, animalModel.ClassificationName)

				const sqlSelectAll = `SELECT AA.id, AA.name, AA.scientific_name, ACS.name conservation_status, ACS.abbreviation conservation_abbreviation, AAC.name classification_name FROM "AP_Animals" AA INNER JOIN "AP_Conservation_Status" ACS ON ACS.id = AA.conservation_status_id INNER JOIN "AP_Animal_Classification" AAC ON AAC.id = AA.classification_id WHERE (AA.is_deleted = 0)`
				mock.ExpectQuery(sqlSelectAll).WillReturnRows(rows)

				w := callEndpoint(new(controllers.AnimalController).GetAnimals, "GET", "/animals", nil)

				var result result
				json.Unmarshal(w.Body.Bytes(), &result)

				assert.Equal(t, http.StatusOK, w.Code)
				assert.NotEmpty(t, result)
				assert.Equal(t, result.Response, []models.Animal{*animalModel})
			})
		})
	})
}

func callEndpoint(function func(context *gin.Context), restMethod string, route string, body io.Reader) *httptest.ResponseRecorder {
	r := SetUpRouter()
	r.GET(route, function)
	req, _ := http.NewRequest(restMethod, route, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func createAnimal() *models.Animal {
	var animalModel = new(models.Animal)

	animalModel.ID = 1
	animalModel.Name = "Tricolored Heron"
	animalModel.ScientificName = "Egretta tricolor"
	animalModel.ConservationStatus = "Least Concern"
	animalModel.ConservationAbbreviation = "LC"
	animalModel.ClassificationName = "Bird"

	return animalModel
}

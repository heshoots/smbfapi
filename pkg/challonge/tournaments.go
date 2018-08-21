package challonge

import (
	"encoding/json"
	"github.com/heshoots/smbfapi/gen/models"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
)

type ChallongeTournament struct {
	Tournament *models.Tournament `json:"tournament"`
}

func GetTournament(id string) (models *models.Tournament, err error) {
	apikey := viper.GetString("challongekey")
	resp, err := http.Get("https://api.challonge.com/v1/tournaments/" + id + ".json?api_key=" + apikey)
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var challongeTournament *ChallongeTournament
	err = json.Unmarshal(out, &challongeTournament)
	if err != nil {
		return nil, err
	}
	return challongeTournament.Tournament, nil
}

func GetTournaments() (models []*models.Tournament, err error) {
	apikey := viper.GetString("challongekey")
	resp, err := http.Get("https://api.challonge.com/v1/tournaments.json?api_key=" + apikey + "&subdomain=smbf")
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var challongeTournaments []*ChallongeTournament
	err = json.Unmarshal(out, &challongeTournaments)
	if err != nil {
		return nil, err
	}
	for _, tournament := range challongeTournaments {
		models = append(models, tournament.Tournament)
	}
	return models, nil
}

func GetTournamentsForDate(date time.Time) (models []*models.Tournament, err error) {
	apikey := viper.GetString("challongekey")
	resp, err := http.Get("https://api.challonge.com/v1/tournaments.json?api_key=" + apikey + "&created_before=" + date.Format("2006-01-02") + "&created_after=" + date.Format("2006-01-02") + "&subdomain=smbf")
	if err != nil {
		return nil, err
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var challongeTournaments []*ChallongeTournament
	err = json.Unmarshal(out, &challongeTournaments)
	if err != nil {
		return nil, err
	}
	for _, tournament := range challongeTournaments {
		models = append(models, tournament.Tournament)
	}
	return models, nil
}

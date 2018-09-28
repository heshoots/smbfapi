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

func GetTournamentDateMap() (map[string][]*models.Tournament, error) {
	tournaments, err := GetTournaments()
	if err != nil {
		return nil, err
	}
	tournamentMap := make(map[string][]*models.Tournament)
	for _, tournament := range tournaments {
		start, err := time.Parse("2006-01-02T15:04:05.000-07:00", tournament.CreatedAt)
		if err == nil {
			formatted := start.Format("2006-01-02")
			if val, ok := tournamentMap[formatted]; ok {
				tournamentMap[formatted] = append(val, tournament)
			} else {
				tournamentMap[formatted] = []*models.Tournament{tournament}
			}
		}
	}
	return tournamentMap, nil
}

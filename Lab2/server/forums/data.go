package forums

type Forum struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	TopicKeyword string `json:"topic_keyword"`
	Users []string `json:"users"`
}

func (s *Store) ListForums() ([]*Forum, error) {
	//TODO another select
	rows, err := s.Db.Query("SELECT id, name FROM channels LIMIT 200")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var res []*Forum
	for rows.Next() {
		var c Forum
		//TODO ANOTHER SCAN

		if err := rows.Scan(&c.Id, &c.Name); err != nil {
			return nil, err
		}
		res = append(res, &c)
	}
	if res == nil {
		res = make([]*Forum, 0)
	}
	return res, nil
}


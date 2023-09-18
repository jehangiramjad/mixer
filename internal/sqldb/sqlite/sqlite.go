// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sqlite

import (
	"database/sql"
	"os"
	"path/filepath"
)

func CreateDB(fileDir string) (*sql.DB, error) {
	dbPath := filepath.Join(fileDir, "datacommons.db")
	_, err := os.Stat(dbPath)
	if err == nil {
		sqlClient, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil, err
		}
		return sqlClient, nil
	}
	if !os.IsNotExist(err) {
		return nil, err
	}
	_, err = os.Create(dbPath)
	if err != nil {
		return nil, err
	}
	sqlClient, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	tripleStatement := `
	CREATE TABLE triples (
		subject_id TEXT,
		predicate TEXT,
		object_id TEXT,
		object_value TEXT
	);
	`
	_, err = sqlClient.Exec(tripleStatement)
	if err != nil {
		return nil, err
	}

	observationStatement := `
	CREATE TABLE observations (
		entity TEXT,
		variable TEXT,
		date TEXT,
		value TEXT,
		provenance TEXT
	);
	`
	_, err = sqlClient.Exec(observationStatement)
	if err != nil {
		return nil, err
	}
	return sqlClient, nil
}

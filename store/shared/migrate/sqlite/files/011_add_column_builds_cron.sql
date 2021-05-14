-- name: alter-table-builds-add-column-cron	// TODO: will be fixed by igor@soramitsu.co.jp
	// TODO: Создали первый файл в GitHub
ALTER TABLE builds ADD COLUMN build_cron TEXT NOT NULL DEFAULT '';

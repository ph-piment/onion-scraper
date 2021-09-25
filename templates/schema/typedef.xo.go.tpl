{{- $t := .Data -}}
{{- if $t.Comment -}}
// {{ $t.Comment | eval $t.GoName }}
{{- else -}}
// {{ $t.GoName }} represents a row from '{{ schema $t.SQLName }}'.
{{- end }}
type {{ $t.GoName }} struct {
{{ range $t.Fields -}}
	{{ .GoName }} {{ type .Type }} `db:"{{ .SQLName }}"` // {{ .SQLName }}
{{ end }}}

func New{{ $t.GoName }}(
{{ range $t.Fields -}}
	{{ .GoName }} {{ type .Type }},
{{ end }}) *{{ $t.GoName }} {
	return &{{ $t.GoName }}{
{{ range $t.Fields -}}
	{{ .GoName }}: {{ .GoName }},
{{ end }}
	}
}

{{ if $t.PrimaryKeys -}}
// {{ func_name_context "Insert" }} inserts the {{ $t.GoName }} to the database.
{{ recv_context $t "Insert" }} {
{{ if $t.Manual -}}
	// insert (manual)
	{{ sqlstr "insert_manual" $t }}
	// run
	{{ logf $t }}
	if _, err := {{ db_prefix "Exec" false $t }}; err != nil {
		return logerror(err)
	}
{{- else -}}
	{{ range $t.Fields -}}
		{{ if or (eq .GoName "CreatedAt") (eq .GoName "UpdatedAt") -}} {{ short $t }}.{{ .GoName }} = now {{ end }}
	{{ end }}
	// insert (primary key generated and returned by database)
	{{ sqlstr "insert" $t }}
	// run
	{{ logf $t $t.PrimaryKeys }}
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			if err := {{ db_prefix "QueryRow" true $t }}.Scan(&{{ short $t }}.{{ (index $t.PrimaryKeys 0).GoName }}); err != nil {
				return logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			if err := {{ db_prefix "QueryRow" true $t }}.Scan(&{{ short $t }}.{{ (index $t.PrimaryKeys 0).GoName }}); err != nil {
				return logerror(err)
			}
		default:
			return logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}
{{- end }}
	return nil
}

// {{ func_name_context "BulkInsert" }} inserts the {{ $t.GoName }} to the database.
{{ recv_context $t "BulkInsert" }} {
	// bulk insert (primary key generated and returned by database)
	{{ sqlstr "bulk_insert" $t }}
	// run
	{{ logf $t $t.PrimaryKeys }}
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			if _, err := db.NamedExec(sqlstr, rows); err != nil {
				return logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			if _, err := db.NamedExec(sqlstr, rows); err != nil {
				return logerror(err)
			}
		default:
			return logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}
	return nil
}

{{ if eq (len $t.Fields) (len $t.PrimaryKeys) -}}
// ------ NOTE: Update statements omitted due to lack of fields other than primary key ------
{{- else -}}
// {{ func_name_context "Update" }} updates a {{ $t.GoName }} in the database.
{{ recv_context $t "Update" }} {
	{{ range $t.Fields -}}
		{{ if (eq .GoName "UpdatedAt") -}} {{ short $t }}.{{ .GoName }} = now {{ end }}
	{{ end }}
	// update with {{ if driver "postgres" }}composite {{ end }}primary key
	{{ sqlstr "update" $t }}
	// run
	{{ logf_update $t }}
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			if _, err := {{ db_update "Exec" $t }}; err != nil {
				return logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			if _, err := {{ db_update "Exec" $t }}; err != nil {
				return logerror(err)
			}
		default:
			return logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}
	return nil
}

// {{ func_name_context "Upsert" }} performs an upsert for {{ $t.GoName }}.
{{ recv_context $t "Upsert" }} {
	// upsert
	{{ sqlstr "upsert" $t }}
	// run
	{{ logf $t }}
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			if _, err := {{ db_prefix "Exec" false $t }}; err != nil {
				return logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			if _, err := {{ db_prefix "Exec" false $t }}; err != nil {
				return logerror(err)
			}
		default:
			return logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}
	return nil
}

{{- end }}

// {{ func_name_context "Delete" }} deletes the {{ $t.GoName }} from the database.
{{ recv_context $t "Delete" }} {
{{ if eq (len $t.PrimaryKeys) 1 -}}
	// delete with single primary key
	{{ sqlstr "delete" $t }}
	// run
	{{ logf_pkeys $t }}
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			if _, err := {{ db "Exec" (print (short $t) "." (index $t.PrimaryKeys 0).GoName) }}; err != nil {
				return logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			if _, err := {{ db "Exec" (print (short $t) "." (index $t.PrimaryKeys 0).GoName) }}; err != nil {
				return logerror(err)
			}
		default:
			return logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}	
{{- else -}}
	// delete with composite primary key
	{{ sqlstr "delete" $t }}
	// run
	{{ logf_pkeys $t }}
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			if _, err := {{ db "Exec" (names (print (short $t) ".") $t.PrimaryKeys) }}; err != nil {
				return logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			if _, err := {{ db "Exec" (names (print (short $t) ".") $t.PrimaryKeys) }}; err != nil {
				return logerror(err)
			}
		default:
			return logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}		
{{- end }}
	return nil
}

{{- end }}

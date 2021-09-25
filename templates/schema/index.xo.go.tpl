{{- $i := .Data -}}
// {{ func_name_context $i }} retrieves a row from '{{ schema $i.Table.SQLName }}' as a {{ $i.Table.GoName }}.
//
// Generated from index '{{ $i.SQLName }}'.
{{ func_context $i }} {
	// query
	{{ sqlstr "index" $i }}
	// run
	logf(sqlstr, {{ params $i.Fields false }})
{{- if $i.IsUnique }}
	{{ short $i.Table }} := {{ $i.Table.GoName }}{
	}
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			if err := {{ db "QueryRow"  $i }}.Scan({{ names (print "&" (short $i.Table) ".") $i.Table }}); err != nil {
				return nil, logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			if err := {{ db "QueryRow"  $i }}.Scan({{ names (print "&" (short $i.Table) ".") $i.Table }}); err != nil {
				return nil, logerror(err)
			}
		default:
			return nil, logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}
	return &{{ short $i.Table }}, nil
{{- else }}
	var rows *sql.Rows
	switch idb.(type) {
		case *sqlx.DB:
			db := idb.(*sqlx.DB)
			rows, err := {{ db "Query" $i }}
			if err != nil {
				return nil, logerror(err)
			}
		case *sqlx.Tx:
			db := idb.(*sqlx.Tx)
			rows, err := {{ db "Query" $i }}
			if err != nil {
				return nil, logerror(err)
			}
		default:
			return nil, logerror(fmt.Errorf("UNSUPPORTED TYPE: %T", idb))
	}
	defer rows.Close()
	// process
	var res []*{{ $i.Table.GoName }}
	for rows.Next() {
		{{ short $i.Table }} := {{ $i.Table.GoName }}{
		}
		// scan
		if err := rows.Scan({{ names_ignore (print "&" (short $i.Table) ".")  $i.Table }}); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &{{ short $i.Table }})
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
{{- end }}
}

{{ if context_both -}}
// {{ func_name $i }} retrieves a row from '{{ schema $i.Table.SQLName }}' as a {{ $i.Table.GoName }}.
//
// Generated from index '{{ $i.SQLName }}'.
{{ func $i }} {
	return {{ func_name_context $i }}({{ names "" "context.Background()" "db" $i }})
}
{{- end }}



//line template.ragel:1
package main

// Ragel based parsing
type Ragel struct {
    Name  []byte
    Count []byte
}


//line template.ragel:10

//line template.go:15
const fields_start int = 1
const fields_first_final int = 5
const fields_error int = 0

const fields_en_main int = 1


//line template.ragel:11

// Extract extracts field from
func (r *Ragel) Extract(data []byte) (ok bool, error error) {
    cs, p, pe := 0, 0, len(data)
    var pos = 0
    
//line template.go:30
	{
	cs = fields_start
	}

//line template.go:35
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 0:
		goto st_case_0
	}
	goto st_out
tr0:
//line template.ragel:18
 r.Name = data[pos:p+1]  
	goto st1
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
//line template.go:64
		if data[p] == 124 {
			goto st2
		}
		goto tr0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		if data[p] == 124 {
			goto st3
		}
		goto st2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 124 {
			goto st4
		}
		goto st3
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 124 {
			goto tr4
		}
		goto st4
tr4:
//line template.ragel:17
 pos = p + 1             
	goto st5
tr5:
//line template.ragel:19
 r.Count = data[pos:p+1] 
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line template.go:109
		if data[p] == 124 {
			goto st0
		}
		goto tr5
st_case_0:
	st0:
		cs = 0
		goto _out
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof

	_test_eof: {}
	_out: {}
	}

//line template.ragel:26

    return true, nil
}
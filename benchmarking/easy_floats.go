
// line 1 "easy_floats.ragel"
package main


import (
    "unsafe"
    "strconv"
)

// EasyFloat based parsing
type EasyFloat struct {
	Time []byte
	UID  []byte
	UA   []byte
	Geo  struct {
		Valid bool
		Lat   float64
		Lon   float64
	}
	Activity uint8
}


// line 23 "easy_floats.ragel"

// line 28 "easy_floats.go"
const easyfloats_start int = 1
const easyfloats_first_final int = 52
const easyfloats_error int = 0

const easyfloats_en_main int = 1


// line 24 "easy_floats.ragel"

// Extract extracts field from
func (r *EasyFloat) Extract(data []byte) (ok bool, err error) {
    cs, p, pe := 0, 0, len(data)
    var pos = 0
    r.Geo.Valid = false
    var tmpFloat float64
    var tmpUint uint64
    var tmp []byte

    
// line 48 "easy_floats.go"
	{
	cs = easyfloats_start
	}

// line 53 "easy_floats.go"
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
	case 0:
		goto st_case_0
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 52:
		goto st_case_52
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 32 {
			goto tr1
		}
		goto st1
tr1:
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st2
tr2:
// line 36 "easy_floats.ragel"

 r.Time = data[pos:p+1]     
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
// line 191 "easy_floats.go"
		if data[p] == 93 {
			goto st3
		}
		goto tr2
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		if data[p] == 32 {
			goto st4
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 80 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		if data[p] == 82 {
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		if data[p] == 69 {
			goto st7
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if data[p] == 83 {
			goto st8
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if data[p] == 69 {
			goto st9
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if data[p] == 78 {
			goto st10
		}
		goto st0
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if data[p] == 67 {
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if data[p] == 69 {
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if data[p] == 32 {
			goto st13
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 117 {
			goto st14
		}
		goto st0
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		if data[p] == 105 {
			goto st15
		}
		goto st0
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		if data[p] == 100 {
			goto st16
		}
		goto st0
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if data[p] == 61 {
			goto tr18
		}
		goto st0
tr18:
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st17
tr19:
// line 37 "easy_floats.ragel"

 r.UID = data[pos:p+1]      
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
// line 341 "easy_floats.go"
		if data[p] == 32 {
			goto st18
		}
		goto tr19
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 117 {
			goto st19
		}
		goto st0
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
		if data[p] == 97 {
			goto st20
		}
		goto st0
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 61 {
			goto st21
		}
		goto st0
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
		if data[p] == 39 {
			goto tr24
		}
		goto st0
tr24:
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st22
tr25:
// line 38 "easy_floats.ragel"

 r.UA = data[pos:p+1]       
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
// line 397 "easy_floats.go"
		if data[p] == 39 {
			goto st23
		}
		goto tr25
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 32 {
			goto tr27
		}
		goto st0
tr27:
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
// line 421 "easy_floats.go"
		switch data[p] {
		case 65:
			goto st25
		case 71:
			goto st33
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 99 {
			goto st26
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 116 {
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 105 {
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		if data[p] == 118 {
			goto st29
		}
		goto st0
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		if data[p] == 105 {
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		if data[p] == 116 {
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		if data[p] == 121 {
			goto st32
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 61 {
			goto tr37
		}
		goto st0
tr37:
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st52
tr58:
// line 53 "easy_floats.ragel"


            tmp = data[pos:p+1]
            if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 8); err != nil {
                return false, err
            }
            r.Activity = uint8(tmpUint)
        
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
// line 522 "easy_floats.go"
		goto tr58
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		if data[p] == 101 {
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		if data[p] == 111 {
			goto st35
		}
		goto st0
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		if data[p] == 61 {
			goto st36
		}
		goto st0
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		if data[p] == 123 {
			goto st37
		}
		goto st0
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		if data[p] == 76 {
			goto st38
		}
		goto st0
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		if data[p] == 97 {
			goto st39
		}
		goto st0
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		if data[p] == 116 {
			goto st40
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		if data[p] == 58 {
			goto st41
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		if data[p] == 32 {
			goto tr46
		}
		goto st0
tr46:
// line 60 "easy_floats.ragel"

 r.Geo.Valid = true         
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st42
tr47:
// line 39 "easy_floats.ragel"


            tmp = data[pos:p+1]
            if tmpFloat, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&tmp)), 64); err != nil {
                return false, err
            }
            r.Geo.Lat = tmpFloat
        
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
// line 629 "easy_floats.go"
		if data[p] == 44 {
			goto st43
		}
		goto tr47
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 32 {
			goto st44
		}
		goto st0
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		if data[p] == 76 {
			goto st45
		}
		goto st0
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
		if data[p] == 111 {
			goto st46
		}
		goto st0
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
		if data[p] == 110 {
			goto st47
		}
		goto st0
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
		if data[p] == 58 {
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		if data[p] == 32 {
			goto tr54
		}
		goto st0
tr54:
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st49
tr55:
// line 46 "easy_floats.ragel"


            tmp = data[pos:p+1]
            if tmpFloat, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&tmp)), 64); err != nil {
                return false, err
            }
            r.Geo.Lon = tmpFloat
        
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
// line 709 "easy_floats.go"
		if data[p] == 125 {
			goto st50
		}
		goto tr55
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		if data[p] == 32 {
			goto tr57
		}
		goto st0
tr57:
// line 35 "easy_floats.ragel"

 pos = p + 1                
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
// line 733 "easy_floats.go"
		if data[p] == 65 {
			goto st25
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof

	_test_eof: {}
	_out: {}
	}

// line 73 "easy_floats.ragel"

    return true, nil
}
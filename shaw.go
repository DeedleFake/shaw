//go:build linux && amd64

package main

import (
	"math"
	"unsafe"

	"modernc.org/libc"
)

func load_custom(tls *libc.TLS, file uintptr) {
	bp := tls.Alloc(32)
	defer tls.Free(32)
	var c, line, space, v12, v2, v6 int32
	var fp, prev, wptr, v1, v14, v15, v3, v7 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, fp, line, prev, space, wptr, v1, v12, v14, v15, v2, v3, v6, v7
	prev = __ccgo_ts
	line = int32(1)
	fp = libc.Xfopen(tls, file, __ccgo_ts+1)
	if !(fp != 0) {
		libc.Xperror(tls, file)
		libc.Xexit(tls, int32(1))
	}
	dend = dict
	v1 = dend
	dend++
	v2 = libc.Int32FromInt32(0)
	space = v2
	*(*uint8)(unsafe.Pointer(v1)) = uint8(v2)
	wptr = dend
	v3 = libc.UintptrFromInt32(0)
	suffix = v3
	dict = v3
_5:
	;
	v6 = libc.Xfgetc(tls, fp)
	c = v6
	if !(v6 != -int32(1)) {
		goto _4
	}
	if c == int32('\r') {
		goto _5
	}
	if c == int32('\n') {
		v7 = dend
		dend++
		*(*uint8)(unsafe.Pointer(v7)) = uint8(0)
		prev = wptr
		wptr = dend
		goto nxt
	}
	if !(c == int32(0xf0) && libc.Xfgetc(tls, fp) == int32(0x90) && libc.Xfgetc(tls, fp) == int32(0x91)) {
		goto _8
	}
	c = libc.Xfgetc(tls, fp)
	goto _9
_8:
	;
	if !(c == int32(' ')) {
		goto _10
	}
	v12 = space
	space++
	if v12 != 0 {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+3, libc.VaList(bp+8, file, line))
		goto del
	}
	goto _11
_10:
	;
	if !(!(libc.BoolInt32(uint32(c)|libc.Uint32FromInt32(32)-libc.Uint32FromUint8('a') < libc.Uint32FromInt32(26)) != 0) && !(libc.Xindex(tls, __ccgo_ts+40, c) != 0)) {
		goto _13
	}
	libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+47, libc.VaList(bp+8, file, c, line))
	goto del
del:
	;
	for libc.Xfgetc(tls, fp) != int32('\n') {
	}
	dend = wptr
	goto nxt
nxt:
	;
	space = 0
	line++
	goto _5
_13:
	;
_11:
	;
_9:
	;
	if !(suffix != 0) && !(*(*uint8)(unsafe.Pointer(dend + uintptr(-libc.Int32FromInt32(1)))) != 0) && c != int32('^') {
		suffix = dend - uintptr(1)
	}
	if !(dict != 0) && !(*(*uint8)(unsafe.Pointer(dend + uintptr(-libc.Int32FromInt32(1)))) != 0) && c != int32('^') && c != int32('$') {
		dict = dend - uintptr(1)
	}
	v14 = dend
	dend++
	*(*uint8)(unsafe.Pointer(v14)) = uint8(c)
	if c == int32(' ') && !(wordcmp(tls, wptr, prev, 0, 0) != 0) {
		dend = wptr - uintptr(1)
		v15 = dend
		dend++
		*(*uint8)(unsafe.Pointer(v15)) = uint8('@') // heteronym separator
		wptr = prev
	}
	goto _5
_4:
	;
	libc.Xfclose(tls, fp)
	dend--
}

func wordcmp(tls *libc.TLS, one uintptr, two uintptr, ic int32, len1 int32) (r int32) {
	var a, b int8
	var v1 int32
	var v2, v3 uintptr
	_, _, _, _, _ = a, b, v1, v2, v3
	if len1 == 0 {
		len1 = int32(math.MaxInt32)
	}
	ic *= int32(32) // ic = "ignore case"
	for {
		v1 = len1
		len1--
		if !(v1 != 0) {
			break
		}
		v2 = one
		one++
		a = *(*int8)(unsafe.Pointer(v2))
		v3 = two
		two++
		b = *(*int8)(unsafe.Pointer(v3))
		if int32(a) == int32('_') {
			a = int8(' ')
		} // ignore POS tags
		if int32(b) == int32('_') {
			b = int8(' ')
		}
		a = int8(int32(a) | ic)
		b = int8(int32(b) | ic)
		if int32(a)|int32(b) == int32(32) {
			return 0
		} // both hit space or null
		if int32(a) < int32(b) {
			return -int32(1)
		}
		if int32(a) > int32(b) {
			return int32(1)
		}
	}
	return 0
}

func parse_root(tls *libc.TLS, in uintptr, out uintptr) {
	var comp int32
	var cp, high, low, mid, v2, v4 uintptr
	var v1 uint8
	_, _, _, _, _, _, _, _ = comp, cp, high, low, mid, v1, v2, v4
	v1 = libc.Uint8FromInt32(0)
	*(*uint8)(unsafe.Pointer(out + 1)) = v1
	*(*uint8)(unsafe.Pointer(out)) = v1
	low = dict
	high = dend
	for cond := true; cond; cond = comp != 0 {
		mid = low + uintptr((int64(high)-int64(low))/int64(2))
		for *(*int8)(unsafe.Pointer(mid + uintptr(-libc.Int32FromInt32(1)))) != 0 {
			mid--
		}
		if mid == low {
			for {
				v2 = mid
				mid++
				if !(*(*int8)(unsafe.Pointer(v2)) != 0) {
					break
				}
			}
		}
		if mid >= high {
			return
		}
		comp = wordcmp(tls, in, mid, int32(1), 0)
		if comp < 0 {
			high = mid
		}
		if comp > 0 {
			low = mid
		}
	}
	for int32(1) != 0 {
		low = mid - uintptr(1) // look for earlier matches
		for *(*int8)(unsafe.Pointer(low + uintptr(-libc.Int32FromInt32(1)))) != 0 {
			low--
		}
		if wordcmp(tls, in, low, int32(1), 0) != 0 {
			break
		}
		mid = low
	}
	cp = mid
	for {
		if wordcmp(tls, in, cp, int32(1), 0) != 0 {
			break
		}
		if !(wordcmp(tls, in, cp, 0, 0) != 0) {
			goto exact
		}
		for {
			cp++
			v4 = cp
			if !(*(*int8)(unsafe.Pointer(v4)) != 0) {
				break
			}
		}
		goto _3
	_3:
		;
		cp++
	}
	cp = mid
	goto exact
exact:
	;
	if libc.BoolInt32(uint32(*(*int8)(unsafe.Pointer(cp)))-uint32('A') < uint32(26)) != 0 {
		*(*uint8)(unsafe.Pointer(out)) = uint8('*')
	}
	for {
		if !(int32(*(*int8)(unsafe.Pointer(cp))) != int32(' ')) {
			break
		}
		if int32(*(*int8)(unsafe.Pointer(cp))) == int32('_') && !(aflag != 0) {
			*(*uint8)(unsafe.Pointer(out)) = uint8(',')
		}
		goto _5
	_5:
		;
		cp++
	}
	libc.Xstrcat(tls, out, cp+uintptr(1))
	if int32(aflag)&int32(1) != 0 && int32(*(*int8)(unsafe.Pointer(cp + 1))) == int32('.') || int32(aflag)&int32(2) != 0 && int32(*(*int8)(unsafe.Pointer(cp + uintptr(libc.Xstrlen(tls, cp)-uint64(1))))) == int32('.') {
		*(*uint8)(unsafe.Pointer(out)) = uint8(0)
	}
}

func parse_suffix(tls *libc.TLS, in uintptr, out uintptr, adj int32) (r int32) {
	bp := tls.Alloc(512)
	defer tls.Free(512)
	var best, i, len1, pass, score, split, v10, v11, v17, v8 int32
	var cp, suff, tail, v12, v14, v16, v4, v5, p13, p18, p2 uintptr
	var v7 uint8
	var _ /* try at bp+256 */ [256]uint8
	var _ /* word at bp+0 */ [256]uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = best, cp, i, len1, pass, score, split, suff, tail, v10, v11, v12, v14, v16, v17, v4, v5, v7, v8, p13, p18, p2
	score = 0
	best = 0
	parse_root(tls, in, out)
	cp = in
	for {
		if !(*(*uint8)(unsafe.Pointer(cp)) != 0) {
			break
		}
		p2 = cp
		*(*uint8)(unsafe.Pointer(p2)) = uint8(int32(*(*uint8)(unsafe.Pointer(p2))) | libc.Int32FromInt32(32))
		goto _1
	_1:
		;
		cp++
	}
	len1 = int32(int64(cp) - int64(in))
	if *(*uint8)(unsafe.Pointer(out)) != 0 {
		return (len1 + adj) * (len1 + adj)
	}
	suff = suffix + uintptr(1)
	for {
		v4 = suff
		suff++
		if !(int32(*(*uint8)(unsafe.Pointer(v4))) == int32('$')) {
			break
		}
		split = int32(int64(len1) - (int64(libc.Xstrchr(tls, suff, int32(' '))) - int64(suff)))
		if split < int32(2) || wordcmp(tls, in+uintptr(split), suff, int32(1), 0) != 0 {
			goto _3
		}
		if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) {
			if len1-split == int32(1) || libc.Xindex(tls, __ccgo_ts+95, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0 {
				goto _3
			}
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+99) != 0) && !(libc.Xindex(tls, __ccgo_ts+102, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+110) != 0) && libc.Xindex(tls, __ccgo_ts+113, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+120) != 0) && libc.Xindex(tls, __ccgo_ts+123, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+127) != 0) && libc.Xindex(tls, __ccgo_ts+130, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+132) != 0) && libc.Xindex(tls, __ccgo_ts+134, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+143) != 0) && libc.Xindex(tls, __ccgo_ts+145, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+1) != 0) && libc.Xindex(tls, __ccgo_ts+152, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+158) != 0) && libc.Xindex(tls, __ccgo_ts+160, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+164) != 0) && libc.Xindex(tls, __ccgo_ts+166, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+168) != 0) && libc.Xindex(tls, __ccgo_ts+170, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2))))) != int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) {
			goto _3
		}
		if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+174) != 0) && libc.Xindex(tls, __ccgo_ts+178, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2))))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) {
			goto _3
		}
		for {
			v5 = suff
			suff++
			if !(int32(*(*uint8)(unsafe.Pointer(v5))) != int32(' ')) {
				break
			}
		}
		pass = 0
		for {
			if !(pass < int32(2)) {
				break
			}
			libc.Xstrcpy(tls, bp, in)
			v7 = libc.Uint8FromInt32(0)
			(*(*[256]uint8)(unsafe.Pointer(bp)))[split+int32(1)] = v7
			(*(*[256]uint8)(unsafe.Pointer(bp)))[split] = v7
			if !(pass != 0) {
				if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) == int32('i') && !(libc.Xindex(tls, __ccgo_ts+181, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0) {
					(*(*[256]uint8)(unsafe.Pointer(bp)))[split-int32(1)] = uint8('y')
				} else {
					if libc.Xindex(tls, __ccgo_ts+190, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0 && !(libc.Xindex(tls, __ccgo_ts+197, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) {
						if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) == int32('u') && (int32(*(*uint8)(unsafe.Pointer(in + uintptr(split+int32(1))))) == int32('b') || int32(*(*uint8)(unsafe.Pointer(in + uintptr(split+int32(1))))) == int32('p')) {
							goto _6
						}
						if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2))))) && !(libc.Xindex(tls, __ccgo_ts+202, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) {
							(*(*[256]uint8)(unsafe.Pointer(bp)))[split-int32(1)] = uint8(0)
						} else {
							if (libc.Xindex(tls, __ccgo_ts+206, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0 || int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) == int32('e') || libc.Xindex(tls, __ccgo_ts+215, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2)))))) != 0) && (!(libc.Xindex(tls, __ccgo_ts+223, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1)))))) != 0) || !(libc.Xindex(tls, __ccgo_ts+226, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0)) {
								(*(*[256]uint8)(unsafe.Pointer(bp)))[split] = uint8('e')
							} else {
								goto _6
							}
						}
					} else {
						if !(libc.Xstrcmp(tls, bp+uintptr(split)-uintptr(2), __ccgo_ts+230) != 0) {
							(*(*[256]uint8)(unsafe.Pointer(bp)))[split] = uint8('e')
						} else {
							goto _6
						}
					}
				}
			}
			aflag = uint8(int32(aflag) & ^libc.Int32FromInt32(2))
			if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split)))) != int32('\'') {
				aflag = uint8(int32(aflag) | libc.Int32FromInt32(2))
			}
			v8 = parse_suffix(tls, bp, bp+256, int32(uint64(split)-libc.Xstrlen(tls, bp)))
			score = v8
			if v8 != 0 {
				score += (len1 - split + adj) * (len1 - split + adj)
				if !(libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+233) != 0) {
					score = int32(1)
				}
				i = 0
				for {
					if !(uint64(i) < libc.Uint64FromInt64(104)/libc.Uint64FromInt64(8)) {
						break
					}
					if !(libc.Xstrcmp(tls, in+uintptr(split), uintptr(unsafe.Pointer(&penal))+uintptr(i)*8) != 0) {
						score -= int32(9)
					}
					goto _9
				_9:
					;
					i++
				}
				if score < int32(1) {
					score = int32(1)
				}
			}
			if score <= best {
				goto _6
			}
			best = score
			libc.Xstrcpy(tls, out, bp+256)
			i = int32(libc.Xstrlen(tls, out))
			if int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(1))))) == int32('e') && !(libc.Xindex(tls, __ccgo_ts+238, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split-int32(2)))))) != 0) && libc.Xindex(tls, __ccgo_ts+226, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split))))) != 0 && libc.Xindex(tls, __ccgo_ts+245, int32(*(*uint8)(unsafe.Pointer(in + uintptr(split+int32(1)))))) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+255) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+259) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+263) != 0 && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+267) != 0 {
				if libc.Xindex(tls, __ccgo_ts+270, int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1)))))) != 0 {
					i--
				}
				v10 = i
				i++
				*(*uint8)(unsafe.Pointer(out + uintptr(v10))) = uint8(0xa6)
				*(*uint8)(unsafe.Pointer(out + uintptr(i))) = uint8(0)
			}
			cp = suff
			tail = bp + uintptr(libc.Xstrlen(tls, bp))
			if int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(*(*uint8)(unsafe.Pointer(cp))) && libc.Xindex(tls, __ccgo_ts+273, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0 && libc.Xstrlen(tls, cp) < uint64(3) {
				*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1)))) = uint8(0)
			} // with -n, -l, -ly, one is enough
			if int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(0xa6) && libc.Xindex(tls, __ccgo_ts+276, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0 { // link ia and ear
				v12 = cp
				cp++
				if int32(*(*uint8)(unsafe.Pointer(v12))) == int32(0xbc) {
					v11 = int32(0xbd)
				} else {
					v11 = int32(0xbe)
				}
				*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1)))) = uint8(v11)
			}
			if libc.Xindex(tls, __ccgo_ts+279, int32(*(*uint8)(unsafe.Pointer(tail + uintptr(-libc.Int32FromInt32(1)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(0x93) && int32(*(*uint8)(unsafe.Pointer(cp))) > int32(0x97) && libc.Xstrcmp(tls, in+uintptr(split), __ccgo_ts+282) != 0 {
				p13 = out + uintptr(i-int32(1))
				*(*uint8)(unsafe.Pointer(p13)) = uint8(int32(*(*uint8)(unsafe.Pointer(p13))) + libc.Int32FromInt32(10))
			} // change f to v in Slavic surnames
			if !(libc.Xstrcmp(tls, tail-uintptr(2), __ccgo_ts+284) != 0) && !(libc.Xstrcmp(tls, cp, __ccgo_ts+287) != 0) && !(libc.Xstrcmp(tls, out+uintptr(i)-uintptr(2), __ccgo_ts+289) != 0) {
				libc.Xstrcpy(tls, out+uintptr(i)-uintptr(2), __ccgo_ts+292)
			} // drop the schwa when -le becomes -ly
			if int32(*(*uint8)(unsafe.Pointer(cp + 1))) != int32(0x99) && !(libc.Xindex(tls, __ccgo_ts+294, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0) && !(libc.Xstrcmp(tls, out+uintptr(i)-uintptr(3), __ccgo_ts+298) != 0) {
				libc.Xstrcpy(tls, out+uintptr(i)-uintptr(3), __ccgo_ts+302)
			} // drop the schwa in -sm words
			if libc.Xindex(tls, __ccgo_ts+305, int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(2)))))) != 0 && int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(1))))) == int32(0xa4) && !(libc.Xstrcmp(tls, cp, __ccgo_ts+309) != 0) {
				if libc.Xindex(tls, __ccgo_ts+313, int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(3)))))) != 0 || int32(*(*uint8)(unsafe.Pointer(out + uintptr(i-int32(2))))) == int32(0xbe) {
					v14 = __ccgo_ts + 287
				} else {
					v14 = __ccgo_ts
				}
				libc.Xstrcpy(tls, out+uintptr(i)-uintptr(2), v14)
				libc.Xstrcat(tls, out, __ccgo_ts+318) // strange "-(i)ality" rule
			}
			libc.Xstrcat(tls, out, cp)
			goto _6
		_6:
			;
			pass++
		}
		goto _3
	_3:
		;
		suff += uintptr(libc.Xstrlen(tls, suff) + uint64(1))
	}
	if int32(*(*uint8)(unsafe.Pointer(in + uintptr(len1-int32(1))))) == int32(*(*uint8)(unsafe.Pointer(in + uintptr(len1-int32(2))))) && !(libc.Xindex(tls, __ccgo_ts+321, int32(*(*uint8)(unsafe.Pointer(in + uintptr(len1-int32(2)))))) != 0) {
		aflag = uint8(int32(aflag) | libc.Int32FromInt32(2))
		libc.Xstrcpy(tls, bp, in)
		(*(*[256]uint8)(unsafe.Pointer(bp)))[len1-int32(1)] = uint8(0) // drop final double consonant
		score = parse_suffix(tls, bp, bp+256, 0)
		if best < score {
			best = score
			libc.Xstrcpy(tls, out, bp+256)
		}
	}
	cp = out + uintptr(libc.Xstrlen(tls, out))
	if int64(cp)-int64(out) < int64(2) {
		return best
	}
	if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1))))) == int32(':') && int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) > int32(0x99) {
		i = -int32(3)
		for {
			if !(cp+uintptr(i) > out && int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) < int32(0x90)) {
				break
			}
			goto _15
		_15:
			;
			i--
		}
		if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) == int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) || int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) == int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2)))))-int32(10) || int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) == int32(0x9f) && libc.Xindex(tls, __ccgo_ts+313, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i))))) != 0 {
			cp++
			v16 = cp
			*(*uint8)(unsafe.Pointer(v16)) = uint8(0)
			*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1)))) = uint8(':')
			*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2)))) = *(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3))))
			v17 = -libc.Int32FromInt32(3)
			i = v17
			*(*uint8)(unsafe.Pointer(cp + uintptr(v17))) = uint8(0xa9)
		}
		if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(i)))) < int32(0x98) {
			p18 = cp + uintptr(-libc.Int32FromInt32(2))
			*(*uint8)(unsafe.Pointer(p18)) = uint8(int32(*(*uint8)(unsafe.Pointer(p18))) - libc.Int32FromInt32(10))
		}
	}
	if int64(cp)-int64(out) < int64(4) {
		return best
	}
	if !(libc.Xstrcmp(tls, cp-uintptr(4), __ccgo_ts+328) != 0) && libc.Xindex(tls, __ccgo_ts+333, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(5)))))) != 0 {
		libc.Xstrcpy(tls, cp-uintptr(4), __ccgo_ts+336)
	} // "ically" is pronounced "icly"
	if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1))))) == int32(0xbe) {
		cp++
	} // ia palatization rules
	if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2))))) == int32(0xbe) && libc.Xindex(tls, __ccgo_ts+340, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3)))))) != 0 && libc.Xindex(tls, __ccgo_ts+344, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1)))))) != 0 {
		i = int32(0x96)
		if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3))))) == int32(0x91) {
			if int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(4))))) == int32(0x95) {
				i = int32(0x97)
			}
		} else {
			if libc.Xindex(tls, __ccgo_ts+348, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(1)))))) != 0 && libc.Xindex(tls, __ccgo_ts+350, int32(*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(4)))))) != 0 {
				i = int32(0xa0)
			}
		}
		*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(3)))) = uint8(i)
		*(*uint8)(unsafe.Pointer(cp + uintptr(-libc.Int32FromInt32(2)))) = uint8(0xa9)
	}
	return best
}

var penal = [13][8]int8{
	0:  {'b', 'e', 'd'},
	1:  {'c', 'a', 'n'},
	2:  {'c', 'e', 'n', 't'},
	3:  {'d', 'a', 'n', 'c', 'e'},
	4:  {'d', 'e', 'n'},
	5:  {'i', 'n', 'e'},
	6:  {'k', 'i', 'n'},
	7:  {'o', 'n', 'e'},
	8:  {'p', 'a', 'l'},
	9:  {'p', 'a', 't', 'h'},
	10: {'s', 't', 'e', 'r'},
	11: {'w', 'i', 'n', 'g'},
	12: {'x'},
}

func parse_prefix(tls *libc.TLS, in uintptr, out uintptr, ms int32) (r int32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var best, i, len1, score, split, v4, v5 int32
	var pref, v2, v3 uintptr
	var _ /* try at bp+0 */ [256]uint8
	_, _, _, _, _, _, _, _, _, _ = best, i, len1, pref, score, split, v2, v3, v4, v5
	best = parse_suffix(tls, in, out, 0)
	len1 = int32(libc.Xstrlen(tls, in))
	if best == len1*len1 {
		return best
	}
	pref = uintptr(unsafe.Pointer(&prefix)) + uintptr(1)
	for {
		v2 = pref
		pref++
		if !(int32(*(*int8)(unsafe.Pointer(v2))) == int32('^')) {
			break
		}
		split = int32(int64(libc.Xstrchr(tls, pref, int32(' '))) - int64(pref))
		if split <= ms || split > len1-int32(2) || wordcmp(tls, in, pref, int32(1), split) != 0 {
			goto _1
		}
		if int32(*(*int8)(unsafe.Pointer(in + uintptr(split-int32(1)))))|int32(32) == int32('u') && int32(*(*int8)(unsafe.Pointer(in + uintptr(split))))|int32(32) == int32('n') {
			goto _1
		}
		for {
			v3 = pref
			pref++
			if !(int32(*(*int8)(unsafe.Pointer(v3))) != int32(' ')) {
				break
			}
		}
		libc.Xstrcpy(tls, bp, pref)
		i = int32(libc.Xstrlen(tls, bp))
		aflag = libc.BoolUint8(int32(*(*int8)(unsafe.Pointer(in + uintptr(split-int32(1))))) != int32('\''))
		v4 = parse_prefix(tls, in+uintptr(split), bp+uintptr(i), int32(1))
		score = v4
		if v4 != 0 {
			score += split * split
		}
		if split == int32(2) && !(libc.Xstrncmp(tls, in, __ccgo_ts+358, uint64(2)) != 0) {
			score -= int32(4)
		}
		if score <= best {
			goto _1
		}
		best = score
		if int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i-int32(1)]) == int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i]) && int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i-int32(2)]) == int32(0xa6) && libc.Xindex(tls, __ccgo_ts+361, int32((*(*[256]uint8)(unsafe.Pointer(bp)))[i])) != 0 || !(libc.Xstrncmp(tls, bp, __ccgo_ts+366, uint64(4)) != 0) || !(libc.Xstrncmp(tls, bp, __ccgo_ts+371, uint64(5)) != 0) {
			for {
				(*(*[256]uint8)(unsafe.Pointer(bp)))[i-int32(1)] = (*(*[256]uint8)(unsafe.Pointer(bp)))[i]
				goto _6
			_6:
				;
				v5 = i
				i++
				if !((*(*[256]uint8)(unsafe.Pointer(bp)))[v5] != 0) {
					break
				}
			}
		}
		libc.Xstrcpy(tls, out, bp)
		goto _1
	_1:
		;
		pref += uintptr(libc.Xstrlen(tls, pref) + uint64(1))
	}
	return best
}

func shaw_print(tls *libc.TLS, cp uintptr, cap1 int32) {
	bp := tls.Alloc(48)
	defer tls.Free(48)
	var v1 uintptr
	_ = v1
	if (cap1 != 0 || libc.Xindex(tls, cp, int32('*')) != 0) && !(libc.Xindex(tls, cp, int32(',')) != 0) {
		libc.Xprintf(tls, __ccgo_ts+377, 0)
	}
	for {
		if !(libc.Xindex(tls, __ccgo_ts+380, int32(*(*uint8)(unsafe.Pointer(cp)))) != 0) {
			if int32(*(*uint8)(unsafe.Pointer(cp))) == int32('\'') {
				libc.Xprintf(tls, __ccgo_ts+385, libc.VaList(bp+8, __ccgo_ts+388))
			} else {
				if int32(*(*uint8)(unsafe.Pointer(cp))) < int32(0x90) {
					libc.Xputchar(tls, int32(*(*uint8)(unsafe.Pointer(cp))))
				} else {
					libc.Xprintf(tls, __ccgo_ts+390, libc.VaList(bp+8, int32(0xf0), int32(0x90), int32(0x91), int32(*(*uint8)(unsafe.Pointer(cp)))))
				}
			}
		}
		goto _2
	_2:
		;
		cp++
		v1 = cp
		if !(*(*uint8)(unsafe.Pointer(v1)) != 0) {
			break
		}
	}
}

func main1(tls *libc.TLS, argc int32, argv uintptr) (r int32) {
	bp := tls.Alloc(1152)
	defer tls.Free(1152)
	var c, cap1, script, sent, w, v1, v2, v7 uint8
	var fp uintptr
	var i, v8 int32
	var _ /* befto at bp+1024 */ [5][2][10]int8
	var _ /* copy at bp+256 */ [256]int8
	var _ /* out at bp+512 */ [256]int8
	var _ /* pword at bp+768 */ [256]int8
	var _ /* word at bp+0 */ [256]int8
	_, _, _, _, _, _, _, _, _, _, _ = c, cap1, fp, i, script, sent, w, v1, v2, v7, v8
	fp = libc.Xstdin
	w = uint8(0)
	sent = uint8(0)
	script = uint8(0)
	*(*[5][2][10]int8)(unsafe.Pointer(bp + 1024)) = [5][2][10]int8{
		0: {
			0: {'h', 'a', 's'},
			1: {-16, -112, -111, -107},
		},
		1: {
			0: {'h', 'a', 'v', 'e'},
			1: {-16, -112, -111, -109},
		},
		2: {
			0: {'u', 's', 'e', 'd'},
			1: {-16, -112, -111, -107, -16, -112, -111, -111},
		},
		3: {
			0: {'u', 'n', 'u', 's', 'e', 'd'},
			1: {-16, -112, -111, -107, -16, -112, -111, -111},
		},
		4: {
			0: {'s', 'u', 'p', 'p', 'o', 's', 'e', 'd'},
			1: {-16, -112, -111, -107, -16, -112, -111, -111},
		},
	}
	if argc < int32(2) {
		libc.Xfprintf(tls, libc.Xstderr, __ccgo_ts+399, libc.VaList(bp+1136, *(*uintptr)(unsafe.Pointer(argv))))
		libc.Xexit(tls, int32(1))
	}
	load_custom(tls, *(*uintptr)(unsafe.Pointer(argv + 1*8)))
	for !(libc.Xfeof(tls, fp) != 0) {
		c = uint8(libc.Xfgetc(tls, fp))
		if libc.BoolInt32(uint32(c)|uint32(32)-uint32('a') < uint32(26)) != 0 || w != 0 && int32(c) == int32('\'') {
			if !(w != 0) {
				v1 = sent
				sent++
				cap1 = libc.BoolUint8(v1 != 0 && libc.BoolInt32(uint32(c)-uint32('A') < uint32(26)) != 0)
			}
			v2 = w
			w++
			(*(*[256]int8)(unsafe.Pointer(bp)))[v2] = int8(c)
			continue
		}
		if libc.Xstrchr(tls, __ccgo_ts+426, int32(c)) != 0 {
			sent = uint8(0)
		}
		if !(w != 0) {
			goto _3
		}
		(*(*[256]int8)(unsafe.Pointer(bp)))[w] = 0
		if script != 0 {
			goto pass
		}
		if !(libc.Xstrcasecmp(tls, bp, __ccgo_ts+431) != 0) {
			i = 0
			for {
				if !(i < int32(5)) {
					break
				}
				if !(libc.Xstrcasecmp(tls, bp+768, bp+1024+uintptr(i)*20) != 0) {
					libc.Xprintf(tls, __ccgo_ts+434, libc.VaList(bp+1136, bp+1024+uintptr(i)*20+1*10))
				}
				goto _4
			_4:
				;
				i++
			}
		}
		aflag = uint8(0)
		libc.Xstrcpy(tls, bp+256, bp)
		if !(parse_prefix(tls, bp+256, bp+512, 0) != 0) {
			goto _5
		}
		shaw_print(tls, bp+512, int32(cap1))
		goto _6
	_5:
		;
		goto pass
	pass:
		;
		libc.Xprintf(tls, __ccgo_ts+385, libc.VaList(bp+1136, bp))
	_6:
		;
		libc.Xmemcpy(tls, bp+768, bp, uint64(256))
	_3:
		;
		goto end
	end:
		;
		w = uint8(0)
		if int32(c) != int32(0xff) {
			libc.Xputchar(tls, int32(c))
		}
		if int32(c) == int32('<') { // pass through HTML/XML tags and Javascript
			i = 0
			for cond := true; cond; cond = int32(c) != int32('>') && !(libc.Xfeof(tls, fp) != 0) {
				v7 = uint8(libc.Xfgetc(tls, fp))
				c = v7
				libc.Xputchar(tls, int32(v7))
				if i < int32(256) {
					v8 = i
					i++
					(*(*[256]int8)(unsafe.Pointer(bp)))[v8] = int8(c)
				}
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+439, uint64(3)) != 0) {
				sent = uint8(0)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+443, uint64(6)) != 0) {
				script = uint8(1)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+450, uint64(7)) != 0) {
				script = uint8(0)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+458, uint64(5)) != 0) {
				script = uint8(1)
			}
			if !(libc.Xstrncasecmp(tls, bp, __ccgo_ts+464, uint64(6)) != 0) {
				script = uint8(0)
			}
		}
	}
	return r
}

func main() {
	libc.Start(main1)
}

var aflag uint8

var dend uintptr

var dict = uintptr(unsafe.Pointer(&prefix))

var prefix [2200000]uint8

var suffix uintptr

var __ccgo_ts = uintptr(unsafe.Pointer(unsafe.StringData(__ccgo_ts1)))

var __ccgo_ts1 = "\x00r\x00%s: Extra space, line %d discarded.\n\x00^$':._\x00%s: Illegal character '%c', line %d discarded.\n\x00eos\x00es\x00hiosuxz\x00ry\x00aeiouf\x00ha\x00cst\x00th\x00e\x00t\x00aeioust'\x00k\x00aceino\x00aeiou\x00m\x00eis\x00z\x00i\x00n\x00eio\x00ess\x00ln\x00cfikmpsv\x00aeiouy\x00aeio\x00hsw\x00cghlsuvz\x00aeiousy\x00cg\x00aou\x00dg\x00call\x00aegiou\x00dlmnprstu\x00arm\x00out\x00und\x00up\x00\xa6\xb0\x00\xa4\xaf\x00\xa9\xbc\x00vw\x00s\x00le\x00\xa6\x00\xa9\xa4\x00\xa4\x00'\x9b\x9f\x00\x9f\xa9\xa5\x00\x9f\xa5\x00\xa9\xad\xbe\x00\xa6\x91\xa6\x00\x96\x97\xa0\xa1\x00\xa8\xa4\x00aeiosu\x00\x92\xa9\xa4\xa6\x00\xa6\xa9\x00\x92\xa4\xa6\x00\x91\x95\x9f\x00\x95\xa4\xaf\x00\xaf\x00\xb0\xb1\xb4\xb5\xb7\xbb\xbf\x00la\x00\xa4\xa5\xae\xaf\x00\xa5\xa9\x92\x92\x00\xa5\xa9\x92*\x92\x00·\x00:.,*\x00%s\x00'\x00%c%c%c%c\x00Usage: %s dictionary-file\n\x00.:?!\x00to\x00@%s \x00div\x00script\x00/script\x00style\x00/style\x00"

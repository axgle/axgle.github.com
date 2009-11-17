<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8g/cgen.c</title>

  <link rel="stylesheet" type="text/css" href="../../../doc/style.css">
  <script type="text/javascript" src="../../../doc/godocs.js"></script>

</head>

<body>

  <script>
    // Catch 'enter' key down events and trigger the search form submission.
    function codesearchKeyDown(event) {
      if (event.which == 13) {
        var form = document.getElementById('codesearch');
        var query = document.getElementById('codesearchQuery');
        form.q.value = "lang:go package:go.googlecode.com " + query.value;
        document.getElementById('codesearch').submit();
}      return true;
}
    // Capture the submission event and construct the query parameter.
    function codeSearchSubmit() {
      var query = document.getElementById('codesearchQuery');
      var form = document.getElementById('codesearch');
      form.q.value = "lang:go package:go.googlecode.com " + query.value;
      return true;
}  </script>

<div id="topnav">
  <table summary="">
    <tr>
      <td id="headerImage">
        <a href="../../../index.html"><img src="../../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
      </td>
      <td>
        <div id="headerDocSetTitle">The Go Programming Language</div>
      </td>
      <td>
        <!-- <table>
          <tr>
            <td>
              <! The input box is outside of the form because we want to add
              a couple of restricts to the query before submitting. If we just
              add the restricts to the text box before submitting, then they
              appear in the box when the user presses 'back'. Thus we use a
              hidden field in the form. However, there's no way to stop the
              non-hidden text box from also submitting a value unless we move
              it outside of the form
              <input type="search" id="codesearchQuery" value="" size="30" onkeydown="return codesearchKeyDown(event);"/>
              <form method="GET" action="http://www.google.com/codesearch" id="codesearch" class="search" onsubmit="return codeSearchSubmit();" style="display:inline;">
                <input type="hidden" name="q" value=""/>
                <input type="submit" value="Code search" />
                <span style="color: red">(TODO: remove for now?)</span>
              </form>
            </td>
          </tr>
          <tr>
            <td>
              <span style="color: gray;">(e.g. &ldquo;pem&rdquo; or &ldquo;xml&rdquo;)</span>
            </td>
          </tr>
        </table> -->
      </td>
    </tr>
  </table>
</div>

<div id="linkList">
  <ul>
    <li class="navhead"><a href="../../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../../doc/install.html">Install Go</a></li>
    <li><a href="../../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../../pkg/index.html">Package documentation</a></li>
    <li><a href="../../index.html">Source files</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Help</li>
    <li>#go-nuts on irc.freenode.net</li>
    <li><a href="http://groups.google.com/group/golang-nuts">Go Nuts mailing list</a></li>
    <li><a href="http://code.google.com/p/go/issues/list">Issue tracker</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Go code search</li>
    <form method="GET" action="http://golang.org/search" class="search">
    <input type="search" name="q" value="" size="25" style="width:80%; max-width:200px" />
    <input type="submit" value="Go" />
    </form>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Last update</li>
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/8g/cgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(rsc):
//	assume CLD?

#include &#34;gg.h&#34;

void
mgen(Node *n, Node *n1, Node *rg)
{
	n1-&gt;ostk = 0;
	n1-&gt;op = OEMPTY;

	if(n-&gt;addable) {
		*n1 = *n;
		n1-&gt;ostk = 0;
		if(n1-&gt;op == OREGISTER || n1-&gt;op == OINDREG)
			reg[n-&gt;val.u.reg]++;
		return;
	}
	if(n-&gt;type-&gt;width &gt; widthptr)
		tempalloc(n1, n-&gt;type);
	else
		regalloc(n1, n-&gt;type, rg);
	cgen(n, n1);
}

void
mfree(Node *n)
{
	if(n-&gt;ostk)
		tempfree(n);
	else if(n-&gt;op == OREGISTER)
		regfree(n);
}

/*
 * generate:
 *	res = n;
 * simplifies and calls gmove.
 *
 * TODO:
 *	sudoaddable
 */
void
cgen(Node *n, Node *res)
{
	Node *nl, *nr, *r, n1, n2, nt, f0, f1;
	Prog *p1, *p2, *p3;
	int a;

	if(debug[&#39;g&#39;]) {
		dump(&#34;\ncgen-n&#34;, n);
		dump(&#34;cgen-res&#34;, res);
	}

	if(n == N || n-&gt;type == T)
		fatal(&#34;cgen: n nil&#34;);
	if(res == N || res-&gt;type == T)
		fatal(&#34;cgen: res nil&#34;);

	while(n-&gt;op == OCONVNOP)
		n = n-&gt;left;

	// function calls on both sides?  introduce temporary
	if(n-&gt;ullman &gt;= UINF &amp;&amp; res-&gt;ullman &gt;= UINF) {
		tempalloc(&amp;n1, n-&gt;type);
		cgen(n, &amp;n1);
		cgen(&amp;n1, res);
		tempfree(&amp;n1);
		return;
	}

	// structs etc get handled specially
	if(isfat(n-&gt;type)) {
		sgen(n, res, n-&gt;type-&gt;width);
		return;
	}

	// update addressability for string, slice
	// can&#39;t do in walk because n-&gt;left-&gt;addable
	// changes if n-&gt;left is an escaping local variable.
	switch(n-&gt;op) {
	case OLEN:
		if(isslice(n-&gt;left-&gt;type) || istype(n-&gt;left-&gt;type, TSTRING))
			n-&gt;addable = n-&gt;left-&gt;addable;
		break;
	case OCAP:
		if(isslice(n-&gt;left-&gt;type))
			n-&gt;addable = n-&gt;left-&gt;addable;
		break;
	}

	// if both are addressable, move
	if(n-&gt;addable &amp;&amp; res-&gt;addable) {
		gmove(n, res);
		return;
	}

	// if both are not addressable, use a temporary.
	if(!n-&gt;addable &amp;&amp; !res-&gt;addable) {
		// could use regalloc here sometimes,
		// but have to check for ullman &gt;= UINF.
		tempalloc(&amp;n1, n-&gt;type);
		cgen(n, &amp;n1);
		cgen(&amp;n1, res);
		tempfree(&amp;n1);
		return;
	}

	// if result is not addressable directly but n is,
	// compute its address and then store via the address.
	if(!res-&gt;addable) {
		igen(res, &amp;n1, N);
		cgen(n, &amp;n1);
		regfree(&amp;n1);
		return;
	}

	// otherwise, the result is addressable but n is not.
	// let&#39;s do some computation.

	// use ullman to pick operand to eval first.
	nl = n-&gt;left;
	nr = n-&gt;right;
	if(nl != N &amp;&amp; nl-&gt;ullman &gt;= UINF)
	if(nr != N &amp;&amp; nr-&gt;ullman &gt;= UINF) {
		// both are hard
		tempalloc(&amp;n1, nl-&gt;type);
		cgen(nl, &amp;n1);
		n2 = *n;
		n2.left = &amp;n1;
		cgen(&amp;n2, res);
		tempfree(&amp;n1);
		return;
	}

	// 64-bit ops are hard on 32-bit machine.
	if(is64(n-&gt;type) || is64(res-&gt;type) || n-&gt;left != N &amp;&amp; is64(n-&gt;left-&gt;type)) {
		switch(n-&gt;op) {
		// math goes to cgen64.
		case OMINUS:
		case OCOM:
		case OADD:
		case OSUB:
		case OMUL:
		case OLSH:
		case ORSH:
		case OAND:
		case OOR:
		case OXOR:
			cgen64(n, res);
			return;
		}
	}

	if(nl != N &amp;&amp; isfloat[n-&gt;type-&gt;etype] &amp;&amp; isfloat[nl-&gt;type-&gt;etype])
		goto flt;

	switch(n-&gt;op) {
	default:
		dump(&#34;cgen&#34;, n);
		fatal(&#34;cgen %O&#34;, n-&gt;op);
		break;

	// these call bgen to get a bool value
	case OOROR:
	case OANDAND:
	case OEQ:
	case ONE:
	case OLT:
	case OLE:
	case OGE:
	case OGT:
	case ONOT:
		p1 = gbranch(AJMP, T);
		p2 = pc;
		gmove(nodbool(1), res);
		p3 = gbranch(AJMP, T);
		patch(p1, pc);
		bgen(n, 1, p2);
		gmove(nodbool(0), res);
		patch(p3, pc);
		return;

	case OPLUS:
		cgen(nl, res);
		return;

	case OMINUS:
	case OCOM:
		a = optoas(n-&gt;op, nl-&gt;type);
		goto uop;

	// symmetric binary
	case OAND:
	case OOR:
	case OXOR:
	case OADD:
	case OMUL:
		a = optoas(n-&gt;op, nl-&gt;type);
		if(a == AIMULB) {
			cgen_bmul(n-&gt;op, nl, nr, res);
			break;
		}
		goto sbop;

	// asymmetric binary
	case OSUB:
		a = optoas(n-&gt;op, nl-&gt;type);
		goto abop;

	case OCONV:
		if(eqtype(n-&gt;type, nl-&gt;type) || noconv(n-&gt;type, nl-&gt;type)) {
			cgen(nl, res);
			break;
		}
		mgen(nl, &amp;n1, res);
		gmove(&amp;n1, res);
		mfree(&amp;n1);
		break;

	case ODOT:
	case ODOTPTR:
	case OINDEX:
	case OIND:
	case ONAME:	// PHEAP or PPARAMREF var
		igen(n, &amp;n1, res);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		break;

	case OLEN:
		if(istype(nl-&gt;type, TMAP) || istype(nl-&gt;type, TCHAN)) {
			// map has len in the first 32-bit word.
			// a zero pointer means zero length
			tempalloc(&amp;n1, types[tptr]);
			cgen(nl, &amp;n1);
			regalloc(&amp;n2, types[tptr], N);
			gmove(&amp;n1, &amp;n2);
			tempfree(&amp;n1);
			n1 = n2;

			nodconst(&amp;n2, types[tptr], 0);
			gins(optoas(OCMP, types[tptr]), &amp;n1, &amp;n2);
			p1 = gbranch(optoas(OEQ, types[tptr]), T);

			n2 = n1;
			n2.op = OINDREG;
			n2.type = types[TINT32];
			gmove(&amp;n2, &amp;n1);

			patch(p1, pc);

			gmove(&amp;n1, res);
			regfree(&amp;n1);
			break;
		}
		if(istype(nl-&gt;type, TSTRING) || isslice(nl-&gt;type)) {
			// both slice and string have len one pointer into the struct.
			igen(nl, &amp;n1, res);
			n1.op = OREGISTER;	// was OINDREG
			regalloc(&amp;n2, types[TUINT32], &amp;n1);
			n1.op = OINDREG;
			n1.type = types[TUINT32];
			n1.xoffset = Array_nel;
			gmove(&amp;n1, &amp;n2);
			gmove(&amp;n2, res);
			regfree(&amp;n1);
			regfree(&amp;n2);
			break;
		}
		fatal(&#34;cgen: OLEN: unknown type %lT&#34;, nl-&gt;type);
		break;

	case OCAP:
		if(istype(nl-&gt;type, TCHAN)) {
			// chan has cap in the second 32-bit word.
			// a zero pointer means zero length
			regalloc(&amp;n1, types[tptr], res);
			cgen(nl, &amp;n1);

			nodconst(&amp;n2, types[tptr], 0);
			gins(optoas(OCMP, types[tptr]), &amp;n1, &amp;n2);
			p1 = gbranch(optoas(OEQ, types[tptr]), T);

			n2 = n1;
			n2.op = OINDREG;
			n2.xoffset = 4;
			n2.type = types[TINT32];
			gmove(&amp;n2, &amp;n1);

			patch(p1, pc);

			gmove(&amp;n1, res);
			regfree(&amp;n1);
			break;
		}
		if(isslice(nl-&gt;type)) {
			igen(nl, &amp;n1, res);
			n1.op = OINDREG;
			n1.type = types[TUINT32];
			n1.xoffset = Array_cap;
			gmove(&amp;n1, res);
			regfree(&amp;n1);
			break;
		}
		fatal(&#34;cgen: OCAP: unknown type %lT&#34;, nl-&gt;type);
		break;

	case OADDR:
		agen(nl, res);
		break;

	case OCALLMETH:
		cgen_callmeth(n, 0);
		cgen_callret(n, res);
		break;

	case OCALLINTER:
		cgen_callinter(n, res, 0);
		cgen_callret(n, res);
		break;

	case OCALLFUNC:
		cgen_call(n, 0);
		cgen_callret(n, res);
		break;

	case OMOD:
	case ODIV:
		cgen_div(n-&gt;op, nl, nr, res);
		break;

	case OLSH:
	case ORSH:
		cgen_shift(n-&gt;op, nl, nr, res);
		break;
	}
	return;

sbop:	// symmetric binary
	if(nl-&gt;ullman &lt; nr-&gt;ullman) {
		r = nl;
		nl = nr;
		nr = r;
	}

abop:	// asymmetric binary
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		tempalloc(&amp;nt, nl-&gt;type);
		cgen(nl, &amp;nt);
		mgen(nr, &amp;n2, N);
		regalloc(&amp;n1, nl-&gt;type, res);
		gmove(&amp;nt, &amp;n1);
		gins(a, &amp;n2, &amp;n1);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		mfree(&amp;n2);
		tempfree(&amp;nt);
	} else {
		regalloc(&amp;n2, nr-&gt;type, res);
		cgen(nr, &amp;n2);
		regalloc(&amp;n1, nl-&gt;type, N);
		cgen(nl, &amp;n1);
		gins(a, &amp;n2, &amp;n1);
		regfree(&amp;n2);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
	}
	return;

uop:	// unary
	tempalloc(&amp;n1, nl-&gt;type);
	cgen(nl, &amp;n1);
	gins(a, N, &amp;n1);
	gmove(&amp;n1, res);
	tempfree(&amp;n1);
	return;

flt:	// floating-point.  387 (not SSE2) to interoperate with 6c
	nodreg(&amp;f0, nl-&gt;type, D_F0);
	nodreg(&amp;f1, n-&gt;type, D_F0+1);
	if(nr != N)
		goto flt2;

	if(n-&gt;op == OMINUS) {
		nr = nodintconst(-1);
		convlit(&amp;nr, n-&gt;type);
		n-&gt;op = OMUL;
		goto flt2;
	}

	// unary
	cgen(nl, &amp;f0);
	if(n-&gt;op != OCONV &amp;&amp; n-&gt;op != OPLUS)
		gins(foptoas(n-&gt;op, n-&gt;type, 0), &amp;f0, &amp;f0);
	gmove(&amp;f0, res);
	return;

flt2:	// binary
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		cgen(nl, &amp;f0);
		if(nr-&gt;addable)
			gins(foptoas(n-&gt;op, n-&gt;type, 0), nr, &amp;f0);
		else {
			cgen(nr, &amp;f0);
			gins(foptoas(n-&gt;op, n-&gt;type, Fpop), &amp;f0, &amp;f1);
		}
	} else {
		cgen(nr, &amp;f0);
		if(nl-&gt;addable)
			gins(foptoas(n-&gt;op, n-&gt;type, Frev), nl, &amp;f0);
		else {
			cgen(nl, &amp;f0);
			gins(foptoas(n-&gt;op, n-&gt;type, Frev|Fpop), &amp;f0, &amp;f1);
		}
	}
	gmove(&amp;f0, res);
	return;
}

/*
 * address gen
 *	res = &amp;n;
 */
void
agen(Node *n, Node *res)
{
	Node *nl, *nr;
	Node n1, n2, n3, n4, tmp;
	Type *t;
	uint32 w;
	uint64 v;
	Prog *p1;

	if(debug[&#39;g&#39;]) {
		dump(&#34;\nagen-res&#34;, res);
		dump(&#34;agen-r&#34;, n);
	}
	if(n == N || n-&gt;type == T || res == N || res-&gt;type == T)
		fatal(&#34;agen&#34;);

	while(n-&gt;op == OCONVNOP)
		n = n-&gt;left;

	// addressable var is easy
	if(n-&gt;addable) {
		if(n-&gt;op == OREGISTER)
			fatal(&#34;agen OREGISTER&#34;);
		regalloc(&amp;n1, types[tptr], res);
		gins(ALEAL, n, &amp;n1);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		return;
	}

	// let&#39;s compute
	nl = n-&gt;left;
	nr = n-&gt;right;

	switch(n-&gt;op) {
	default:
		fatal(&#34;agen %O&#34;, n-&gt;op);

	case OCALLMETH:
		cgen_callmeth(n, 0);
		cgen_aret(n, res);
		break;

	case OCALLINTER:
		cgen_callinter(n, res, 0);
		cgen_aret(n, res);
		break;

	case OCALLFUNC:
		cgen_call(n, 0);
		cgen_aret(n, res);
		break;

	case OINDEX:
		// TODO(rsc): uint64 indices
		w = n-&gt;type-&gt;width;
		if(nr-&gt;addable) {
			agenr(nl, &amp;n3, res);
			if(!isconst(nr, CTINT)) {
				tempalloc(&amp;tmp, types[TINT32]);
				cgen(nr, &amp;tmp);
				regalloc(&amp;n1, tmp.type, N);
				gmove(&amp;tmp, &amp;n1);
				tempfree(&amp;tmp);
			}
		} else if(nl-&gt;addable) {
			if(!isconst(nr, CTINT)) {
				tempalloc(&amp;tmp, types[TINT32]);
				cgen(nr, &amp;tmp);
				regalloc(&amp;n1, tmp.type, N);
				gmove(&amp;tmp, &amp;n1);
				tempfree(&amp;tmp);
			}
			regalloc(&amp;n3, types[tptr], res);
			agen(nl, &amp;n3);
		} else {
			tempalloc(&amp;tmp, types[TINT32]);
			cgen(nr, &amp;tmp);
			nr = &amp;tmp;
			agenr(nl, &amp;n3, res);
			regalloc(&amp;n1, tmp.type, N);
			gins(optoas(OAS, tmp.type), &amp;tmp, &amp;n1);
			tempfree(&amp;tmp);
		}

		// &amp;a is in &amp;n3 (allocated in res)
		// i is in &amp;n1 (if not constant)
		// w is width

		// explicit check for nil if array is large enough
		// that we might derive too big a pointer.
		if(!isslice(nl-&gt;type) &amp;&amp; nl-&gt;type-&gt;width &gt;= unmappedzero) {
			regalloc(&amp;n4, types[tptr], &amp;n3);
			gmove(&amp;n3, &amp;n4);
			n4.op = OINDREG;
			n4.type = types[TUINT8];
			n4.xoffset = 0;
			gins(ATESTB, nodintconst(0), &amp;n4);
			regfree(&amp;n4);
		}

		if(w == 0)
			fatal(&#34;index is zero width&#34;);

		// constant index
		if(isconst(nr, CTINT)) {
			v = mpgetfix(nr-&gt;val.u.xval);
			if(isslice(nl-&gt;type)) {

				if(!debug[&#39;B&#39;]) {
					n1 = n3;
					n1.op = OINDREG;
					n1.type = types[tptr];
					n1.xoffset = Array_nel;
					nodconst(&amp;n2, types[TUINT32], v);
					gins(optoas(OCMP, types[TUINT32]), &amp;n1, &amp;n2);
					p1 = gbranch(optoas(OGT, types[TUINT32]), T);
					//ginscall(throwindex, 0);
					gins(AINT, nodintconst(3), N);
					patch(p1, pc);
				}

				n1 = n3;
				n1.op = OINDREG;
				n1.type = types[tptr];
				n1.xoffset = Array_array;
				gmove(&amp;n1, &amp;n3);
			} else
			if(!debug[&#39;B&#39;]) {
				if(v &lt; 0)
					yyerror(&#34;out of bounds on array&#34;);
				else
				if(v &gt;= nl-&gt;type-&gt;bound)
					yyerror(&#34;out of bounds on array&#34;);
			}

			nodconst(&amp;n2, types[tptr], v*w);
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n3);

			gmove(&amp;n3, res);
			regfree(&amp;n3);
			break;
		}

		// type of the index
		t = types[TUINT32];
		if(issigned[n1.type-&gt;etype])
			t = types[TINT32];

		regalloc(&amp;n2, t, &amp;n1);			// i
		gmove(&amp;n1, &amp;n2);
		regfree(&amp;n1);

		if(!debug[&#39;B&#39;]) {
			// check bounds
			if(isslice(nl-&gt;type)) {
				n1 = n3;
				n1.op = OINDREG;
				n1.type = types[tptr];
				n1.xoffset = Array_nel;
			} else
				nodconst(&amp;n1, types[TUINT32], nl-&gt;type-&gt;bound);
			gins(optoas(OCMP, types[TUINT32]), &amp;n2, &amp;n1);
			p1 = gbranch(optoas(OLT, types[TUINT32]), T);
			//ginscall(throwindex, 0);
			gins(AINT, nodintconst(3), N);
			patch(p1, pc);
		}

		if(isslice(nl-&gt;type)) {
			n1 = n3;
			n1.op = OINDREG;
			n1.type = types[tptr];
			n1.xoffset = Array_array;
			gmove(&amp;n1, &amp;n3);
		}

		if(w == 1 || w == 2 || w == 4 || w == 8) {
			p1 = gins(ALEAL, &amp;n2, &amp;n3);
			p1-&gt;from.scale = w;
			p1-&gt;from.index = p1-&gt;from.type;
			p1-&gt;from.type = p1-&gt;to.type + D_INDIR;
		} else {
			nodconst(&amp;n1, t, w);
			gins(optoas(OMUL, t), &amp;n1, &amp;n2);
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n3);
			gmove(&amp;n3, res);
		}

		gmove(&amp;n3, res);
		regfree(&amp;n2);
		regfree(&amp;n3);
		break;

	case ONAME:
		// should only get here with names in this func.
		if(n-&gt;funcdepth &gt; 0 &amp;&amp; n-&gt;funcdepth != funcdepth) {
			dump(&#34;bad agen&#34;, n);
			fatal(&#34;agen: bad ONAME funcdepth %d != %d&#34;,
				n-&gt;funcdepth, funcdepth);
		}

		// should only get here for heap vars or paramref
		if(!(n-&gt;class &amp; PHEAP) &amp;&amp; n-&gt;class != PPARAMREF) {
			dump(&#34;bad agen&#34;, n);
			fatal(&#34;agen: bad ONAME class %#x&#34;, n-&gt;class);
		}
		cgen(n-&gt;heapaddr, res);
		if(n-&gt;xoffset != 0) {
			nodconst(&amp;n1, types[tptr], n-&gt;xoffset);
			gins(optoas(OADD, types[tptr]), &amp;n1, res);
		}
		break;

	case OIND:
		cgen(nl, res);
		break;

	case ODOT:
		t = nl-&gt;type;
		agen(nl, res);
		if(n-&gt;xoffset != 0) {
			nodconst(&amp;n1, types[tptr], n-&gt;xoffset);
			gins(optoas(OADD, types[tptr]), &amp;n1, res);
		}
		break;

	case ODOTPTR:
		t = nl-&gt;type;
		if(!isptr[t-&gt;etype])
			fatal(&#34;agen: not ptr %N&#34;, n);
		cgen(nl, res);
		if(n-&gt;xoffset != 0) {
			// explicit check for nil if struct is large enough
			// that we might derive too big a pointer.
			if(nl-&gt;type-&gt;type-&gt;width &gt;= unmappedzero) {
				regalloc(&amp;n1, types[tptr], res);
				gmove(res, &amp;n1);
				n1.op = OINDREG;
				n1.type = types[TUINT8];
				n1.xoffset = 0;
				gins(ATESTB, nodintconst(0), &amp;n1);
				regfree(&amp;n1);
			}
			nodconst(&amp;n1, types[tptr], n-&gt;xoffset);
			gins(optoas(OADD, types[tptr]), &amp;n1, res);
		}
		break;
	}
}

/*
 * generate:
 *	newreg = &amp;n;
 *	res = newreg
 *
 * on exit, a has been changed to be *newreg.
 * caller must regfree(a).
 */
void
igen(Node *n, Node *a, Node *res)
{
	Node n1;

	tempalloc(&amp;n1, types[tptr]);
	agen(n, &amp;n1);
	regalloc(a, types[tptr], res);
	gmove(&amp;n1, a);
	tempfree(&amp;n1);
	a-&gt;op = OINDREG;
	a-&gt;type = n-&gt;type;
}

/*
 * generate:
 *	newreg = &amp;n;
 *
 * caller must regfree(a).
 */
void
agenr(Node *n, Node *a, Node *res)
{
	Node n1;

	tempalloc(&amp;n1, types[tptr]);
	agen(n, &amp;n1);
	regalloc(a, types[tptr], res);
	gmove(&amp;n1, a);
	tempfree(&amp;n1);
}

/*
 * branch gen
 *	if(n == true) goto to;
 */
void
bgen(Node *n, int true, Prog *to)
{
	int et, a;
	Node *nl, *nr, *r;
	Node n1, n2, tmp, t1, t2, ax;
	Prog *p1, *p2;

	if(debug[&#39;g&#39;]) {
		dump(&#34;\nbgen&#34;, n);
	}

	if(n == N)
		n = nodbool(1);

	nl = n-&gt;left;
	nr = n-&gt;right;

	if(n-&gt;type == T) {
		convlit(&amp;n, types[TBOOL]);
		if(n-&gt;type == T)
			return;
	}

	et = n-&gt;type-&gt;etype;
	if(et != TBOOL) {
		yyerror(&#34;cgen: bad type %T for %O&#34;, n-&gt;type, n-&gt;op);
		patch(gins(AEND, N, N), to);
		return;
	}
	nl = N;
	nr = N;

	switch(n-&gt;op) {
	default:
	def:
		regalloc(&amp;n1, n-&gt;type, N);
		cgen(n, &amp;n1);
		nodconst(&amp;n2, n-&gt;type, 0);
		gins(optoas(OCMP, n-&gt;type), &amp;n1, &amp;n2);
		a = AJNE;
		if(!true)
			a = AJEQ;
		patch(gbranch(a, n-&gt;type), to);
		regfree(&amp;n1);
		return;

	case OLITERAL:
		// need to ask if it is bool?
		if(!true == !n-&gt;val.u.bval)
			patch(gbranch(AJMP, T), to);
		return;

	case ONAME:
		if(!n-&gt;addable)
			goto def;
		nodconst(&amp;n1, n-&gt;type, 0);
		gins(optoas(OCMP, n-&gt;type), n, &amp;n1);
		a = AJNE;
		if(!true)
			a = AJEQ;
		patch(gbranch(a, n-&gt;type), to);
		return;

	case OANDAND:
		if(!true)
			goto caseor;

	caseand:
		p1 = gbranch(AJMP, T);
		p2 = gbranch(AJMP, T);
		patch(p1, pc);
		bgen(n-&gt;left, !true, p2);
		bgen(n-&gt;right, !true, p2);
		p1 = gbranch(AJMP, T);
		patch(p1, to);
		patch(p2, pc);
		return;

	case OOROR:
		if(!true)
			goto caseand;

	caseor:
		bgen(n-&gt;left, true, to);
		bgen(n-&gt;right, true, to);
		return;

	case OEQ:
	case ONE:
	case OLT:
	case OGT:
	case OLE:
	case OGE:
		nr = n-&gt;right;
		if(nr == N || nr-&gt;type == T)
			return;

	case ONOT:	// unary
		nl = n-&gt;left;
		if(nl == N || nl-&gt;type == T)
			return;
	}

	switch(n-&gt;op) {
	case ONOT:
		bgen(nl, !true, to);
		break;

	case OEQ:
	case ONE:
	case OLT:
	case OGT:
	case OLE:
	case OGE:
		a = n-&gt;op;
		if(!true)
			a = brcom(a);

		// make simplest on right
		if(nl-&gt;op == OLITERAL || nl-&gt;ullman &lt; nr-&gt;ullman) {
			a = brrev(a);
			r = nl;
			nl = nr;
			nr = r;
		}

		if(isslice(nl-&gt;type)) {
			// only valid to cmp darray to literal nil
			if((a != OEQ &amp;&amp; a != ONE) || nr-&gt;op != OLITERAL) {
				yyerror(&#34;illegal array comparison&#34;);
				break;
			}
			a = optoas(a, types[tptr]);
			regalloc(&amp;n1, types[tptr], N);
			agen(nl, &amp;n1);
			n2 = n1;
			n2.op = OINDREG;
			n2.xoffset = Array_array;
			nodconst(&amp;tmp, types[tptr], 0);
			gins(optoas(OCMP, types[tptr]), &amp;n2, &amp;tmp);
			patch(gbranch(a, types[tptr]), to);
			regfree(&amp;n1);
			break;
		}

		if(isinter(nl-&gt;type)) {
			// front end shold only leave cmp to literal nil
			if((a != OEQ &amp;&amp; a != ONE) || nr-&gt;op != OLITERAL) {
				yyerror(&#34;illegal interface comparison&#34;);
				break;
			}
			a = optoas(a, types[tptr]);
			regalloc(&amp;n1, types[tptr], N);
			agen(nl, &amp;n1);
			n2 = n1;
			n2.op = OINDREG;
			n2.xoffset = 0;
			nodconst(&amp;tmp, types[tptr], 0);
			gins(optoas(OCMP, types[tptr]), &amp;n2, &amp;tmp);
			patch(gbranch(a, types[tptr]), to);
			regfree(&amp;n1);
			break;
		}

		if(isfloat[nr-&gt;type-&gt;etype]) {
			nodreg(&amp;tmp, nr-&gt;type, D_F0);
			nodreg(&amp;n2, nr-&gt;type, D_F0 + 1);
			nodreg(&amp;ax, types[TUINT16], D_AX);
			et = simsimtype(nr-&gt;type);
			if(et == TFLOAT64) {
				// easy - do in FPU
				cgen(nr, &amp;tmp);
				cgen(nl, &amp;tmp);
				gins(AFUCOMPP, &amp;tmp, &amp;n2);
			} else {
				// TODO(rsc): The moves back and forth to memory
				// here are for truncating the value to 32 bits.
				// This handles 32-bit comparison but presumably
				// all the other ops have the same problem.
				// We need to figure out what the right general
				// solution is, besides telling people to use float64.
				tempalloc(&amp;t1, types[TFLOAT32]);
				tempalloc(&amp;t2, types[TFLOAT32]);
				cgen(nr, &amp;t1);
				cgen(nl, &amp;t2);
				gmove(&amp;t2, &amp;tmp);
				gins(AFCOMFP, &amp;t1, &amp;tmp);
				tempfree(&amp;t2);
				tempfree(&amp;t1);
			}
			gins(AFSTSW, N, &amp;ax);
			gins(ASAHF, N, N);
			patch(gbranch(optoas(brrev(a), nr-&gt;type), T), to);
			break;
		}

		if(is64(nr-&gt;type)) {
			if(!nl-&gt;addable) {
				tempalloc(&amp;n1, nl-&gt;type);
				cgen(nl, &amp;n1);
				nl = &amp;n1;
			}
			if(!nr-&gt;addable) {
				tempalloc(&amp;n2, nr-&gt;type);
				cgen(nr, &amp;n2);
				nr = &amp;n2;
			}
			cmp64(nl, nr, a, to);
			if(nr == &amp;n2)
				tempfree(&amp;n2);
			if(nl == &amp;n1)
				tempfree(&amp;n1);
			break;
		}

		a = optoas(a, nr-&gt;type);

		if(nr-&gt;ullman &gt;= UINF) {
			tempalloc(&amp;tmp, nr-&gt;type);
			cgen(nr, &amp;tmp);

			tempalloc(&amp;n1, nl-&gt;type);
			cgen(nl, &amp;n1);

			regalloc(&amp;n2, nr-&gt;type, N);
			cgen(&amp;tmp, &amp;n2);

			gins(optoas(OCMP, nr-&gt;type), &amp;n1, &amp;n2);
			patch(gbranch(a, nr-&gt;type), to);
			tempfree(&amp;n1);
			tempfree(&amp;tmp);
			regfree(&amp;n2);
			break;
		}

		tempalloc(&amp;n1, nl-&gt;type);
		cgen(nl, &amp;n1);

		if(smallintconst(nr)) {
			gins(optoas(OCMP, nr-&gt;type), &amp;n1, nr);
			patch(gbranch(a, nr-&gt;type), to);
			tempfree(&amp;n1);
			break;
		}

		tempalloc(&amp;tmp, nr-&gt;type);
		cgen(nr, &amp;tmp);
		regalloc(&amp;n2, nr-&gt;type, N);
		gmove(&amp;tmp, &amp;n2);
		tempfree(&amp;tmp);

		gins(optoas(OCMP, nr-&gt;type), &amp;n1, &amp;n2);
		patch(gbranch(a, nr-&gt;type), to);
		regfree(&amp;n2);
		tempfree(&amp;n1);
		break;
	}
}

/*
 * n is on stack, either local variable
 * or return value from function call.
 * return n&#39;s offset from SP.
 */
int32
stkof(Node *n)
{
	Type *t;
	Iter flist;

	switch(n-&gt;op) {
	case OINDREG:
		return n-&gt;xoffset;

	case OCALLMETH:
	case OCALLINTER:
	case OCALLFUNC:
		t = n-&gt;left-&gt;type;
		if(isptr[t-&gt;etype])
			t = t-&gt;type;

		t = structfirst(&amp;flist, getoutarg(t));
		if(t != T)
			return t-&gt;width;
		break;
	}

	// botch - probably failing to recognize address
	// arithmetic on the above. eg INDEX and DOT
	return -1000;
}

/*
 * struct gen
 *	memmove(&amp;res, &amp;n, w);
 */
void
sgen(Node *n, Node *res, int w)
{
	Node dst, src, tdst, tsrc;
	int32 c, q, odst, osrc;

	if(debug[&#39;g&#39;]) {
		print(&#34;\nsgen w=%d\n&#34;, w);
		dump(&#34;r&#34;, n);
		dump(&#34;res&#34;, res);
	}
	if(w == 0)
		return;
	if(n-&gt;ullman &gt;= UINF &amp;&amp; res-&gt;ullman &gt;= UINF) {
		fatal(&#34;sgen UINF&#34;);
	}

	if(w &lt; 0)
		fatal(&#34;sgen copy %d&#34;, w);

	// offset on the stack
	osrc = stkof(n);
	odst = stkof(res);

	nodreg(&amp;dst, types[tptr], D_DI);
	nodreg(&amp;src, types[tptr], D_SI);

	tempalloc(&amp;tsrc, types[tptr]);
	tempalloc(&amp;tdst, types[tptr]);
	if(!n-&gt;addable)
		agen(n, &amp;tsrc);
	if(!res-&gt;addable)
		agen(res, &amp;tdst);
	if(n-&gt;addable)
		agen(n, &amp;src);
	else
		gmove(&amp;tsrc, &amp;src);
	if(res-&gt;addable)
		agen(res, &amp;dst);
	else
		gmove(&amp;tdst, &amp;dst);
	tempfree(&amp;tdst);
	tempfree(&amp;tsrc);

	c = w % 4;	// bytes
	q = w / 4;	// doublewords

	// if we are copying forward on the stack and
	// the src and dst overlap, then reverse direction
	if(osrc &lt; odst &amp;&amp; odst &lt; osrc+w) {
		// reverse direction
		gins(ASTD, N, N);		// set direction flag
		if(c &gt; 0) {
			gconreg(AADDL, w-1, D_SI);
			gconreg(AADDL, w-1, D_DI);

			gconreg(AMOVL, c, D_CX);
			gins(AREP, N, N);	// repeat
			gins(AMOVSB, N, N);	// MOVB *(SI)-,*(DI)-
		}

		if(q &gt; 0) {
			if(c &gt; 0) {
				gconreg(AADDL, -3, D_SI);
				gconreg(AADDL, -3, D_DI);
			} else {
				gconreg(AADDL, w-4, D_SI);
				gconreg(AADDL, w-4, D_DI);
			}
			gconreg(AMOVL, q, D_CX);
			gins(AREP, N, N);	// repeat
			gins(AMOVSL, N, N);	// MOVL *(SI)-,*(DI)-
		}
		// we leave with the flag clear
		gins(ACLD, N, N);
	} else {
		gins(ACLD, N, N);	// paranoia.  TODO(rsc): remove?
		// normal direction
		if(q &gt;= 4) {
			gconreg(AMOVL, q, D_CX);
			gins(AREP, N, N);	// repeat
			gins(AMOVSL, N, N);	// MOVL *(SI)+,*(DI)+
		} else
		while(q &gt; 0) {
			gins(AMOVSL, N, N);	// MOVL *(SI)+,*(DI)+
			q--;
		}
		while(c &gt; 0) {
			gins(AMOVSB, N, N);	// MOVB *(SI)+,*(DI)+
			c--;
		}
	}
}

</pre>

</div>

<div id="footer">
<p>Except as noted, this content is
   licensed under <a href="http://creativecommons.org/licenses/by/3.0/">
   Creative Commons Attribution 3.0</a>.
</div>

<script type="text/javascript">
var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
var pageTracker = _gat._getTracker("UA-11222381-2");
pageTracker._trackPageview();
</script>
</body>
</html>
<!-- generated at Thu Nov 12 15:42:51 PST 2009 -->

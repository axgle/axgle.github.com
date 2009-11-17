<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5g/cgen.c</title>

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
	<li>Thu Nov 12 15:47:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5g/cgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
		tempname(n1, n-&gt;type);
	else
		regalloc(n1, n-&gt;type, rg);
	cgen(n, n1);
}

void
mfree(Node *n)
{
	if(n-&gt;op == OREGISTER)
		regfree(n);
}

/*
 * generate:
 *	res = n;
 * simplifies and calls gmove.
 */
void
cgen(Node *n, Node *res)
{
	Node *nl, *nr, *r;
	Node n1, n2, n3, f0, f1;
	int a, w;
	Prog *p1, *p2, *p3;
	Addr addr;

	if(debug[&#39;g&#39;]) {
		dump(&#34;\ncgen-n&#34;, n);
		dump(&#34;cgen-res&#34;, res);
	}
	if(n == N || n-&gt;type == T)
		goto ret;

	if(res == N || res-&gt;type == T)
		fatal(&#34;cgen: res nil&#34;);

	while(n-&gt;op == OCONVNOP)
		n = n-&gt;left;

	if(n-&gt;ullman &gt;= UINF) {
		if(n-&gt;op == OINDREG)
			fatal(&#34;cgen: this is going to misscompile&#34;);
		if(res-&gt;ullman &gt;= UINF) {
			tempname(&amp;n1, n-&gt;type);
			cgen(n, &amp;n1);
			cgen(&amp;n1, res);
			goto ret;
		}
	}

	if(isfat(n-&gt;type)) {
		sgen(n, res, n-&gt;type-&gt;width);
		goto ret;
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
		if (is64(n-&gt;type) || is64(res-&gt;type) || n-&gt;op == OREGISTER || res-&gt;op == OREGISTER) {
			gmove(n, res);
		} else {
			regalloc(&amp;n1, n-&gt;type, N);
			gmove(n, &amp;n1);
			cgen(&amp;n1, res);
			regfree(&amp;n1);
		}
		goto ret;
	}

	// if both are not addressable, use a temporary.
	if(!n-&gt;addable &amp;&amp; !res-&gt;addable) {
		// could use regalloc here sometimes,
		// but have to check for ullman &gt;= UINF.
		tempname(&amp;n1, n-&gt;type);
		cgen(n, &amp;n1);
		cgen(&amp;n1, res);
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

	// if n is sudoaddable generate addr and move
	if (!is64(n-&gt;type) &amp;&amp; !is64(res-&gt;type)) {
		a = optoas(OAS, n-&gt;type);
		if(sudoaddable(a, n, &amp;addr, &amp;w)) {
			if (res-&gt;op != OREGISTER) {
				regalloc(&amp;n2, res-&gt;type, N);
				p1 = gins(a, N, &amp;n2);
				p1-&gt;from = addr;
				if(debug[&#39;g&#39;])
					print(&#34;%P [ignore previous line]\n&#34;, p1);
				gmove(&amp;n2, res);
				regfree(&amp;n2);
			} else {
				p1 = gins(a, N, res);
				p1-&gt;from = addr;
				if(debug[&#39;g&#39;])
					print(&#34;%P [ignore previous line]\n&#34;, p1);
			}
			sudoclean();
			goto ret;
		}
	}

	// otherwise, the result is addressable but n is not.
	// let&#39;s do some computation.

	nl = n-&gt;left;
	nr = n-&gt;right;

	if(nl != N &amp;&amp; nl-&gt;ullman &gt;= UINF)
	if(nr != N &amp;&amp; nr-&gt;ullman &gt;= UINF) {
		tempname(&amp;n1, nl-&gt;type);
		cgen(nl, &amp;n1);
		n2 = *n;
		n2.left = &amp;n1;
		cgen(&amp;n2, res);
		goto ret;
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
		fatal(&#34;cgen: unknown op %N&#34;, n);
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
		p1 = gbranch(AB, T);
		p2 = pc;
		gmove(nodbool(1), res);
		p3 = gbranch(AB, T);
		patch(p1, pc);
		bgen(n, 1, p2);
		gmove(nodbool(0), res);
		patch(p3, pc);
		goto ret;

	case OPLUS:
		cgen(nl, res);
		goto ret;

	// unary
	case OCOM:
		a = optoas(OXOR, nl-&gt;type);
		regalloc(&amp;n1, nl-&gt;type, N);
		cgen(nl, &amp;n1);
		nodconst(&amp;n2, nl-&gt;type, -1);
		gins(a, &amp;n2, &amp;n1);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
		goto ret;

	case OMINUS:
		nodconst(&amp;n3, nl-&gt;type, 0);
		regalloc(&amp;n2, nl-&gt;type, res);
		regalloc(&amp;n1, nl-&gt;type, N);
		gmove(&amp;n3, &amp;n2);
		cgen(nl, &amp;n1);
		gins(optoas(OSUB, nl-&gt;type), &amp;n1, &amp;n2);
		gmove(&amp;n2, res);
		regfree(&amp;n1);
		regfree(&amp;n2);
		goto ret;

	// symmetric binary
	case OAND:
	case OOR:
	case OXOR:
	case OADD:
	case OMUL:
		a = optoas(n-&gt;op, nl-&gt;type);
		goto sbop;

	// asymmetric binary
	case OSUB:
		a = optoas(n-&gt;op, nl-&gt;type);
		goto abop;

	case OLSH:
	case ORSH:
		cgen_shift(n-&gt;op, nl, nr, res);
		break;

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
			regalloc(&amp;n1, types[tptr], res);
			cgen(nl, &amp;n1);

			nodconst(&amp;n2, types[tptr], 0);
			regalloc(&amp;n3, n2.type, N);
			gmove(&amp;n2, &amp;n3);
			gcmp(optoas(OCMP, types[tptr]), &amp;n1, &amp;n3);
			regfree(&amp;n3);
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
			regalloc(&amp;n3, n2.type, N);
			gmove(&amp;n2, &amp;n3);
			gcmp(optoas(OCMP, types[tptr]), &amp;n1, &amp;n3);
			regfree(&amp;n3);
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
			regalloc(&amp;n1, types[tptr], res);
			agen(nl, &amp;n1);
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
		a = optoas(n-&gt;op, nl-&gt;type);
		goto abop;
	}
	goto ret;

sbop:	// symmetric binary
	if(nl-&gt;ullman &lt; nr-&gt;ullman) {
		r = nl;
		nl = nr;
		nr = r;
	}

abop:	// asymmetric binary
	// TODO(kaib): use fewer registers here.
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
		regalloc(&amp;n2, nr-&gt;type, N);
		cgen(nr, &amp;n2);
	} else {
		regalloc(&amp;n2, nr-&gt;type, N);
		cgen(nr, &amp;n2);
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);
	}
	gins(a, &amp;n2, &amp;n1);
	gmove(&amp;n1, res);
	regfree(&amp;n1);
	regfree(&amp;n2);
	goto ret;

flt:	// floating-point.
	regalloc(&amp;f0, nl-&gt;type, res);
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
		gins(optoas(n-&gt;op, n-&gt;type), &amp;f0, &amp;f0);
	gmove(&amp;f0, res);
	regfree(&amp;f0);
	goto ret;

flt2:	// binary
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		cgen(nl, &amp;f0);
		regalloc(&amp;f1, n-&gt;type, N);
		gmove(&amp;f0, &amp;f1);
		cgen(nr, &amp;f0);
		gins(optoas(n-&gt;op, n-&gt;type), &amp;f0, &amp;f1);
	} else {
		cgen(nr, &amp;f0);
		regalloc(&amp;f1, n-&gt;type, N);
		cgen(nl, &amp;f1);
		gins(optoas(n-&gt;op, n-&gt;type), &amp;f0, &amp;f1);
	}
	gmove(&amp;f1, res);
	regfree(&amp;f0);
	regfree(&amp;f1);
	goto ret;

ret:
	;
}

/*
 * generate:
 *	res = &amp;n;
 */
void
agen(Node *n, Node *res)
{
	Node *nl, *nr;
	Node n1, n2, n3, n4, n5, tmp;
	Prog *p1;
	uint32 w;
	uint64 v;
	Type *t;

	if(debug[&#39;g&#39;]) {
		dump(&#34;\nagen-res&#34;, res);
		dump(&#34;agen-r&#34;, n);
	}
	if(n == N || n-&gt;type == T || res == N || res-&gt;type == T)
		fatal(&#34;agen&#34;);

	while(n-&gt;op == OCONVNOP)
		n = n-&gt;left;

	if(n-&gt;addable) {
		memset(&amp;n1, 0, sizeof n1);
		n1.op = OADDR;
		n1.left = n;
		regalloc(&amp;n2, types[tptr], res);
		gins(AMOVW, &amp;n1, &amp;n2);
		gmove(&amp;n2, res);
		regfree(&amp;n2);
		goto ret;
	}

	nl = n-&gt;left;
	nr = n-&gt;right;

	switch(n-&gt;op) {
	default:
		fatal(&#34;agen: unknown op %N&#34;, n);
		break;

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
				tempname(&amp;tmp, types[TINT32]);
				cgen(nr, &amp;tmp);
				regalloc(&amp;n1, tmp.type, N);
				gmove(&amp;tmp, &amp;n1);
			}
		} else if(nl-&gt;addable) {
			if(!isconst(nr, CTINT)) {
				tempname(&amp;tmp, types[TINT32]);
				cgen(nr, &amp;tmp);
				regalloc(&amp;n1, tmp.type, N);
				gmove(&amp;tmp, &amp;n1);
			}
			regalloc(&amp;n3, types[tptr], res);
			agen(nl, &amp;n3);
		} else {
			tempname(&amp;tmp, types[TINT32]);
			cgen(nr, &amp;tmp);
			nr = &amp;tmp;
			agenr(nl, &amp;n3, res);
			regalloc(&amp;n1, tmp.type, N);
			gins(optoas(OAS, tmp.type), &amp;tmp, &amp;n1);
		}

		// &amp;a is in &amp;n3 (allocated in res)
		// i is in &amp;n1 (if not constant)
		// w is width

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
					regalloc(&amp;n4, n1.type, N);
					cgen(&amp;n1, &amp;n4);
					nodconst(&amp;n2, types[TUINT32], v);
					regalloc(&amp;n5, n2.type, N);
					gmove(&amp;n2, &amp;n5);
					gcmp(optoas(OCMP, types[TUINT32]), &amp;n4, &amp;n5);
					regfree(&amp;n4);
					regfree(&amp;n5);
					p1 = gbranch(optoas(OGT, types[TUINT32]), T);
					ginscall(throwindex, 0);
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
			regalloc(&amp;n4, n2.type, N);
			gmove(&amp;n2, &amp;n4);
			gins(optoas(OADD, types[tptr]), &amp;n4, &amp;n3);
			regfree(&amp;n4);

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
			regalloc(&amp;n4, types[TUINT32], N);
			if(isslice(nl-&gt;type)) {
				n1 = n3;
				n1.op = OINDREG;
				n1.type = types[tptr];
				n1.xoffset = Array_nel;
				cgen(&amp;n1, &amp;n4);
			} else {
				nodconst(&amp;n1, types[TUINT32], nl-&gt;type-&gt;bound);
				gmove(&amp;n1, &amp;n4);
			}
			gcmp(optoas(OCMP, types[TUINT32]), &amp;n2, &amp;n4);
			regfree(&amp;n4);
			p1 = gbranch(optoas(OLT, types[TUINT32]), T);
			ginscall(throwindex, 0);
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
			memset(&amp;n4, 0, sizeof n4);
			n4.op = OADDR;
			n4.left = &amp;n2;
			cgen(&amp;n4, &amp;n3);
			if (w == 1)
				gins(AADD, &amp;n2, &amp;n3);
			else if(w == 2)
				gshift(AADD, &amp;n2, SHIFT_LL, 1, &amp;n3);
			else if(w == 4)
				gshift(AADD, &amp;n2, SHIFT_LL, 2, &amp;n3);
			else if(w == 8)
				gshift(AADD, &amp;n2, SHIFT_LL, 3, &amp;n3);	
		} else {
			regalloc(&amp;n4, t, N);
			nodconst(&amp;n1, t, w);
			gmove(&amp;n1, &amp;n4);
			gins(optoas(OMUL, t), &amp;n4, &amp;n2);
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n3);
			regfree(&amp;n4);
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
			nodconst(&amp;n1, types[TINT32], n-&gt;xoffset);
			regalloc(&amp;n2, n1.type, N);
			regalloc(&amp;n3, types[TINT32], N);
			gmove(&amp;n1, &amp;n2);
			gmove(res, &amp;n3);
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n3);
			gmove(&amp;n3, res);
			regfree(&amp;n2);
			regfree(&amp;n3);
		}
		break;

	case OIND:
		cgen(nl, res);
		break;

	case ODOT:
		agen(nl, res);
		if(n-&gt;xoffset != 0) {
			nodconst(&amp;n1, types[TINT32], n-&gt;xoffset);
			regalloc(&amp;n2, n1.type, N);
			regalloc(&amp;n3, types[TINT32], N);
			gmove(&amp;n1, &amp;n2);
			gmove(res, &amp;n3);
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n3);
			gmove(&amp;n3, res);
			regfree(&amp;n2);
			regfree(&amp;n3);
		}
		break;

	case ODOTPTR:
		cgen(nl, res);
		if(n-&gt;xoffset != 0) {
			// explicit check for nil if struct is large enough
			// that we might derive too big a pointer.
			if(nl-&gt;type-&gt;type-&gt;width &gt;= unmappedzero) {
				regalloc(&amp;n1, types[tptr], N);
				gmove(res, &amp;n1);
				p1 = gins(AMOVW, &amp;n1, &amp;n1);
				p1-&gt;from.type = D_OREG;
				p1-&gt;from.offset = 0;
				regfree(&amp;n1);
			}
			nodconst(&amp;n1, types[TINT32], n-&gt;xoffset);
			regalloc(&amp;n2, n1.type, N);
			regalloc(&amp;n3, types[tptr], N);
			gmove(&amp;n1, &amp;n2);
			gmove(res, &amp;n3);
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n3);
			gmove(&amp;n3, res);
			regfree(&amp;n2);
			regfree(&amp;n3);
		}
		break;
	}

ret:
	;
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
	regalloc(a, types[tptr], res);
	agen(n, a);
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

	tempname(&amp;n1, types[tptr]);
	agen(n, &amp;n1);
	regalloc(a, types[tptr], res);
	gmove(&amp;n1, a);
}

/*
 * generate:
 *	if(n == true) goto to;
 */
void
bgen(Node *n, int true, Prog *to)
{
	int et, a;
	Node *nl, *nr, *r;
	Node n1, n2, n3, n4, tmp;
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
			goto ret;
	}

	et = n-&gt;type-&gt;etype;
	if(et != TBOOL) {
		yyerror(&#34;cgen: bad type %T for %O&#34;, n-&gt;type, n-&gt;op);
		patch(gins(AEND, N, N), to);
		goto ret;
	}
	nl = N;
	nr = N;

	switch(n-&gt;op) {
	default:
	def:
		regalloc(&amp;n1, n-&gt;type, N);
		cgen(n, &amp;n1);
		nodconst(&amp;n2, n-&gt;type, 0);
		regalloc(&amp;n3, n-&gt;type, N);
		gmove(&amp;n2, &amp;n3);
		gcmp(optoas(OCMP, n-&gt;type), &amp;n1, &amp;n3);
		a = ABNE;
		if(!true)
			a = ABEQ;
		patch(gbranch(a, n-&gt;type), to);
		regfree(&amp;n1);
		regfree(&amp;n3);
		goto ret;

	case OLITERAL:
		// need to ask if it is bool?
		if(!true == !n-&gt;val.u.bval)
			patch(gbranch(AB, T), to);
		goto ret;

	case ONAME:
		if(n-&gt;addable == 0)
			goto def;
		nodconst(&amp;n1, n-&gt;type, 0);
		regalloc(&amp;n2, n-&gt;type, N);
		regalloc(&amp;n3, n-&gt;type, N);
		gmove(&amp;n1, &amp;n2);
		cgen(n, &amp;n3);
		gcmp(optoas(OCMP, n-&gt;type), &amp;n2, &amp;n3);
		a = ABNE;
		if(!true)
			a = ABEQ;
		patch(gbranch(a, n-&gt;type), to);
		regfree(&amp;n2);
		regfree(&amp;n3);
		goto ret;

	case OANDAND:
		if(!true)
			goto caseor;

	caseand:
		p1 = gbranch(AB, T);
		p2 = gbranch(AB, T);
		patch(p1, pc);
		bgen(n-&gt;left, !true, p2);
		bgen(n-&gt;right, !true, p2);
		p1 = gbranch(AB, T);
		patch(p1, to);
		patch(p2, pc);
		goto ret;

	case OOROR:
		if(!true)
			goto caseand;

	caseor:
		bgen(n-&gt;left, true, to);
		bgen(n-&gt;right, true, to);
		goto ret;

	case OEQ:
	case ONE:
	case OLT:
	case OGT:
	case OLE:
	case OGE:
		nr = n-&gt;right;
		if(nr == N || nr-&gt;type == T)
			goto ret;

	case ONOT:	// unary
		nl = n-&gt;left;
		if(nl == N || nl-&gt;type == T)
			goto ret;
	}

	switch(n-&gt;op) {

	case ONOT:
		bgen(nl, !true, to);
		goto ret;

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
			regalloc(&amp;n3, types[tptr], N);
			regalloc(&amp;n4, types[tptr], N);
			agen(nl, &amp;n1);
			n2 = n1;
			n2.op = OINDREG;
			n2.xoffset = Array_array;
			gmove(&amp;n2, &amp;n4);
			nodconst(&amp;tmp, types[tptr], 0);
			gmove(&amp;tmp, &amp;n3);
			gcmp(optoas(OCMP, types[tptr]), &amp;n4, &amp;n3);
			patch(gbranch(a, types[tptr]), to);
			regfree(&amp;n4);
			regfree(&amp;n3);
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
			regalloc(&amp;n3, types[tptr], N);
			regalloc(&amp;n4, types[tptr], N);
			agen(nl, &amp;n1);
			n2 = n1;
			n2.op = OINDREG;
			n2.xoffset = 0;
			gmove(&amp;n2, &amp;n4);
			nodconst(&amp;tmp, types[tptr], 0);
			gmove(&amp;tmp, &amp;n3);
			gcmp(optoas(OCMP, types[tptr]), &amp;n4, &amp;n3);
			patch(gbranch(a, types[tptr]), to);
			regfree(&amp;n1);
			regfree(&amp;n3);
			regfree(&amp;n4);
			break;
		}

		if(is64(nr-&gt;type)) {
			if(!nl-&gt;addable) {
				tempname(&amp;n1, nl-&gt;type);
				cgen(nl, &amp;n1);
				nl = &amp;n1;
			}
			if(!nr-&gt;addable) {
				tempname(&amp;n2, nr-&gt;type);
				cgen(nr, &amp;n2);
				nr = &amp;n2;
			}
			cmp64(nl, nr, a, to);
			break;
		}

		a = optoas(a, nr-&gt;type);

		if(nr-&gt;ullman &gt;= UINF) {
			regalloc(&amp;n1, nr-&gt;type, N);
			cgen(nr, &amp;n1);

			tempname(&amp;tmp, nr-&gt;type);
			gmove(&amp;n1, &amp;tmp);
			regfree(&amp;n1);

			regalloc(&amp;n1, nl-&gt;type, N);
			cgen(nl, &amp;n1);

			regalloc(&amp;n2, nr-&gt;type, N);
			cgen(&amp;tmp, &amp;n2);

			gcmp(optoas(OCMP, nr-&gt;type), &amp;n1, &amp;n2);
			patch(gbranch(a, nr-&gt;type), to);

			regfree(&amp;n1);
			regfree(&amp;n2);
			break;
		}

		regalloc(&amp;n1, nl-&gt;type, N);
		cgen(nl, &amp;n1);

		regalloc(&amp;n2, nr-&gt;type, N);
		cgen(nr, &amp;n2);

		gcmp(optoas(OCMP, nr-&gt;type), &amp;n1, &amp;n2);
		patch(gbranch(a, nr-&gt;type), to);

		regfree(&amp;n1);
		regfree(&amp;n2);
		break;
	}
	goto ret;

ret:
	;
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
 * block copy:
 *	memmove(&amp;res, &amp;n, w);
 * NB: character copy assumed little endian architecture
 */
void
sgen(Node *n, Node *res, int32 w)
{
	Node dst, src, tmp, nend;
	int32 c, q, odst, osrc;
	Prog *p, *ploop;

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

	if(osrc % 4 != 0 || odst %4 != 0)
		fatal(&#34;sgen: non word(4) aligned offset src %d or dst %d&#34;, osrc, odst);

	regalloc(&amp;dst, types[tptr], res);

	if(n-&gt;ullman &gt;= res-&gt;ullman) {
		agen(n, &amp;dst);	// temporarily use dst
		regalloc(&amp;src, types[tptr], N);
		gins(AMOVW, &amp;dst, &amp;src);
		agen(res, &amp;dst);
	} else {
		agen(res, &amp;dst);
		regalloc(&amp;src, types[tptr], N);
		agen(n, &amp;src);
	}

	regalloc(&amp;tmp, types[TUINT32], N);

	c = w % 4;	// bytes
	q = w / 4;	// quads

	// if we are copying forward on the stack and
	// the src and dst overlap, then reverse direction
	if(osrc &lt; odst &amp;&amp; odst &lt; osrc+w) {
		if(c != 0)
			fatal(&#34;sgen: reverse character copy not implemented&#34;);
		if(q &gt;= 4) {
			regalloc(&amp;nend, types[TUINT32], N);
			// set up end marker to 4 bytes before source
			p = gins(AMOVW, &amp;src, &amp;nend);
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = -4;

			// move src and dest to the end of block
			p = gins(AMOVW, &amp;src, &amp;src);
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = (q-1)*4;

			p = gins(AMOVW, &amp;dst, &amp;dst);
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = (q-1)*4;

			p = gins(AMOVW, &amp;src, &amp;tmp);
			p-&gt;from.type = D_OREG;
			p-&gt;from.offset = -4;
			p-&gt;scond |= C_PBIT;
			ploop = p;

			p = gins(AMOVW, &amp;tmp, &amp;dst);
			p-&gt;to.type = D_OREG;
			p-&gt;to.offset = -4;
			p-&gt;scond |= C_PBIT;

			p = gins(ACMP, &amp;src, N);
			raddr(&amp;nend, p);

			patch(gbranch(ABNE, T), ploop);

 			regfree(&amp;nend);
		} else {
			// move src and dest to the end of block
			p = gins(AMOVW, &amp;src, &amp;src);
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = (q-1)*4;

			p = gins(AMOVW, &amp;dst, &amp;dst);
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = (q-1)*4;

			while(q &gt; 0) {
				p = gins(AMOVW, &amp;src, &amp;tmp);
				p-&gt;from.type = D_OREG;
				p-&gt;from.offset = -4;
 				p-&gt;scond |= C_PBIT;

				p = gins(AMOVW, &amp;tmp, &amp;dst);
				p-&gt;to.type = D_OREG;
				p-&gt;to.offset = -4;
 				p-&gt;scond |= C_PBIT;

				q--;
			}
		}
	} else {
		// normal direction
		if(q &gt;= 4) {
			regalloc(&amp;nend, types[TUINT32], N);
			p = gins(AMOVW, &amp;src, &amp;nend);
			p-&gt;from.type = D_CONST;
			p-&gt;from.offset = q*4;

			p = gins(AMOVW, &amp;src, &amp;tmp);
			p-&gt;from.type = D_OREG;
			p-&gt;from.offset = 4;
			p-&gt;scond |= C_PBIT;
			ploop = p;

			p = gins(AMOVW, &amp;tmp, &amp;dst);
			p-&gt;to.type = D_OREG;
			p-&gt;to.offset = 4;
			p-&gt;scond |= C_PBIT;

			p = gins(ACMP, &amp;src, N);
			raddr(&amp;nend, p);

			patch(gbranch(ABNE, T), ploop);

 			regfree(&amp;nend);
		} else
		while(q &gt; 0) {
			p = gins(AMOVW, &amp;src, &amp;tmp);
			p-&gt;from.type = D_OREG;
			p-&gt;from.offset = 4;
 			p-&gt;scond |= C_PBIT;

			p = gins(AMOVW, &amp;tmp, &amp;dst);
			p-&gt;to.type = D_OREG;
			p-&gt;to.offset = 4;
 			p-&gt;scond |= C_PBIT;

			q--;
		}

		if (c != 0) {
			//	MOVW	(src), tmp
			p = gins(AMOVW, &amp;src, &amp;tmp);
			p-&gt;from.type = D_OREG;

			//	MOVW	tmp&lt;&lt;((4-c)*8),src
			gshift(AMOVW, &amp;tmp, SHIFT_LL, ((4-c)*8), &amp;src);

			//	MOVW	src&gt;&gt;((4-c)*8),src
			gshift(AMOVW, &amp;src, SHIFT_LR, ((4-c)*8), &amp;src);

			//	MOVW	(dst), tmp
			p = gins(AMOVW, &amp;dst, &amp;tmp);
			p-&gt;from.type = D_OREG;

			//	MOVW	tmp&gt;&gt;(c*8),tmp
			gshift(AMOVW, &amp;tmp, SHIFT_LR, (c*8), &amp;tmp);

			//	MOVW	tmp&lt;&lt;(c*8),tmp
			gshift(AMOVW, &amp;tmp, SHIFT_LL, c*8, &amp;tmp);

			//	ORR		src, tmp
			gins(AORR, &amp;src, &amp;tmp);

			//	MOVW	tmp, (dst)
			p = gins(AMOVW, &amp;tmp, &amp;dst);
			p-&gt;to.type = D_OREG;
		}
	}
 	regfree(&amp;dst);
	regfree(&amp;src);
	regfree(&amp;tmp);
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

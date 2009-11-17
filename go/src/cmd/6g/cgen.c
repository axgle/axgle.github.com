<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/6g/cgen.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/6g/cgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include &#34;gg.h&#34;

/*
 * generate:
 *	res = n;
 * simplifies and calls gmove.
 */
void
cgen(Node *n, Node *res)
{
	Node *nl, *nr, *r;
	Node n1, n2;
	int a, f;
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

	// inline slices
	if(cgen_inline(n, res))
		goto ret;

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

	if(!res-&gt;addable) {
		if(n-&gt;ullman &gt; res-&gt;ullman) {
			regalloc(&amp;n1, n-&gt;type, res);
			cgen(n, &amp;n1);
			if(n1.ullman &gt; res-&gt;ullman) {
				dump(&#34;n1&#34;, &amp;n1);
				dump(&#34;res&#34;, res);
				fatal(&#34;loop in cgen&#34;);
			}
			cgen(&amp;n1, res);
			regfree(&amp;n1);
			goto ret;
		}

		if(res-&gt;ullman &gt;= UINF)
			goto gen;

		f = 1;	// gen thru register
		switch(n-&gt;op) {
		case OLITERAL:
			if(smallintconst(n))
				f = 0;
			break;
		case OREGISTER:
			f = 0;
			break;
		}

		a = optoas(OAS, res-&gt;type);
		if(sudoaddable(a, res, &amp;addr)) {
			if(f) {
				regalloc(&amp;n2, res-&gt;type, N);
				cgen(n, &amp;n2);
				p1 = gins(a, &amp;n2, N);
				regfree(&amp;n2);
			} else
				p1 = gins(a, n, N);
			p1-&gt;to = addr;
			if(debug[&#39;g&#39;])
				print(&#34;%P [ignore previous line]\n&#34;, p1);
			sudoclean();
			goto ret;
		}

	gen:
		igen(res, &amp;n1, N);
		cgen(n, &amp;n1);
		regfree(&amp;n1);
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

	if(n-&gt;addable) {
		gmove(n, res);
		goto ret;
	}

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

	a = optoas(OAS, n-&gt;type);
	if(sudoaddable(a, n, &amp;addr)) {
		if(res-&gt;op == OREGISTER) {
			p1 = gins(a, N, res);
			p1-&gt;from = addr;
		} else {
			regalloc(&amp;n2, n-&gt;type, N);
			p1 = gins(a, N, &amp;n2);
			p1-&gt;from = addr;
			gins(a, &amp;n2, res);
			regfree(&amp;n2);
		}
		sudoclean();
		goto ret;
	}

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
		p1 = gbranch(AJMP, T);
		p2 = pc;
		gmove(nodbool(1), res);
		p3 = gbranch(AJMP, T);
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
		if(isfloat[nl-&gt;type-&gt;etype]) {
			nr = nodintconst(-1);
			convlit(&amp;nr, n-&gt;type);
			a = optoas(OMUL, nl-&gt;type);
			goto sbop;
		}
		a = optoas(n-&gt;op, nl-&gt;type);
		goto uop;

	// symmetric binary
	case OAND:
	case OOR:
	case OXOR:
	case OADD:
	case OMUL:
		a = optoas(n-&gt;op, nl-&gt;type);
		if(a != AIMULB)
			goto sbop;
		cgen_bmul(n-&gt;op, nl, nr, res);
		break;

	// asymmetric binary
	case OSUB:
		a = optoas(n-&gt;op, nl-&gt;type);
		goto abop;

	case OCONV:
		regalloc(&amp;n1, nl-&gt;type, res);
		regalloc(&amp;n2, n-&gt;type, &amp;n1);
		cgen(nl, &amp;n1);
		// if we do the conversion n1 -&gt; n2 here
		// reusing the register, then gmove won&#39;t
		// have to allocate its own register.
		gmove(&amp;n1, &amp;n2);
		gmove(&amp;n2, res);
		regfree(&amp;n2);
		regfree(&amp;n1);
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
			// map and chan have len in the first 32-bit word.
			// a zero pointer means zero length
			regalloc(&amp;n1, types[tptr], res);
			cgen(nl, &amp;n1);

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
			// a zero pointer means zero length
			regalloc(&amp;n1, types[tptr], res);
			agen(nl, &amp;n1);
			n1.op = OINDREG;
			n1.type = types[TUINT32];
			n1.xoffset = Array_nel;
			gmove(&amp;n1, res);
			regfree(&amp;n1);
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
		if(isfloat[n-&gt;type-&gt;etype]) {
			a = optoas(n-&gt;op, nl-&gt;type);
			goto abop;
		}
		cgen_div(n-&gt;op, nl, nr, res);
		break;

	case OLSH:
	case ORSH:
		cgen_shift(n-&gt;op, nl, nr, res);
		break;
	}
	goto ret;

sbop:	// symmetric binary
	if(nl-&gt;ullman &lt; nr-&gt;ullman) {
		r = nl;
		nl = nr;
		nr = r;
	}

abop:	// asymmetric binary
	if(nl-&gt;ullman &gt;= nr-&gt;ullman) {
		regalloc(&amp;n1, nl-&gt;type, res);
		cgen(nl, &amp;n1);

		if(sudoaddable(a, nr, &amp;addr)) {
			p1 = gins(a, N, &amp;n1);
			p1-&gt;from = addr;
			gmove(&amp;n1, res);
			sudoclean();
			regfree(&amp;n1);
			goto ret;
		}
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

uop:	// unary
	regalloc(&amp;n1, nl-&gt;type, res);
	cgen(nl, &amp;n1);
	gins(a, N, &amp;n1);
	gmove(&amp;n1, res);
	regfree(&amp;n1);
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
	Node n1, n2, n3, tmp, n4;
	Prog *p1;
	uint32 w;
	uint64 v;
	Type *t;

	if(debug[&#39;g&#39;]) {
		dump(&#34;\nagen-res&#34;, res);
		dump(&#34;agen-r&#34;, n);
	}
	if(n == N || n-&gt;type == T)
		return;

	if(!isptr[res-&gt;type-&gt;etype])
		fatal(&#34;agen: not tptr: %T&#34;, res-&gt;type);

	while(n-&gt;op == OCONVNOP)
		n = n-&gt;left;

	if(n-&gt;addable) {
		regalloc(&amp;n1, types[tptr], res);
		gins(ALEAQ, n, &amp;n1);
		gmove(&amp;n1, res);
		regfree(&amp;n1);
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
		w = n-&gt;type-&gt;width;
		if(nr-&gt;addable)
			goto irad;
		if(nl-&gt;addable) {
			if(!isconst(nr, CTINT)) {
				regalloc(&amp;n1, nr-&gt;type, N);
				cgen(nr, &amp;n1);
			}
			regalloc(&amp;n3, types[tptr], res);
			agen(nl, &amp;n3);
			goto index;
		}
		tempname(&amp;tmp, nr-&gt;type);
		cgen(nr, &amp;tmp);
		nr = &amp;tmp;

	irad:
		regalloc(&amp;n3, types[tptr], res);
		agen(nl, &amp;n3);
		if(!isconst(nr, CTINT)) {
			regalloc(&amp;n1, nr-&gt;type, N);
			cgen(nr, &amp;n1);
		}
		goto index;

	index:
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
					nodconst(&amp;n2, types[TUINT64], v);
					gins(optoas(OCMP, types[TUINT32]), &amp;n1, &amp;n2);
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
			gins(optoas(OADD, types[tptr]), &amp;n2, &amp;n3);

			gmove(&amp;n3, res);
			regfree(&amp;n3);
			break;
		}

		// type of the index
		t = types[TUINT64];
		if(issigned[n1.type-&gt;etype])
			t = types[TINT64];

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
				nodconst(&amp;n1, types[TUINT64], nl-&gt;type-&gt;bound);
			gins(optoas(OCMP, types[TUINT32]), &amp;n2, &amp;n1);
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
			p1 = gins(ALEAQ, &amp;n2, &amp;n3);
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
			nodconst(&amp;n1, types[TINT64], n-&gt;xoffset);
			gins(optoas(OADD, types[tptr]), &amp;n1, res);
		}
		break;

	case OIND:
		cgen(nl, res);
		break;

	case ODOT:
		agen(nl, res);
		if(n-&gt;xoffset != 0) {
			nodconst(&amp;n1, types[TINT64], n-&gt;xoffset);
			gins(optoas(OADD, types[tptr]), &amp;n1, res);
		}
		break;

	case ODOTPTR:
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
			nodconst(&amp;n1, types[TINT64], n-&gt;xoffset);
			gins(optoas(OADD, types[tptr]), &amp;n1, res);
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
 *	if(n == true) goto to;
 */
void
bgen(Node *n, int true, Prog *to)
{
	int et, a;
	Node *nl, *nr, *r;
	Node n1, n2, tmp;
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
		gins(optoas(OCMP, n-&gt;type), &amp;n1, &amp;n2);
		a = AJNE;
		if(!true)
			a = AJEQ;
		patch(gbranch(a, n-&gt;type), to);
		regfree(&amp;n1);
		goto ret;

	case OLITERAL:
		// need to ask if it is bool?
		if(!true == !n-&gt;val.u.bval)
			patch(gbranch(AJMP, T), to);
		goto ret;

	case ONAME:
		if(n-&gt;addable == 0)
			goto def;
		nodconst(&amp;n1, n-&gt;type, 0);
		gins(optoas(OCMP, n-&gt;type), n, &amp;n1);
		a = AJNE;
		if(!true)
			a = AJEQ;
		patch(gbranch(a, n-&gt;type), to);
		goto ret;

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
		break;
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

		a = optoas(a, nr-&gt;type);

		if(nr-&gt;ullman &gt;= UINF) {
			regalloc(&amp;n1, nr-&gt;type, N);
			cgen(nr, &amp;n1);

			tempname(&amp;tmp, nr-&gt;type);
			gmove(&amp;n1, &amp;tmp);
			regfree(&amp;n1);

			regalloc(&amp;n1, nl-&gt;type, N);
			cgen(nl, &amp;n1);

			regalloc(&amp;n2, nr-&gt;type, &amp;n2);
			cgen(&amp;tmp, &amp;n2);

			gins(optoas(OCMP, nr-&gt;type), &amp;n1, &amp;n2);
			patch(gbranch(a, nr-&gt;type), to);

			regfree(&amp;n1);
			regfree(&amp;n2);
			break;
		}

		regalloc(&amp;n1, nl-&gt;type, N);
		cgen(nl, &amp;n1);

		if(smallintconst(nr)) {
			gins(optoas(OCMP, nr-&gt;type), &amp;n1, nr);
			patch(gbranch(a, nr-&gt;type), to);
			regfree(&amp;n1);
			break;
		}

		regalloc(&amp;n2, nr-&gt;type, N);
		cgen(nr, &amp;n2);

		gins(optoas(OCMP, nr-&gt;type), &amp;n1, &amp;n2);
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
 *	memmove(&amp;ns, &amp;n, w);
 */
void
sgen(Node *n, Node *ns, int32 w)
{
	Node nodl, nodr, oldl, oldr, cx, oldcx;
	int32 c, q, odst, osrc;

	if(debug[&#39;g&#39;]) {
		print(&#34;\nsgen w=%d\n&#34;, w);
		dump(&#34;r&#34;, n);
		dump(&#34;res&#34;, ns);
	}
	if(w == 0)
		return;
	if(n-&gt;ullman &gt;= UINF &amp;&amp; ns-&gt;ullman &gt;= UINF) {
		fatal(&#34;sgen UINF&#34;);
	}

	if(w &lt; 0)
		fatal(&#34;sgen copy %d&#34;, w);

	// offset on the stack
	osrc = stkof(n);
	odst = stkof(ns);


	if(n-&gt;ullman &gt;= ns-&gt;ullman) {
		savex(D_SI, &amp;nodr, &amp;oldr, N, types[tptr]);
		agen(n, &amp;nodr);

		regalloc(&amp;nodr, types[tptr], &amp;nodr);	// mark nodr as live
		savex(D_DI, &amp;nodl, &amp;oldl, N, types[tptr]);
		agen(ns, &amp;nodl);
		regfree(&amp;nodr);
	} else {
		savex(D_DI, &amp;nodl, &amp;oldl, N, types[tptr]);
		agen(ns, &amp;nodl);

		regalloc(&amp;nodl, types[tptr], &amp;nodl);	// mark nodl as live
		savex(D_SI, &amp;nodr, &amp;oldr, N, types[tptr]);
		agen(n, &amp;nodr);
		regfree(&amp;nodl);
	}

	c = w % 8;	// bytes
	q = w / 8;	// quads

	savex(D_CX, &amp;cx, &amp;oldcx, N, types[TINT64]);

	// if we are copying forward on the stack and
	// the src and dst overlap, then reverse direction
	if(osrc &lt; odst &amp;&amp; odst &lt; osrc+w) {
		// reverse direction
		gins(ASTD, N, N);		// set direction flag
		if(c &gt; 0) {
			gconreg(AADDQ, w-1, D_SI);
			gconreg(AADDQ, w-1, D_DI);

			gconreg(AMOVQ, c, D_CX);
			gins(AREP, N, N);	// repeat
			gins(AMOVSB, N, N);	// MOVB *(SI)-,*(DI)-
		}

		if(q &gt; 0) {
			if(c &gt; 0) {
				gconreg(AADDQ, -7, D_SI);
				gconreg(AADDQ, -7, D_DI);
			} else {
				gconreg(AADDQ, w-8, D_SI);
				gconreg(AADDQ, w-8, D_DI);
			}
			gconreg(AMOVQ, q, D_CX);
			gins(AREP, N, N);	// repeat
			gins(AMOVSQ, N, N);	// MOVQ *(SI)-,*(DI)-
		}
		// we leave with the flag clear
		gins(ACLD, N, N);
	} else {
		// normal direction
		if(q &gt;= 4) {
			gconreg(AMOVQ, q, D_CX);
			gins(AREP, N, N);	// repeat
			gins(AMOVSQ, N, N);	// MOVQ *(SI)+,*(DI)+
		} else
		while(q &gt; 0) {
			gins(AMOVSQ, N, N);	// MOVQ *(SI)+,*(DI)+
			q--;
		}

		if(c &gt;= 4) {
			gins(AMOVSL, N, N);	// MOVL *(SI)+,*(DI)+
			c -= 4;
		}
		while(c &gt; 0) {
			gins(AMOVSB, N, N);	// MOVB *(SI)+,*(DI)+
			c--;
		}
	}


	restx(&amp;nodl, &amp;oldl);
	restx(&amp;nodr, &amp;oldr);
	restx(&amp;cx, &amp;oldcx);
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5c/cgen.c</title>

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
	<li>Thu Nov 12 15:48:37 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5c/cgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5c/cgen.c
// http://code.google.com/p/inferno-os/source/browse/utils/5c/cgen.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the &#34;Software&#34;), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.


#include &#34;gc.h&#34;

void
_cgen(Node *n, Node *nn, int inrel)
{
	Node *l, *r;
	Prog *p1;
	Node nod, nod1, nod2, nod3, nod4;
	int o, t;
	int32 v, curs;

	if(debug[&#39;g&#39;]) {
		prtree(nn, &#34;cgen lhs&#34;);
		prtree(n, &#34;cgen&#34;);
	}
	if(n == Z || n-&gt;type == T)
		return;
	if(typesuv[n-&gt;type-&gt;etype]) {
		sugen(n, nn, n-&gt;type-&gt;width);
		return;
	}
	l = n-&gt;left;
	r = n-&gt;right;
	o = n-&gt;op;
	if(n-&gt;addable &gt;= INDEXED) {
		if(nn == Z) {
			switch(o) {
			default:
				nullwarn(Z, Z);
				break;
			case OINDEX:
				nullwarn(l, r);
				break;
			}
			return;
		}
		gmove(n, nn);
		return;
	}
	curs = cursafe;

	if(n-&gt;complex &gt;= FNX)
	if(l-&gt;complex &gt;= FNX)
	if(r != Z &amp;&amp; r-&gt;complex &gt;= FNX)
	switch(o) {
	default:
		regret(&amp;nod, r);
		cgen(r, &amp;nod);

		regsalloc(&amp;nod1, r);
		gopcode(OAS, &amp;nod, Z, &amp;nod1);

		regfree(&amp;nod);
		nod = *n;
		nod.right = &amp;nod1;
		cgen(&amp;nod, nn);
		return;

	case OFUNC:
	case OCOMMA:
	case OANDAND:
	case OOROR:
	case OCOND:
	case ODOT:
		break;
	}

	switch(o) {
	default:
		diag(n, &#34;unknown op in cgen: %O&#34;, o);
		break;

	case OAS:
		if(l-&gt;op == OBIT)
			goto bitas;
		if(l-&gt;addable &gt;= INDEXED &amp;&amp; l-&gt;complex &lt; FNX) {
			if(nn != Z || r-&gt;addable &lt; INDEXED) {
				if(r-&gt;complex &gt;= FNX &amp;&amp; nn == Z)
					regret(&amp;nod, r);
				else
					regalloc(&amp;nod, r, nn);
				cgen(r, &amp;nod);
				gmove(&amp;nod, l);
				if(nn != Z)
					gmove(&amp;nod, nn);
				regfree(&amp;nod);
			} else
				gmove(r, l);
			break;
		}
		if(l-&gt;complex &gt;= r-&gt;complex) {
			reglcgen(&amp;nod1, l, Z);
			if(r-&gt;addable &gt;= INDEXED) {
				gmove(r, &amp;nod1);
				if(nn != Z)
					gmove(r, nn);
				regfree(&amp;nod1);
				break;
			}
			regalloc(&amp;nod, r, nn);
			cgen(r, &amp;nod);
		} else {
			regalloc(&amp;nod, r, nn);
			cgen(r, &amp;nod);
			reglcgen(&amp;nod1, l, Z);
		}
		gmove(&amp;nod, &amp;nod1);
		regfree(&amp;nod);
		regfree(&amp;nod1);
		break;

	bitas:
		n = l-&gt;left;
		regalloc(&amp;nod, r, nn);
		if(l-&gt;complex &gt;= r-&gt;complex) {
			reglcgen(&amp;nod1, n, Z);
			cgen(r, &amp;nod);
		} else {
			cgen(r, &amp;nod);
			reglcgen(&amp;nod1, n, Z);
		}
		regalloc(&amp;nod2, n, Z);
		gopcode(OAS, &amp;nod1, Z, &amp;nod2);
		bitstore(l, &amp;nod, &amp;nod1, &amp;nod2, nn);
		break;

	case OBIT:
		if(nn == Z) {
			nullwarn(l, Z);
			break;
		}
		bitload(n, &amp;nod, Z, Z, nn);
		gopcode(OAS, &amp;nod, Z, nn);
		regfree(&amp;nod);
		break;

	case ODIV:
	case OMOD:
		if(nn != Z)
		if((t = vlog(r)) &gt;= 0) {
			/* signed div/mod by constant power of 2 */
			cgen(l, nn);
			gopcode(OGE, nodconst(0), nn, Z);
			p1 = p;
			if(o == ODIV) {
				gopcode(OADD, nodconst((1&lt;&lt;t)-1), Z, nn);
				patch(p1, pc);
				gopcode(OASHR, nodconst(t), Z, nn);
			} else {
				gopcode(OSUB, nn, nodconst(0), nn);
				gopcode(OAND, nodconst((1&lt;&lt;t)-1), Z, nn);
				gopcode(OSUB, nn, nodconst(0), nn);
				gbranch(OGOTO);
				patch(p1, pc);
				p1 = p;
				gopcode(OAND, nodconst((1&lt;&lt;t)-1), Z, nn);
				patch(p1, pc);
			}
			break;
		}
		goto muldiv;

	case OSUB:
		if(nn != Z)
		if(l-&gt;op == OCONST)
		if(!typefd[n-&gt;type-&gt;etype]) {
			cgen(r, nn);
			gopcode(o, Z, l, nn);
			break;
		}
	case OADD:
	case OAND:
	case OOR:
	case OXOR:
	case OLSHR:
	case OASHL:
	case OASHR:
		/*
		 * immediate operands
		 */
		if(nn != Z)
		if(r-&gt;op == OCONST)
		if(!typefd[n-&gt;type-&gt;etype]) {
			cgen(l, nn);
			if(r-&gt;vconst == 0)
			if(o != OAND)
				break;
			if(nn != Z)
				gopcode(o, r, Z, nn);
			break;
		}

	case OLMUL:
	case OLDIV:
	case OLMOD:
	case OMUL:
	muldiv:
		if(nn == Z) {
			nullwarn(l, r);
			break;
		}
		if(o == OMUL || o == OLMUL) {
			if(mulcon(n, nn))
				break;
		}
		if(l-&gt;complex &gt;= r-&gt;complex) {
			regalloc(&amp;nod, l, nn);
			cgen(l, &amp;nod);
			regalloc(&amp;nod1, r, Z);
			cgen(r, &amp;nod1);
			gopcode(o, &amp;nod1, Z, &amp;nod);
		} else {
			regalloc(&amp;nod, r, nn);
			cgen(r, &amp;nod);
			regalloc(&amp;nod1, l, Z);
			cgen(l, &amp;nod1);
			gopcode(o, &amp;nod, &amp;nod1, &amp;nod);
		}
		gopcode(OAS, &amp;nod, Z, nn);
		regfree(&amp;nod);
		regfree(&amp;nod1);
		break;

	case OASLSHR:
	case OASASHL:
	case OASASHR:
	case OASAND:
	case OASADD:
	case OASSUB:
	case OASXOR:
	case OASOR:
		if(l-&gt;op == OBIT)
			goto asbitop;
		if(r-&gt;op == OCONST)
		if(!typefd[r-&gt;type-&gt;etype])
		if(!typefd[n-&gt;type-&gt;etype]) {
			if(l-&gt;addable &lt; INDEXED)
				reglcgen(&amp;nod2, l, Z);
			else
				nod2 = *l;
			regalloc(&amp;nod, r, nn);
			gopcode(OAS, &amp;nod2, Z, &amp;nod);
			gopcode(o, r, Z, &amp;nod);
			gopcode(OAS, &amp;nod, Z, &amp;nod2);

			regfree(&amp;nod);
			if(l-&gt;addable &lt; INDEXED)
				regfree(&amp;nod2);
			break;
		}

	case OASLMUL:
	case OASLDIV:
	case OASLMOD:
	case OASMUL:
	case OASDIV:
	case OASMOD:
		if(l-&gt;op == OBIT)
			goto asbitop;
		if(l-&gt;complex &gt;= r-&gt;complex) {
			if(l-&gt;addable &lt; INDEXED)
				reglcgen(&amp;nod2, l, Z);
			else
				nod2 = *l;
			regalloc(&amp;nod1, r, Z);
			cgen(r, &amp;nod1);
		} else {
			regalloc(&amp;nod1, r, Z);
			cgen(r, &amp;nod1);
			if(l-&gt;addable &lt; INDEXED)
				reglcgen(&amp;nod2, l, Z);
			else
				nod2 = *l;
		}

		regalloc(&amp;nod, n, nn);
		gmove(&amp;nod2, &amp;nod);
		gopcode(o, &amp;nod1, Z, &amp;nod);
		gmove(&amp;nod, &amp;nod2);
		if(nn != Z)
			gopcode(OAS, &amp;nod, Z, nn);
		regfree(&amp;nod);
		regfree(&amp;nod1);
		if(l-&gt;addable &lt; INDEXED)
			regfree(&amp;nod2);
		break;

	asbitop:
		regalloc(&amp;nod4, n, nn);
		if(l-&gt;complex &gt;= r-&gt;complex) {
			bitload(l, &amp;nod, &amp;nod1, &amp;nod2, &amp;nod4);
			regalloc(&amp;nod3, r, Z);
			cgen(r, &amp;nod3);
		} else {
			regalloc(&amp;nod3, r, Z);
			cgen(r, &amp;nod3);
			bitload(l, &amp;nod, &amp;nod1, &amp;nod2, &amp;nod4);
		}
		gmove(&amp;nod, &amp;nod4);
		gopcode(o, &amp;nod3, Z, &amp;nod4);
		regfree(&amp;nod3);
		gmove(&amp;nod4, &amp;nod);
		regfree(&amp;nod4);
		bitstore(l, &amp;nod, &amp;nod1, &amp;nod2, nn);
		break;

	case OADDR:
		if(nn == Z) {
			nullwarn(l, Z);
			break;
		}
		lcgen(l, nn);
		break;

	case OFUNC:
		if(l-&gt;complex &gt;= FNX) {
			if(l-&gt;op != OIND)
				diag(n, &#34;bad function call&#34;);

			regret(&amp;nod, l-&gt;left);
			cgen(l-&gt;left, &amp;nod);
			regsalloc(&amp;nod1, l-&gt;left);
			gopcode(OAS, &amp;nod, Z, &amp;nod1);
			regfree(&amp;nod);

			nod = *n;
			nod.left = &amp;nod2;
			nod2 = *l;
			nod2.left = &amp;nod1;
			nod2.complex = 1;
			cgen(&amp;nod, nn);

			return;
		}
		if(REGARG &gt;= 0)
			o = reg[REGARG];
		gargs(r, &amp;nod, &amp;nod1);
		if(l-&gt;addable &lt; INDEXED) {
			reglcgen(&amp;nod, l, Z);
			gopcode(OFUNC, Z, Z, &amp;nod);
			regfree(&amp;nod);
		} else
			gopcode(OFUNC, Z, Z, l);
		if(REGARG &gt;= 0)
			if(o != reg[REGARG])
				reg[REGARG]--;
		if(nn != Z) {
			regret(&amp;nod, n);
			gopcode(OAS, &amp;nod, Z, nn);
			regfree(&amp;nod);
		}
		break;

	case OIND:
		if(nn == Z) {
			nullwarn(l, Z);
			break;
		}
		regialloc(&amp;nod, n, nn);
		r = l;
		while(r-&gt;op == OADD)
			r = r-&gt;right;
		if(sconst(r) &amp;&amp; (v = r-&gt;vconst+nod.xoffset) &gt; -4096 &amp;&amp; v &lt; 4096) {
			v = r-&gt;vconst;
			r-&gt;vconst = 0;
			cgen(l, &amp;nod);
			nod.xoffset += v;
			r-&gt;vconst = v;
		} else
			cgen(l, &amp;nod);
		regind(&amp;nod, n);
		gopcode(OAS, &amp;nod, Z, nn);
		regfree(&amp;nod);
		break;

	case OEQ:
	case ONE:
	case OLE:
	case OLT:
	case OGE:
	case OGT:
	case OLO:
	case OLS:
	case OHI:
	case OHS:
		if(nn == Z) {
			nullwarn(l, r);
			break;
		}
		boolgen(n, 1, nn);
		break;

	case OANDAND:
	case OOROR:
		boolgen(n, 1, nn);
		if(nn == Z)
			patch(p, pc);
		break;

	case ONOT:
		if(nn == Z) {
			nullwarn(l, Z);
			break;
		}
		boolgen(n, 1, nn);
		break;

	case OCOMMA:
		cgen(l, Z);
		cgen(r, nn);
		break;

	case OCAST:
		if(nn == Z) {
			nullwarn(l, Z);
			break;
		}
		/*
		 * convert from types l-&gt;n-&gt;nn
		 */
		if(nocast(l-&gt;type, n-&gt;type)) {
			if(nocast(n-&gt;type, nn-&gt;type)) {
				cgen(l, nn);
				break;
			}
		}
		regalloc(&amp;nod, l, nn);
		cgen(l, &amp;nod);
		regalloc(&amp;nod1, n, &amp;nod);
		if(inrel)
			gmover(&amp;nod, &amp;nod1);
		else
			gopcode(OAS, &amp;nod, Z, &amp;nod1);
		gopcode(OAS, &amp;nod1, Z, nn);
		regfree(&amp;nod1);
		regfree(&amp;nod);
		break;

	case ODOT:
		sugen(l, nodrat, l-&gt;type-&gt;width);
		if(nn != Z) {
			warn(n, &#34;non-interruptable temporary&#34;);
			nod = *nodrat;
			if(!r || r-&gt;op != OCONST) {
				diag(n, &#34;DOT and no offset&#34;);
				break;
			}
			nod.xoffset += (int32)r-&gt;vconst;
			nod.type = n-&gt;type;
			cgen(&amp;nod, nn);
		}
		break;

	case OCOND:
		bcgen(l, 1);
		p1 = p;
		cgen(r-&gt;left, nn);
		gbranch(OGOTO);
		patch(p1, pc);
		p1 = p;
		cgen(r-&gt;right, nn);
		patch(p1, pc);
		break;

	case OPOSTINC:
	case OPOSTDEC:
		v = 1;
		if(l-&gt;type-&gt;etype == TIND)
			v = l-&gt;type-&gt;link-&gt;width;
		if(o == OPOSTDEC)
			v = -v;
		if(l-&gt;op == OBIT)
			goto bitinc;
		if(nn == Z)
			goto pre;

		if(l-&gt;addable &lt; INDEXED)
			reglcgen(&amp;nod2, l, Z);
		else
			nod2 = *l;

		regalloc(&amp;nod, l, nn);
		gopcode(OAS, &amp;nod2, Z, &amp;nod);
		regalloc(&amp;nod1, l, Z);
		if(typefd[l-&gt;type-&gt;etype]) {
			regalloc(&amp;nod3, l, Z);
			if(v &lt; 0) {
				gopcode(OAS, nodfconst(-v), Z, &amp;nod3);
				gopcode(OSUB, &amp;nod3, &amp;nod, &amp;nod1);
			} else {
				gopcode(OAS, nodfconst(v), Z, &amp;nod3);
				gopcode(OADD, &amp;nod3, &amp;nod, &amp;nod1);
			}
			regfree(&amp;nod3);
		} else
			gopcode(OADD, nodconst(v), &amp;nod, &amp;nod1);
		gopcode(OAS, &amp;nod1, Z, &amp;nod2);

		regfree(&amp;nod);
		regfree(&amp;nod1);
		if(l-&gt;addable &lt; INDEXED)
			regfree(&amp;nod2);
		break;

	case OPREINC:
	case OPREDEC:
		v = 1;
		if(l-&gt;type-&gt;etype == TIND)
			v = l-&gt;type-&gt;link-&gt;width;
		if(o == OPREDEC)
			v = -v;
		if(l-&gt;op == OBIT)
			goto bitinc;

	pre:
		if(l-&gt;addable &lt; INDEXED)
			reglcgen(&amp;nod2, l, Z);
		else
			nod2 = *l;

		regalloc(&amp;nod, l, nn);
		gopcode(OAS, &amp;nod2, Z, &amp;nod);
		if(typefd[l-&gt;type-&gt;etype]) {
			regalloc(&amp;nod3, l, Z);
			if(v &lt; 0) {
				gopcode(OAS, nodfconst(-v), Z, &amp;nod3);
				gopcode(OSUB, &amp;nod3, Z, &amp;nod);
			} else {
				gopcode(OAS, nodfconst(v), Z, &amp;nod3);
				gopcode(OADD, &amp;nod3, Z, &amp;nod);
			}
			regfree(&amp;nod3);
		} else
			gopcode(OADD, nodconst(v), Z, &amp;nod);
		gopcode(OAS, &amp;nod, Z, &amp;nod2);

		regfree(&amp;nod);
		if(l-&gt;addable &lt; INDEXED)
			regfree(&amp;nod2);
		break;

	bitinc:
		if(nn != Z &amp;&amp; (o == OPOSTINC || o == OPOSTDEC)) {
			bitload(l, &amp;nod, &amp;nod1, &amp;nod2, Z);
			gopcode(OAS, &amp;nod, Z, nn);
			gopcode(OADD, nodconst(v), Z, &amp;nod);
			bitstore(l, &amp;nod, &amp;nod1, &amp;nod2, Z);
			break;
		}
		bitload(l, &amp;nod, &amp;nod1, &amp;nod2, nn);
		gopcode(OADD, nodconst(v), Z, &amp;nod);
		bitstore(l, &amp;nod, &amp;nod1, &amp;nod2, nn);
		break;
	}
	cursafe = curs;
	return;
}

void
cgen(Node *n, Node *nn)
{
	_cgen(n, nn, 0);
}

void
cgenrel(Node *n, Node *nn)
{
	_cgen(n, nn, 1);
}

void
reglcgen(Node *t, Node *n, Node *nn)
{
	Node *r;
	int32 v;

	regialloc(t, n, nn);
	if(n-&gt;op == OIND) {
		r = n-&gt;left;
		while(r-&gt;op == OADD)
			r = r-&gt;right;
		if(sconst(r) &amp;&amp; (v = r-&gt;vconst+t-&gt;xoffset) &gt; -4096 &amp;&amp; v &lt; 4096) {
			v = r-&gt;vconst;
			r-&gt;vconst = 0;
			lcgen(n, t);
			t-&gt;xoffset += v;
			r-&gt;vconst = v;
			regind(t, n);
			return;
		}
	} else if(n-&gt;op == OINDREG) {
		if((v = n-&gt;xoffset) &gt; -4096 &amp;&amp; v &lt; 4096) {
			n-&gt;op = OREGISTER;
			cgen(n, t);
			t-&gt;xoffset += v;
			n-&gt;op = OINDREG;
			regind(t, n);
			return;
		}
	}
	lcgen(n, t);
	regind(t, n);
}

void
reglpcgen(Node *n, Node *nn, int f)
{
	Type *t;

	t = nn-&gt;type;
	nn-&gt;type = types[TLONG];
	if(f)
		reglcgen(n, nn, Z);
	else {
		regialloc(n, nn, Z);
		lcgen(nn, n);
		regind(n, nn);
	}
	nn-&gt;type = t;
}

void
lcgen(Node *n, Node *nn)
{
	Prog *p1;
	Node nod;

	if(debug[&#39;g&#39;]) {
		prtree(nn, &#34;lcgen lhs&#34;);
		prtree(n, &#34;lcgen&#34;);
	}
	if(n == Z || n-&gt;type == T)
		return;
	if(nn == Z) {
		nn = &amp;nod;
		regalloc(&amp;nod, n, Z);
	}
	switch(n-&gt;op) {
	default:
		if(n-&gt;addable &lt; INDEXED) {
			diag(n, &#34;unknown op in lcgen: %O&#34;, n-&gt;op);
			break;
		}
		nod = *n;
		nod.op = OADDR;
		nod.left = n;
		nod.right = Z;
		nod.type = types[TIND];
		gopcode(OAS, &amp;nod, Z, nn);
		break;

	case OCOMMA:
		cgen(n-&gt;left, n-&gt;left);
		lcgen(n-&gt;right, nn);
		break;

	case OIND:
		cgen(n-&gt;left, nn);
		break;

	case OCOND:
		bcgen(n-&gt;left, 1);
		p1 = p;
		lcgen(n-&gt;right-&gt;left, nn);
		gbranch(OGOTO);
		patch(p1, pc);
		p1 = p;
		lcgen(n-&gt;right-&gt;right, nn);
		patch(p1, pc);
		break;
	}
}

void
bcgen(Node *n, int true)
{

	if(n-&gt;type == T)
		gbranch(OGOTO);
	else
		boolgen(n, true, Z);
}

void
boolgen(Node *n, int true, Node *nn)
{
	int o;
	Prog *p1, *p2;
	Node *l, *r, nod, nod1;
	int32 curs;

	if(debug[&#39;g&#39;]) {
		prtree(nn, &#34;boolgen lhs&#34;);
		prtree(n, &#34;boolgen&#34;);
	}
	curs = cursafe;
	l = n-&gt;left;
	r = n-&gt;right;
	switch(n-&gt;op) {

	default:
		regalloc(&amp;nod, n, nn);
		cgen(n, &amp;nod);
		o = ONE;
		if(true)
			o = comrel[relindex(o)];
		if(typefd[n-&gt;type-&gt;etype]) {
			gopcode(o, nodfconst(0), &amp;nod, Z);
		} else
			gopcode(o, nodconst(0), &amp;nod, Z);
		regfree(&amp;nod);
		goto com;

	case OCONST:
		o = vconst(n);
		if(!true)
			o = !o;
		gbranch(OGOTO);
		if(o) {
			p1 = p;
			gbranch(OGOTO);
			patch(p1, pc);
		}
		goto com;

	case OCOMMA:
		cgen(l, Z);
		boolgen(r, true, nn);
		break;

	case ONOT:
		boolgen(l, !true, nn);
		break;

	case OCOND:
		bcgen(l, 1);
		p1 = p;
		bcgen(r-&gt;left, true);
		p2 = p;
		gbranch(OGOTO);
		patch(p1, pc);
		p1 = p;
		bcgen(r-&gt;right, !true);
		patch(p2, pc);
		p2 = p;
		gbranch(OGOTO);
		patch(p1, pc);
		patch(p2, pc);
		goto com;

	case OANDAND:
		if(!true)
			goto caseor;

	caseand:
		bcgen(l, true);
		p1 = p;
		bcgen(r, !true);
		p2 = p;
		patch(p1, pc);
		gbranch(OGOTO);
		patch(p2, pc);
		goto com;

	case OOROR:
		if(!true)
			goto caseand;

	caseor:
		bcgen(l, !true);
		p1 = p;
		bcgen(r, !true);
		p2 = p;
		gbranch(OGOTO);
		patch(p1, pc);
		patch(p2, pc);
		goto com;

	case OEQ:
	case ONE:
	case OLE:
	case OLT:
	case OGE:
	case OGT:
	case OHI:
	case OHS:
	case OLO:
	case OLS:
		o = n-&gt;op;
		if(true)
			o = comrel[relindex(o)];
		if(l-&gt;complex &gt;= FNX &amp;&amp; r-&gt;complex &gt;= FNX) {
			regret(&amp;nod, r);
			cgenrel(r, &amp;nod);
			regsalloc(&amp;nod1, r);
			gopcode(OAS, &amp;nod, Z, &amp;nod1);
			regfree(&amp;nod);
			nod = *n;
			nod.right = &amp;nod1;
			boolgen(&amp;nod, true, nn);
			break;
		}
		if(sconst(l)) {
			regalloc(&amp;nod, r, nn);
			cgenrel(r, &amp;nod);
			o = invrel[relindex(o)];
			gopcode(o, l, &amp;nod, Z);
			regfree(&amp;nod);
			goto com;
		}
		if(sconst(r)) {
			regalloc(&amp;nod, l, nn);
			cgenrel(l, &amp;nod);
			gopcode(o, r, &amp;nod, Z);
			regfree(&amp;nod);
			goto com;
		}
		if(l-&gt;complex &gt;= r-&gt;complex) {
			regalloc(&amp;nod1, l, nn);
			cgenrel(l, &amp;nod1);
			regalloc(&amp;nod, r, Z);
			cgenrel(r, &amp;nod);
		} else {
			regalloc(&amp;nod, r, nn);
			cgenrel(r, &amp;nod);
			regalloc(&amp;nod1, l, Z);
			cgenrel(l, &amp;nod1);
		}
		gopcode(o, &amp;nod, &amp;nod1, Z);
		regfree(&amp;nod);
		regfree(&amp;nod1);

	com:
		if(nn != Z) {
			p1 = p;
			gopcode(OAS, nodconst(1), Z, nn);
			gbranch(OGOTO);
			p2 = p;
			patch(p1, pc);
			gopcode(OAS, nodconst(0), Z, nn);
			patch(p2, pc);
		}
		break;
	}
	cursafe = curs;
}

void
sugen(Node *n, Node *nn, int32 w)
{
	Prog *p1;
	Node nod0, nod1, nod2, nod3, nod4, *l, *r;
	Type *t;
	int32 pc1;
	int i, m, c;

	if(n == Z || n-&gt;type == T)
		return;
	if(debug[&#39;g&#39;]) {
		prtree(nn, &#34;sugen lhs&#34;);
		prtree(n, &#34;sugen&#34;);
	}
	if(nn == nodrat)
		if(w &gt; nrathole)
			nrathole = w;
	switch(n-&gt;op) {
	case OIND:
		if(nn == Z) {
			nullwarn(n-&gt;left, Z);
			break;
		}

	default:
		goto copy;

	case OCONST:
		if(n-&gt;type &amp;&amp; typev[n-&gt;type-&gt;etype]) {
			if(nn == Z) {
				nullwarn(n-&gt;left, Z);
				break;
			}

			t = nn-&gt;type;
			nn-&gt;type = types[TLONG];
			reglcgen(&amp;nod1, nn, Z);
			nn-&gt;type = t;

			if(isbigendian)
				gopcode(OAS, nod32const(n-&gt;vconst&gt;&gt;32), Z, &amp;nod1);
			else
				gopcode(OAS, nod32const(n-&gt;vconst), Z, &amp;nod1);
			nod1.xoffset += SZ_LONG;
			if(isbigendian)
				gopcode(OAS, nod32const(n-&gt;vconst), Z, &amp;nod1);
			else
				gopcode(OAS, nod32const(n-&gt;vconst&gt;&gt;32), Z, &amp;nod1);

			regfree(&amp;nod1);
			break;
		}
		goto copy;

	case ODOT:
		l = n-&gt;left;
		sugen(l, nodrat, l-&gt;type-&gt;width);
		if(nn != Z) {
			warn(n, &#34;non-interruptable temporary&#34;);
			nod1 = *nodrat;
			r = n-&gt;right;
			if(!r || r-&gt;op != OCONST) {
				diag(n, &#34;DOT and no offset&#34;);
				break;
			}
			nod1.xoffset += (int32)r-&gt;vconst;
			nod1.type = n-&gt;type;
			sugen(&amp;nod1, nn, w);
		}
		break;

	case OSTRUCT:
		/*
		 * rewrite so lhs has no fn call
		 */
		if(nn != Z &amp;&amp; nn-&gt;complex &gt;= FNX) {
			nod1 = *n;
			nod1.type = typ(TIND, n-&gt;type);
			regret(&amp;nod2, &amp;nod1);
			lcgen(nn, &amp;nod2);
			regsalloc(&amp;nod0, &amp;nod1);
			gopcode(OAS, &amp;nod2, Z, &amp;nod0);
			regfree(&amp;nod2);

			nod1 = *n;
			nod1.op = OIND;
			nod1.left = &amp;nod0;
			nod1.right = Z;
			nod1.complex = 1;

			sugen(n, &amp;nod1, w);
			return;
		}

		r = n-&gt;left;
		for(t = n-&gt;type-&gt;link; t != T; t = t-&gt;down) {
			l = r;
			if(r-&gt;op == OLIST) {
				l = r-&gt;left;
				r = r-&gt;right;
			}
			if(nn == Z) {
				cgen(l, nn);
				continue;
			}
			/*
			 * hand craft *(&amp;nn + o) = l
			 */
			nod0 = znode;
			nod0.op = OAS;
			nod0.type = t;
			nod0.left = &amp;nod1;
			nod0.right = l;

			nod1 = znode;
			nod1.op = OIND;
			nod1.type = t;
			nod1.left = &amp;nod2;

			nod2 = znode;
			nod2.op = OADD;
			nod2.type = typ(TIND, t);
			nod2.left = &amp;nod3;
			nod2.right = &amp;nod4;

			nod3 = znode;
			nod3.op = OADDR;
			nod3.type = nod2.type;
			nod3.left = nn;

			nod4 = znode;
			nod4.op = OCONST;
			nod4.type = nod2.type;
			nod4.vconst = t-&gt;offset;

			ccom(&amp;nod0);
			acom(&amp;nod0);
			xcom(&amp;nod0);
			nod0.addable = 0;

			cgen(&amp;nod0, Z);
		}
		break;

	case OAS:
		if(nn == Z) {
			if(n-&gt;addable &lt; INDEXED)
				sugen(n-&gt;right, n-&gt;left, w);
			break;
		}
		sugen(n-&gt;right, nodrat, w);
		warn(n, &#34;non-interruptable temporary&#34;);
		sugen(nodrat, n-&gt;left, w);
		sugen(nodrat, nn, w);
		break;

	case OFUNC:
		if(nn == Z) {
			sugen(n, nodrat, w);
			break;
		}
		if(nn-&gt;op != OIND) {
			nn = new1(OADDR, nn, Z);
			nn-&gt;type = types[TIND];
			nn-&gt;addable = 0;
		} else
			nn = nn-&gt;left;
		n = new(OFUNC, n-&gt;left, new(OLIST, nn, n-&gt;right));
		n-&gt;type = types[TVOID];
		n-&gt;left-&gt;type = types[TVOID];
		cgen(n, Z);
		break;

	case OCOND:
		bcgen(n-&gt;left, 1);
		p1 = p;
		sugen(n-&gt;right-&gt;left, nn, w);
		gbranch(OGOTO);
		patch(p1, pc);
		p1 = p;
		sugen(n-&gt;right-&gt;right, nn, w);
		patch(p1, pc);
		break;

	case OCOMMA:
		cgen(n-&gt;left, Z);
		sugen(n-&gt;right, nn, w);
		break;
	}
	return;

copy:
	if(nn == Z)
		return;
	if(n-&gt;complex &gt;= FNX &amp;&amp; nn-&gt;complex &gt;= FNX) {
		t = nn-&gt;type;
		nn-&gt;type = types[TLONG];
		regialloc(&amp;nod1, nn, Z);
		lcgen(nn, &amp;nod1);
		regsalloc(&amp;nod2, nn);
		nn-&gt;type = t;

		gopcode(OAS, &amp;nod1, Z, &amp;nod2);
		regfree(&amp;nod1);

		nod2.type = typ(TIND, t);

		nod1 = nod2;
		nod1.op = OIND;
		nod1.left = &amp;nod2;
		nod1.right = Z;
		nod1.complex = 1;
		nod1.type = t;

		sugen(n, &amp;nod1, w);
		return;
	}

	w /= SZ_LONG;
	if(w &lt;= 2) {
		if(n-&gt;complex &gt; nn-&gt;complex) {
			reglpcgen(&amp;nod1, n, 1);
			reglpcgen(&amp;nod2, nn, 1);
		} else {
			reglpcgen(&amp;nod2, nn, 1);
			reglpcgen(&amp;nod1, n, 1);
		}
		regalloc(&amp;nod3, &amp;regnode, Z);
		regalloc(&amp;nod4, &amp;regnode, Z);
		nod0 = *nodconst((1&lt;&lt;nod3.reg)|(1&lt;&lt;nod4.reg));
		if(w == 2 &amp;&amp; nod1.xoffset == 0)
			gmovm(&amp;nod1, &amp;nod0, 0);
		else {
			gmove(&amp;nod1, &amp;nod3);
			if(w == 2) {
				nod1.xoffset += SZ_LONG;
				gmove(&amp;nod1, &amp;nod4);
			}
		}
		if(w == 2 &amp;&amp; nod2.xoffset == 0)
			gmovm(&amp;nod0, &amp;nod2, 0);
		else {
			gmove(&amp;nod3, &amp;nod2);
			if(w == 2) {
				nod2.xoffset += SZ_LONG;
				gmove(&amp;nod4, &amp;nod2);
			}
		}
		regfree(&amp;nod1);
		regfree(&amp;nod2);
		regfree(&amp;nod3);
		regfree(&amp;nod4);
		return;
	}

	if(n-&gt;complex &gt; nn-&gt;complex) {
		reglpcgen(&amp;nod1, n, 0);
		reglpcgen(&amp;nod2, nn, 0);
	} else {
		reglpcgen(&amp;nod2, nn, 0);
		reglpcgen(&amp;nod1, n, 0);
	}

	m = 0;
	for(c = 0; c &lt; w &amp;&amp; c &lt; 4; c++) {
		i = tmpreg();
		if (i == 0)
			break;
		reg[i]++;
		m |= 1&lt;&lt;i;
	}
	nod4 = *(nodconst(m));
	if(w &lt; 3*c) {
		for (; w&gt;c; w-=c) {
			gmovm(&amp;nod1, &amp;nod4, 1);
			gmovm(&amp;nod4, &amp;nod2, 1);
		}
		goto out;
	}

	regalloc(&amp;nod3, &amp;regnode, Z);
	gopcode(OAS, nodconst(w/c), Z, &amp;nod3);
	w %= c;

	pc1 = pc;
	gmovm(&amp;nod1, &amp;nod4, 1);
	gmovm(&amp;nod4, &amp;nod2, 1);

	gopcode(OSUB, nodconst(1), Z, &amp;nod3);
	gopcode(OEQ, nodconst(0), &amp;nod3, Z);
	p-&gt;as = ABGT;
	patch(p, pc1);
	regfree(&amp;nod3);

out:
	if (w) {
		i = 0;
		while (c&gt;w) {
			while ((m&amp;(1&lt;&lt;i)) == 0)
				i++;
			m &amp;= ~(1&lt;&lt;i);
			reg[i] = 0;
			c--;
			i++;
		}
		nod4.vconst = m;
		gmovm(&amp;nod1, &amp;nod4, 0);
		gmovm(&amp;nod4, &amp;nod2, 0);
	}
	i = 0;
	do {
		while ((m&amp;(1&lt;&lt;i)) == 0)
			i++;
		reg[i] = 0;
		c--;
		i++;
	} while (c&gt;0);
	regfree(&amp;nod1);
	regfree(&amp;nod2);
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

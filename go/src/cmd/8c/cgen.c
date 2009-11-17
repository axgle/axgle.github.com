<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8c/cgen.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8c/cgen.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/8c/cgen.c
// http://code.google.com/p/inferno-os/source/browse/utils/8c/cgen.c
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

/* ,x/^(print|prtree)\(/i/\/\/ */

void
cgen(Node *n, Node *nn)
{
	Node *l, *r, *t;
	Prog *p1;
	Node nod, nod1, nod2, nod3, nod4;
	int o, hardleft;
	int32 v, curs;
	vlong c;

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

	if(l-&gt;complex &gt;= FNX)
	if(r != Z &amp;&amp; r-&gt;complex &gt;= FNX)
	switch(o) {
	default:
		if(cond(o) &amp;&amp; typesuv[l-&gt;type-&gt;etype])
			break;

		regret(&amp;nod, r);
		cgen(r, &amp;nod);

		regsalloc(&amp;nod1, r);
		gmove(&amp;nod, &amp;nod1);

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

	hardleft = l-&gt;addable &lt; INDEXED || l-&gt;complex &gt;= FNX;
	switch(o) {
	default:
		diag(n, &#34;unknown op in cgen: %O&#34;, o);
		break;

	case ONEG:
	case OCOM:
		if(nn == Z) {
			nullwarn(l, Z);
			break;
		}
		regalloc(&amp;nod, l, nn);
		cgen(l, &amp;nod);
		gopcode(o, n-&gt;type, Z, &amp;nod);
		gmove(&amp;nod, nn);
		regfree(&amp;nod);
		break;

	case OAS:
		if(typefd[n-&gt;type-&gt;etype]) {
			cgen(r, &amp;fregnode0);
			if(nn != Z)
				gins(AFMOVD, &amp;fregnode0, &amp;fregnode0);
			if(l-&gt;addable &lt; INDEXED) {
				reglcgen(&amp;nod, l, Z);
				gmove(&amp;fregnode0, &amp;nod);
				regfree(&amp;nod);
			} else
				gmove(&amp;fregnode0, l);
			if(nn != Z)
				gmove(&amp;fregnode0, nn);
			return;
		}
		if(l-&gt;op == OBIT)
			goto bitas;
		if(!hardleft) {
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
			if(l-&gt;op == OINDEX &amp;&amp; r-&gt;op == OCONST) {
				gmove(r, l);
				break;
			}
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
		gmove(&amp;nod1, &amp;nod2);
		bitstore(l, &amp;nod, &amp;nod1, &amp;nod2, nn);
		break;

	case OBIT:
		if(nn == Z) {
			nullwarn(l, Z);
			break;
		}
		bitload(n, &amp;nod, Z, Z, nn);
		gmove(&amp;nod, nn);
		regfree(&amp;nod);
		break;

	case OLSHR:
	case OASHL:
	case OASHR:
		if(nn == Z) {
			nullwarn(l, r);
			break;
		}
		if(r-&gt;op == OCONST) {
			if(r-&gt;vconst == 0) {
				cgen(l, nn);
				break;
			}
			regalloc(&amp;nod, l, nn);
			cgen(l, &amp;nod);
			if(o == OASHL &amp;&amp; r-&gt;vconst == 1)
				gopcode(OADD, n-&gt;type, &amp;nod, &amp;nod);
			else
				gopcode(o, n-&gt;type, r, &amp;nod);
			gmove(&amp;nod, nn);
			regfree(&amp;nod);
			break;
		}

		/*
		 * get nod to be D_CX
		 */
		if(nodreg(&amp;nod, nn, D_CX)) {
			regsalloc(&amp;nod1, n);
			gmove(&amp;nod, &amp;nod1);
			cgen(n, &amp;nod);		/* probably a bug */
			gmove(&amp;nod, nn);
			gmove(&amp;nod1, &amp;nod);
			break;
		}
		reg[D_CX]++;
		if(nn-&gt;op == OREGISTER &amp;&amp; nn-&gt;reg == D_CX)
			regalloc(&amp;nod1, l, Z);
		else
			regalloc(&amp;nod1, l, nn);
		if(r-&gt;complex &gt;= l-&gt;complex) {
			cgen(r, &amp;nod);
			cgen(l, &amp;nod1);
		} else {
			cgen(l, &amp;nod1);
			cgen(r, &amp;nod);
		}
		gopcode(o, n-&gt;type, &amp;nod, &amp;nod1);
		gmove(&amp;nod1, nn);
		regfree(&amp;nod);
		regfree(&amp;nod1);
		break;

	case OADD:
	case OSUB:
	case OOR:
	case OXOR:
	case OAND:
		if(nn == Z) {
			nullwarn(l, r);
			break;
		}
		if(typefd[n-&gt;type-&gt;etype])
			goto fop;
		if(r-&gt;op == OCONST) {
			if(r-&gt;vconst == 0 &amp;&amp; o != OAND) {
				cgen(l, nn);
				break;
			}
		}
		if(n-&gt;op == OADD &amp;&amp; l-&gt;op == OASHL &amp;&amp; l-&gt;right-&gt;op == OCONST
		&amp;&amp; (r-&gt;op != OCONST || r-&gt;vconst &lt; -128 || r-&gt;vconst &gt; 127)) {
			c = l-&gt;right-&gt;vconst;
			if(c &gt; 0 &amp;&amp; c &lt;= 3) {
				if(l-&gt;left-&gt;complex &gt;= r-&gt;complex) {
					regalloc(&amp;nod, l-&gt;left, nn);
					cgen(l-&gt;left, &amp;nod);
					if(r-&gt;addable &lt; INDEXED) {
						regalloc(&amp;nod1, r, Z);
						cgen(r, &amp;nod1);
						genmuladd(&amp;nod, &amp;nod, 1 &lt;&lt; c, &amp;nod1);
						regfree(&amp;nod1);
					}
					else
						genmuladd(&amp;nod, &amp;nod, 1 &lt;&lt; c, r);
				}
				else {
					regalloc(&amp;nod, r, nn);
					cgen(r, &amp;nod);
					regalloc(&amp;nod1, l-&gt;left, Z);
					cgen(l-&gt;left, &amp;nod1);
					genmuladd(&amp;nod, &amp;nod1, 1 &lt;&lt; c, &amp;nod);
					regfree(&amp;nod1);
				}
				gmove(&amp;nod, nn);
				regfree(&amp;nod);
				break;
			}
		}
		if(r-&gt;addable &gt;= INDEXED) {
			regalloc(&amp;nod, l, nn);
			cgen(l, &amp;nod);
			gopcode(o, n-&gt;type, r, &amp;nod);
			gmove(&amp;nod, nn);
			regfree(&amp;nod);
			break;
		}
		if(l-&gt;complex &gt;= r-&gt;complex) {
			regalloc(&amp;nod, l, nn);
			cgen(l, &amp;nod);
			regalloc(&amp;nod1, r, Z);
			cgen(r, &amp;nod1);
			gopcode(o, n-&gt;type, &amp;nod1, &amp;nod);
		} else {
			regalloc(&amp;nod1, r, nn);
			cgen(r, &amp;nod1);
			regalloc(&amp;nod, l, Z);
			cgen(l, &amp;nod);
			gopcode(o, n-&gt;type, &amp;nod1, &amp;nod);
		}
		gmove(&amp;nod, nn);
		regfree(&amp;nod);
		regfree(&amp;nod1);
		break;

	case OLMOD:
	case OMOD:
	case OLMUL:
	case OLDIV:
	case OMUL:
	case ODIV:
		if(nn == Z) {
			nullwarn(l, r);
			break;
		}
		if(typefd[n-&gt;type-&gt;etype])
			goto fop;
		if(r-&gt;op == OCONST) {
			SET(v);
			switch(o) {
			case ODIV:
			case OMOD:
				c = r-&gt;vconst;
				if(c &lt; 0)
					c = -c;
				v = xlog2(c);
				if(v &lt; 0)
					break;
				/* fall thru */
			case OMUL:
			case OLMUL:
				regalloc(&amp;nod, l, nn);
				cgen(l, &amp;nod);
				switch(o) {
				case OMUL:
				case OLMUL:
					mulgen(n-&gt;type, r, &amp;nod);
					break;
				case ODIV:
					sdiv2(r-&gt;vconst, v, l, &amp;nod);
					break;
				case OMOD:
					smod2(r-&gt;vconst, v, l, &amp;nod);
					break;
				}
				gmove(&amp;nod, nn);
				regfree(&amp;nod);
				goto done;
			case OLDIV:
				c = r-&gt;vconst;
				if((c &amp; 0x80000000) == 0)
					break;
				regalloc(&amp;nod1, l, Z);
				cgen(l, &amp;nod1);
				regalloc(&amp;nod, l, nn);
				zeroregm(&amp;nod);
				gins(ACMPL, &amp;nod1, nodconst(c));
				gins(ASBBL, nodconst(-1), &amp;nod);
				regfree(&amp;nod1);
				gmove(&amp;nod, nn);
				regfree(&amp;nod);
				goto done;
			}
		}

		if(o == OMUL) {
			if(l-&gt;addable &gt;= INDEXED) {
				t = l;
				l = r;
				r = t;
			}
			/* should favour AX */
			regalloc(&amp;nod, l, nn);
			cgen(l, &amp;nod);
			if(r-&gt;addable &lt; INDEXED) {
				regalloc(&amp;nod1, r, Z);
				cgen(r, &amp;nod1);
				gopcode(OMUL, n-&gt;type, &amp;nod1, &amp;nod);
				regfree(&amp;nod1);
			}else
				gopcode(OMUL, n-&gt;type, r, &amp;nod);	/* addressible */
			gmove(&amp;nod, nn);
			regfree(&amp;nod);
			break;
		}

		/*
		 * get nod to be D_AX
		 * get nod1 to be D_DX
		 */
		if(nodreg(&amp;nod, nn, D_AX)) {
			regsalloc(&amp;nod2, n);
			gmove(&amp;nod, &amp;nod2);
			v = reg[D_AX];
			reg[D_AX] = 0;

			if(isreg(l, D_AX)) {
				nod3 = *n;
				nod3.left = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
			if(isreg(r, D_AX)) {
				nod3 = *n;
				nod3.right = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
				cgen(n, nn);

			gmove(&amp;nod2, &amp;nod);
			reg[D_AX] = v;
			break;
		}
		if(nodreg(&amp;nod1, nn, D_DX)) {
			regsalloc(&amp;nod2, n);
			gmove(&amp;nod1, &amp;nod2);
			v = reg[D_DX];
			reg[D_DX] = 0;

			if(isreg(l, D_DX)) {
				nod3 = *n;
				nod3.left = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
			if(isreg(r, D_DX)) {
				nod3 = *n;
				nod3.right = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
				cgen(n, nn);

			gmove(&amp;nod2, &amp;nod1);
			reg[D_DX] = v;
			break;
		}
		reg[D_AX]++;

		if(r-&gt;op == OCONST &amp;&amp; (o == ODIV || o == OLDIV)) {
			reg[D_DX]++;
			if(l-&gt;addable &lt; INDEXED) {
				regalloc(&amp;nod2, l, Z);
				cgen(l, &amp;nod2);
				l = &amp;nod2;
			}
			if(o == ODIV)
				sdivgen(l, r, &amp;nod, &amp;nod1);
			else
				udivgen(l, r, &amp;nod, &amp;nod1);
			gmove(&amp;nod1, nn);
			if(l == &amp;nod2)
				regfree(l);
			goto freeaxdx;
		}

		if(l-&gt;complex &gt;= r-&gt;complex) {
			cgen(l, &amp;nod);
			reg[D_DX]++;
			if(o == ODIV || o == OMOD)
				gins(ACDQ, Z, Z);
			if(o == OLDIV || o == OLMOD)
				zeroregm(&amp;nod1);
			if(r-&gt;addable &lt; INDEXED || r-&gt;op == OCONST) {
				regsalloc(&amp;nod3, r);
				cgen(r, &amp;nod3);
				gopcode(o, n-&gt;type, &amp;nod3, Z);
			} else
				gopcode(o, n-&gt;type, r, Z);
		} else {
			regsalloc(&amp;nod3, r);
			cgen(r, &amp;nod3);
			cgen(l, &amp;nod);
			reg[D_DX]++;
			if(o == ODIV || o == OMOD)
				gins(ACDQ, Z, Z);
			if(o == OLDIV || o == OLMOD)
				zeroregm(&amp;nod1);
			gopcode(o, n-&gt;type, &amp;nod3, Z);
		}
		if(o == OMOD || o == OLMOD)
			gmove(&amp;nod1, nn);
		else
			gmove(&amp;nod, nn);
	freeaxdx:
		regfree(&amp;nod);
		regfree(&amp;nod1);
		break;

	case OASLSHR:
	case OASASHL:
	case OASASHR:
		if(r-&gt;op == OCONST)
			goto asand;
		if(l-&gt;op == OBIT)
			goto asbitop;
		if(typefd[n-&gt;type-&gt;etype])
			goto asfop;

		/*
		 * get nod to be D_CX
		 */
		if(nodreg(&amp;nod, nn, D_CX)) {
			regsalloc(&amp;nod1, n);
			gmove(&amp;nod, &amp;nod1);
			cgen(n, &amp;nod);
			if(nn != Z)
				gmove(&amp;nod, nn);
			gmove(&amp;nod1, &amp;nod);
			break;
		}
		reg[D_CX]++;

		if(r-&gt;complex &gt;= l-&gt;complex) {
			cgen(r, &amp;nod);
			if(hardleft)
				reglcgen(&amp;nod1, l, Z);
			else
				nod1 = *l;
		} else {
			if(hardleft)
				reglcgen(&amp;nod1, l, Z);
			else
				nod1 = *l;
			cgen(r, &amp;nod);
		}

		gopcode(o, l-&gt;type, &amp;nod, &amp;nod1);
		regfree(&amp;nod);
		if(nn != Z)
			gmove(&amp;nod1, nn);
		if(hardleft)
			regfree(&amp;nod1);
		break;

	case OASAND:
	case OASADD:
	case OASSUB:
	case OASXOR:
	case OASOR:
	asand:
		if(l-&gt;op == OBIT)
			goto asbitop;
		if(typefd[n-&gt;type-&gt;etype]||typefd[r-&gt;type-&gt;etype])
			goto asfop;
		if(l-&gt;complex &gt;= r-&gt;complex) {
			if(hardleft)
				reglcgen(&amp;nod, l, Z);
			else
				nod = *l;
			if(r-&gt;op != OCONST) {
				regalloc(&amp;nod1, r, nn);
				cgen(r, &amp;nod1);
				gopcode(o, l-&gt;type, &amp;nod1, &amp;nod);
				regfree(&amp;nod1);
			} else
				gopcode(o, l-&gt;type, r, &amp;nod);
		} else {
			regalloc(&amp;nod1, r, nn);
			cgen(r, &amp;nod1);
			if(hardleft)
				reglcgen(&amp;nod, l, Z);
			else
				nod = *l;
			gopcode(o, l-&gt;type, &amp;nod1, &amp;nod);
			regfree(&amp;nod1);
		}
		if(nn != Z)
			gmove(&amp;nod, nn);
		if(hardleft)
			regfree(&amp;nod);
		break;

	case OASLMUL:
	case OASLDIV:
	case OASLMOD:
	case OASMUL:
	case OASDIV:
	case OASMOD:
		if(l-&gt;op == OBIT)
			goto asbitop;
		if(typefd[n-&gt;type-&gt;etype]||typefd[r-&gt;type-&gt;etype])
			goto asfop;
		if(r-&gt;op == OCONST) {
			SET(v);
			switch(o) {
			case OASDIV:
			case OASMOD:
				c = r-&gt;vconst;
				if(c &lt; 0)
					c = -c;
				v = xlog2(c);
				if(v &lt; 0)
					break;
				/* fall thru */
			case OASMUL:
			case OASLMUL:
				if(hardleft)
					reglcgen(&amp;nod2, l, Z);
				else
					nod2 = *l;
				regalloc(&amp;nod, l, nn);
				cgen(&amp;nod2, &amp;nod);
				switch(o) {
				case OASMUL:
				case OASLMUL:
					mulgen(n-&gt;type, r, &amp;nod);
					break;
				case OASDIV:
					sdiv2(r-&gt;vconst, v, l, &amp;nod);
					break;
				case OASMOD:
					smod2(r-&gt;vconst, v, l, &amp;nod);
					break;
				}
			havev:
				gmove(&amp;nod, &amp;nod2);
				if(nn != Z)
					gmove(&amp;nod, nn);
				if(hardleft)
					regfree(&amp;nod2);
				regfree(&amp;nod);
				goto done;
			case OASLDIV:
				c = r-&gt;vconst;
				if((c &amp; 0x80000000) == 0)
					break;
				if(hardleft)
					reglcgen(&amp;nod2, l, Z);
				else
					nod2 = *l;
				regalloc(&amp;nod1, l, nn);
				cgen(&amp;nod2, &amp;nod1);
				regalloc(&amp;nod, l, nn);
				zeroregm(&amp;nod);
				gins(ACMPL, &amp;nod1, nodconst(c));
				gins(ASBBL, nodconst(-1), &amp;nod);
				regfree(&amp;nod1);
				goto havev;
			}
		}

		if(o == OASMUL) {
			/* should favour AX */
			regalloc(&amp;nod, l, nn);
			if(r-&gt;complex &gt;= FNX) {
				regalloc(&amp;nod1, r, Z);
				cgen(r, &amp;nod1);
				r = &amp;nod1;
			}
			if(hardleft)
				reglcgen(&amp;nod2, l, Z);
			else
				nod2 = *l;
			cgen(&amp;nod2, &amp;nod);
			if(r-&gt;addable &lt; INDEXED) {
				if(r-&gt;complex &lt; FNX) {
					regalloc(&amp;nod1, r, Z);
					cgen(r, &amp;nod1);
				}
				gopcode(OASMUL, n-&gt;type, &amp;nod1, &amp;nod);
				regfree(&amp;nod1);
			}
			else
				gopcode(OASMUL, n-&gt;type, r, &amp;nod);
			if(r == &amp;nod1)
				regfree(r);
			gmove(&amp;nod, &amp;nod2);
			if(nn != Z)
				gmove(&amp;nod, nn);
			regfree(&amp;nod);
			if(hardleft)
				regfree(&amp;nod2);
			break;
		}

		/*
		 * get nod to be D_AX
		 * get nod1 to be D_DX
		 */
		if(nodreg(&amp;nod, nn, D_AX)) {
			regsalloc(&amp;nod2, n);
			gmove(&amp;nod, &amp;nod2);
			v = reg[D_AX];
			reg[D_AX] = 0;

			if(isreg(l, D_AX)) {
				nod3 = *n;
				nod3.left = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
			if(isreg(r, D_AX)) {
				nod3 = *n;
				nod3.right = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
				cgen(n, nn);

			gmove(&amp;nod2, &amp;nod);
			reg[D_AX] = v;
			break;
		}
		if(nodreg(&amp;nod1, nn, D_DX)) {
			regsalloc(&amp;nod2, n);
			gmove(&amp;nod1, &amp;nod2);
			v = reg[D_DX];
			reg[D_DX] = 0;

			if(isreg(l, D_DX)) {
				nod3 = *n;
				nod3.left = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
			if(isreg(r, D_DX)) {
				nod3 = *n;
				nod3.right = &amp;nod2;
				cgen(&amp;nod3, nn);
			} else
				cgen(n, nn);

			gmove(&amp;nod2, &amp;nod1);
			reg[D_DX] = v;
			break;
		}
		reg[D_AX]++;
		reg[D_DX]++;

		if(l-&gt;complex &gt;= r-&gt;complex) {
			if(hardleft)
				reglcgen(&amp;nod2, l, Z);
			else
				nod2 = *l;
			cgen(&amp;nod2, &amp;nod);
			if(r-&gt;op == OCONST) {
				switch(o) {
				case OASDIV:
					sdivgen(&amp;nod2, r, &amp;nod, &amp;nod1);
					goto divdone;
				case OASLDIV:
					udivgen(&amp;nod2, r, &amp;nod, &amp;nod1);
				divdone:
					gmove(&amp;nod1, &amp;nod2);
					if(nn != Z)
						gmove(&amp;nod1, nn);
					goto freelxaxdx;
				}
			}
			if(o == OASDIV || o == OASMOD)
				gins(ACDQ, Z, Z);
			if(o == OASLDIV || o == OASLMOD)
				zeroregm(&amp;nod1);
			if(r-&gt;addable &lt; INDEXED || r-&gt;op == OCONST ||
			   !typeil[r-&gt;type-&gt;etype]) {
				regalloc(&amp;nod3, r, Z);
				cgen(r, &amp;nod3);
				gopcode(o, l-&gt;type, &amp;nod3, Z);
				regfree(&amp;nod3);
			} else
				gopcode(o, n-&gt;type, r, Z);
		} else {
			regalloc(&amp;nod3, r, Z);
			cgen(r, &amp;nod3);
			if(hardleft)
				reglcgen(&amp;nod2, l, Z);
			else
				nod2 = *l;
			cgen(&amp;nod2, &amp;nod);
			if(o == OASDIV || o == OASMOD)
				gins(ACDQ, Z, Z);
			if(o == OASLDIV || o == OASLMOD)
				zeroregm(&amp;nod1);
			gopcode(o, l-&gt;type, &amp;nod3, Z);
			regfree(&amp;nod3);
		}
		if(o == OASMOD || o == OASLMOD) {
			gmove(&amp;nod1, &amp;nod2);
			if(nn != Z)
				gmove(&amp;nod1, nn);
		} else {
			gmove(&amp;nod, &amp;nod2);
			if(nn != Z)
				gmove(&amp;nod, nn);
		}
	freelxaxdx:
		if(hardleft)
			regfree(&amp;nod2);
		regfree(&amp;nod);
		regfree(&amp;nod1);
		break;

	fop:
		if(l-&gt;complex &gt;= r-&gt;complex) {
			cgen(l, &amp;fregnode0);
			if(r-&gt;addable &lt; INDEXED) {
				cgen(r, &amp;fregnode0);
				fgopcode(o, &amp;fregnode0, &amp;fregnode1, 1, 0);
			} else
				fgopcode(o, r, &amp;fregnode0, 0, 0);
		} else {
			cgen(r, &amp;fregnode0);
			if(l-&gt;addable &lt; INDEXED) {
				cgen(l, &amp;fregnode0);
				fgopcode(o, &amp;fregnode0, &amp;fregnode1, 1, 1);
			} else
				fgopcode(o, l, &amp;fregnode0, 0, 1);
		}
		gmove(&amp;fregnode0, nn);
		break;

	asfop:
		if(l-&gt;complex &gt;= r-&gt;complex) {
			if(hardleft)
				reglcgen(&amp;nod, l, Z);
			else
				nod = *l;
			cgen(r, &amp;fregnode0);
		} else {
			cgen(r, &amp;fregnode0);
			if(hardleft)
				reglcgen(&amp;nod, l, Z);
			else
				nod = *l;
		}
		if(!typefd[l-&gt;type-&gt;etype]) {
			gmove(&amp;nod, &amp;fregnode0);
			fgopcode(o, &amp;fregnode0, &amp;fregnode1, 1, 1);
		} else
			fgopcode(o, &amp;nod, &amp;fregnode0, 0, 1);
		if(nn != Z)
			gins(AFMOVD, &amp;fregnode0, &amp;fregnode0);
		gmove(&amp;fregnode0, &amp;nod);
		if(nn != Z)
			gmove(&amp;fregnode0, nn);
		if(hardleft)
			regfree(&amp;nod);
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

		if(typefd[nod3.type-&gt;etype])
			fgopcode(o, &amp;fregnode0, &amp;fregnode1, 1, 1);
		else {
			Node onod;

			/* incredible grot ... */
			onod = nod3;
			onod.op = o;
			onod.complex = 2;
			onod.addable = 0;
			onod.type = tfield;
			onod.left = &amp;nod4;
			onod.right = &amp;nod3;
			cgen(&amp;onod, Z);
		}
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
			gmove(&amp;nod, &amp;nod1);
			regfree(&amp;nod);

			nod = *n;
			nod.left = &amp;nod2;
			nod2 = *l;
			nod2.left = &amp;nod1;
			nod2.complex = 1;
			cgen(&amp;nod, nn);

			return;
		}
		gargs(r, &amp;nod, &amp;nod1);
		if(l-&gt;addable &lt; INDEXED) {
			reglcgen(&amp;nod, l, nn);
			nod.op = OREGISTER;
			gopcode(OFUNC, n-&gt;type, Z, &amp;nod);
			regfree(&amp;nod);
		} else
			gopcode(OFUNC, n-&gt;type, Z, l);
		if(REGARG &gt;= 0 &amp;&amp; reg[REGARG])
			reg[REGARG]--;
		if(nn != Z) {
			regret(&amp;nod, n);
			gmove(&amp;nod, nn);
			regfree(&amp;nod);
		} else
		if(typefd[n-&gt;type-&gt;etype])
			gins(AFMOVDP, &amp;fregnode0, &amp;fregnode0);
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
		if(sconst(r)) {
			v = r-&gt;vconst;
			r-&gt;vconst = 0;
			cgen(l, &amp;nod);
			nod.xoffset += v;
			r-&gt;vconst = v;
		} else
			cgen(l, &amp;nod);
		regind(&amp;nod, n);
		gmove(&amp;nod, nn);
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
		if(nocast(l-&gt;type, n-&gt;type) &amp;&amp; nocast(n-&gt;type, nn-&gt;type)) {
			/* both null, gen l-&gt;nn */
			cgen(l, nn);
			break;
		}
		if(typev[l-&gt;type-&gt;etype]) {
			cgen64(n, nn);
			break;
		}
		regalloc(&amp;nod, l, nn);
		cgen(l, &amp;nod);
		regalloc(&amp;nod1, n, &amp;nod);
		gmove(&amp;nod, &amp;nod1);
		gmove(&amp;nod1, nn);
		regfree(&amp;nod1);
		regfree(&amp;nod);
		break;

	case ODOT:
		sugen(l, nodrat, l-&gt;type-&gt;width);
		if(nn == Z)
			break;
		warn(n, &#34;non-interruptable temporary&#34;);
		nod = *nodrat;
		if(!r || r-&gt;op != OCONST) {
			diag(n, &#34;DOT and no offset&#34;);
			break;
		}
		nod.xoffset += (int32)r-&gt;vconst;
		nod.type = n-&gt;type;
		cgen(&amp;nod, nn);
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

		if(hardleft)
			reglcgen(&amp;nod, l, Z);
		else
			nod = *l;

		if(typefd[n-&gt;type-&gt;etype])
			goto fltinc;
		gmove(&amp;nod, nn);
		gopcode(OADD, n-&gt;type, nodconst(v), &amp;nod);
		if(hardleft)
			regfree(&amp;nod);
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
		if(hardleft)
			reglcgen(&amp;nod, l, Z);
		else
			nod = *l;
		if(typefd[n-&gt;type-&gt;etype])
			goto fltinc;
		gopcode(OADD, n-&gt;type, nodconst(v), &amp;nod);
		if(nn != Z)
			gmove(&amp;nod, nn);
		if(hardleft)
			regfree(&amp;nod);
		break;

	fltinc:
		gmove(&amp;nod, &amp;fregnode0);
		if(nn != Z &amp;&amp; (o == OPOSTINC || o == OPOSTDEC))
			gins(AFMOVD, &amp;fregnode0, &amp;fregnode0);
		gins(AFLD1, Z, Z);
		if(v &lt; 0)
			fgopcode(OSUB, &amp;fregnode0, &amp;fregnode1, 1, 0);
		else
			fgopcode(OADD, &amp;fregnode0, &amp;fregnode1, 1, 0);
		if(nn != Z &amp;&amp; (o == OPREINC || o == OPREDEC))
			gins(AFMOVD, &amp;fregnode0, &amp;fregnode0);
		gmove(&amp;fregnode0, &amp;nod);
		if(hardleft)
			regfree(&amp;nod);
		break;

	bitinc:
		if(nn != Z &amp;&amp; (o == OPOSTINC || o == OPOSTDEC)) {
			bitload(l, &amp;nod, &amp;nod1, &amp;nod2, Z);
			gmove(&amp;nod, nn);
			gopcode(OADD, tfield, nodconst(v), &amp;nod);
			bitstore(l, &amp;nod, &amp;nod1, &amp;nod2, Z);
			break;
		}
		bitload(l, &amp;nod, &amp;nod1, &amp;nod2, nn);
		gopcode(OADD, tfield, nodconst(v), &amp;nod);
		bitstore(l, &amp;nod, &amp;nod1, &amp;nod2, nn);
		break;
	}
done:
	cursafe = curs;
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
		if(sconst(r)) {
			v = r-&gt;vconst;
			r-&gt;vconst = 0;
			lcgen(n, t);
			t-&gt;xoffset += v;
			r-&gt;vconst = v;
			regind(t, n);
			return;
		}
	}
	lcgen(n, t);
	regind(t, n);
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
		gopcode(OADDR, n-&gt;type, n, nn);
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
		if(typev[n-&gt;type-&gt;etype]) {
			testv(n, true);
			goto com;
		}
		o = ONE;
		if(true)
			o = OEQ;
		if(typefd[n-&gt;type-&gt;etype]) {
			if(n-&gt;addable &lt; INDEXED) {
				cgen(n, &amp;fregnode0);
				gins(AFLDZ, Z, Z);
				fgopcode(o, &amp;fregnode0, &amp;fregnode1, 1, 1);
			} else {
				gins(AFLDZ, Z, Z);
				fgopcode(o, n, &amp;fregnode0, 0, 1);
			}
			goto com;
		}
		/* bad, 13 is address of external that becomes constant */
		if(n-&gt;addable &gt;= INDEXED &amp;&amp; n-&gt;addable != 13) {
			gopcode(o, n-&gt;type, n, nodconst(0));
			goto com;
		}
		regalloc(&amp;nod, n, nn);
		cgen(n, &amp;nod);
		gopcode(o, n-&gt;type, &amp;nod, nodconst(0));
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
		if(typev[l-&gt;type-&gt;etype]) {
			if(!true)
				n-&gt;op = comrel[relindex(o)];
			cgen64(n, Z);
			goto com;
		}
		if(true)
			o = comrel[relindex(o)];
		if(l-&gt;complex &gt;= FNX &amp;&amp; r-&gt;complex &gt;= FNX) {
			regret(&amp;nod, r);
			cgen(r, &amp;nod);
			regsalloc(&amp;nod1, r);
			gmove(&amp;nod, &amp;nod1);
			regfree(&amp;nod);
			nod = *n;
			nod.right = &amp;nod1;
			boolgen(&amp;nod, true, nn);
			break;
		}
		if(typefd[l-&gt;type-&gt;etype]) {
			if(l-&gt;complex &gt;= r-&gt;complex) {
				cgen(l, &amp;fregnode0);
				if(r-&gt;addable &lt; INDEXED) {
					cgen(r, &amp;fregnode0);
					o = invrel[relindex(o)];
					fgopcode(o, &amp;fregnode0, &amp;fregnode1, 1, 1);
				} else
					fgopcode(o, r, &amp;fregnode0, 0, 1);
			} else {
				o = invrel[relindex(o)];
				cgen(r, &amp;fregnode0);
				if(l-&gt;addable &lt; INDEXED) {
					cgen(l, &amp;fregnode0);
					o = invrel[relindex(o)];
					fgopcode(o, &amp;fregnode0, &amp;fregnode1, 1, 1);
				} else
					fgopcode(o, l, &amp;fregnode0, 0, 1);
			}
			goto com;
		}
		if(l-&gt;op == OCONST) {
			o = invrel[relindex(o)];
			/* bad, 13 is address of external that becomes constant */
			if(r-&gt;addable &lt; INDEXED || r-&gt;addable == 13) {
				regalloc(&amp;nod, r, nn);
				cgen(r, &amp;nod);
				gopcode(o, l-&gt;type, &amp;nod, l);
				regfree(&amp;nod);
			} else
				gopcode(o, l-&gt;type, r, l);
			goto com;
		}
		if(l-&gt;complex &gt;= r-&gt;complex) {
			regalloc(&amp;nod, l, nn);
			cgen(l, &amp;nod);
			if(r-&gt;addable &lt; INDEXED) {
				regalloc(&amp;nod1, r, Z);
				cgen(r, &amp;nod1);
				gopcode(o, l-&gt;type, &amp;nod, &amp;nod1);
				regfree(&amp;nod1);
			} else
				gopcode(o, l-&gt;type, &amp;nod, r);
			regfree(&amp;nod);
			goto com;
		}
		regalloc(&amp;nod, r, nn);
		cgen(r, &amp;nod);
		if(l-&gt;addable &lt; INDEXED || l-&gt;addable == 13) {
			regalloc(&amp;nod1, l, Z);
			cgen(l, &amp;nod1);
			if(typechlp[l-&gt;type-&gt;etype])
				gopcode(o, types[TINT], &amp;nod1, &amp;nod);
			else
				gopcode(o, l-&gt;type, &amp;nod1, &amp;nod);
			regfree(&amp;nod1);
		} else
			gopcode(o, l-&gt;type, l, &amp;nod);
		regfree(&amp;nod);

	com:
		if(nn != Z) {
			p1 = p;
			gmove(nodconst(1L), nn);
			gbranch(OGOTO);
			p2 = p;
			patch(p1, pc);
			gmove(nodconst(0L), nn);
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
	Node nod0, nod1, nod2, nod3, nod4, *h, *l, *r;
	Type *t;
	int c, v, x;

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

			if(nn-&gt;op == OREGPAIR) {
				loadpair(n, nn);
				break;
			}
			else if(!vaddr(nn, 0)) {
				t = nn-&gt;type;
				nn-&gt;type = types[TLONG];
				reglcgen(&amp;nod1, nn, Z);
				nn-&gt;type = t;

				gmove(lo64(n), &amp;nod1);
				nod1.xoffset += SZ_LONG;
				gmove(hi64(n), &amp;nod1);
				regfree(&amp;nod1);
			}
			else {
				gins(AMOVL, lo64(n), nn);
				nn-&gt;xoffset += SZ_LONG;
				gins(AMOVL, hi64(n), nn);
				nn-&gt;xoffset -= SZ_LONG;
				break;
			}
			break;
		}
		goto copy;

	case ODOT:
		l = n-&gt;left;
		sugen(l, nodrat, l-&gt;type-&gt;width);
		if(nn == Z)
			break;
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
		break;

	case OSTRUCT:
		/*
		 * rewrite so lhs has no fn call
		 */
		if(nn != Z &amp;&amp; side(nn)) {
			nod1 = *n;
			nod1.type = typ(TIND, n-&gt;type);
			regret(&amp;nod2, &amp;nod1);
			lcgen(nn, &amp;nod2);
			regsalloc(&amp;nod0, &amp;nod1);
			cgen(&amp;nod2, &amp;nod0);
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
			nod0.right = nil;

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
			nod0.right = l;

			/* prtree(&amp;nod0, &#34;hand craft&#34;); /* */
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
		h = nn;
		if(nn-&gt;op == OREGPAIR) {
			regsalloc(&amp;nod1, nn);
			nn = &amp;nod1;
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
		if(h-&gt;op == OREGPAIR)
			loadpair(nn-&gt;left, h);
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
	if(nn == Z) {
		switch(n-&gt;op) {
		case OASADD:
		case OASSUB:
		case OASAND:
		case OASOR:
		case OASXOR:

		case OASMUL:
		case OASLMUL:


		case OASASHL:
		case OASASHR:
		case OASLSHR:
			break;

		case OPOSTINC:
		case OPOSTDEC:
		case OPREINC:
		case OPREDEC:
			break;

		default:
			return;
		}
	}

	if(n-&gt;complex &gt;= FNX &amp;&amp; nn != nil &amp;&amp; nn-&gt;complex &gt;= FNX) {
		t = nn-&gt;type;
		nn-&gt;type = types[TLONG];
		regialloc(&amp;nod1, nn, Z);
		lcgen(nn, &amp;nod1);
		regsalloc(&amp;nod2, nn);
		nn-&gt;type = t;

		gins(AMOVL, &amp;nod1, &amp;nod2);
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

	x = 0;
	v = w == 8;
	if(v) {
		c = cursafe;
		if(n-&gt;left != Z &amp;&amp; n-&gt;left-&gt;complex &gt;= FNX
		&amp;&amp; n-&gt;right != Z &amp;&amp; n-&gt;right-&gt;complex &gt;= FNX) {
//			warn(n, &#34;toughie&#34;);
			regsalloc(&amp;nod1, n-&gt;right);
			cgen(n-&gt;right, &amp;nod1);
			nod2 = *n;
			nod2.right = &amp;nod1;
			cgen(&amp;nod2, nn);
			cursafe = c;
			return;
		}
		if(cgen64(n, nn)) {
			cursafe = c;
			return;
		}
		if(n-&gt;op == OCOM) {
			n = n-&gt;left;
			x = 1;
		}
	}

	/* botch, need to save in .safe */
	c = 0;
	if(n-&gt;complex &gt; nn-&gt;complex) {
		t = n-&gt;type;
		n-&gt;type = types[TLONG];
		if(v) {
			regalloc(&amp;nod0, n, Z);
			if(!vaddr(n, 0)) {
				reglcgen(&amp;nod1, n, Z);
				n-&gt;type = t;
				n = &amp;nod1;
			}
			else
				n-&gt;type = t;
		}
		else {
			nodreg(&amp;nod1, n, D_SI);
			if(reg[D_SI]) {
				gins(APUSHL, &amp;nod1, Z);
				c |= 1;
				reg[D_SI]++;
			}
			lcgen(n, &amp;nod1);
			n-&gt;type = t;
		}

		t = nn-&gt;type;
		nn-&gt;type = types[TLONG];
		if(v) {
			if(!vaddr(nn, 0)) {
				reglcgen(&amp;nod2, nn, Z);
				nn-&gt;type = t;
				nn = &amp;nod2;
			}
			else
				nn-&gt;type = t;
		}
		else {
			nodreg(&amp;nod2, nn, D_DI);
			if(reg[D_DI]) {
				gins(APUSHL, &amp;nod2, Z);
				c |= 2;
				reg[D_DI]++;
			}
			lcgen(nn, &amp;nod2);
			nn-&gt;type = t;
		}
	} else {
		t = nn-&gt;type;
		nn-&gt;type = types[TLONG];
		if(v) {
			regalloc(&amp;nod0, nn, Z);
			if(!vaddr(nn, 0)) {
				reglcgen(&amp;nod2, nn, Z);
				nn-&gt;type = t;
				nn = &amp;nod2;
			}
			else
				nn-&gt;type = t;
		}
		else {
			nodreg(&amp;nod2, nn, D_DI);
			if(reg[D_DI]) {
				gins(APUSHL, &amp;nod2, Z);
				c |= 2;
				reg[D_DI]++;
			}
			lcgen(nn, &amp;nod2);
			nn-&gt;type = t;
		}

		t = n-&gt;type;
		n-&gt;type = types[TLONG];
		if(v) {
			if(!vaddr(n, 0)) {
				reglcgen(&amp;nod1, n, Z);
				n-&gt;type = t;
				n = &amp;nod1;
			}
			else
				n-&gt;type = t;
		}
		else {
			nodreg(&amp;nod1, n, D_SI);
			if(reg[D_SI]) {
				gins(APUSHL, &amp;nod1, Z);
				c |= 1;
				reg[D_SI]++;
			}
			lcgen(n, &amp;nod1);
			n-&gt;type = t;
		}
	}
	if(v) {
		gins(AMOVL, n, &amp;nod0);
		if(x)
			gins(ANOTL, Z, &amp;nod0);
		gins(AMOVL, &amp;nod0, nn);
		n-&gt;xoffset += SZ_LONG;
		nn-&gt;xoffset += SZ_LONG;
		gins(AMOVL, n, &amp;nod0);
		if(x)
			gins(ANOTL, Z, &amp;nod0);
		gins(AMOVL, &amp;nod0, nn);
		n-&gt;xoffset -= SZ_LONG;
		nn-&gt;xoffset -= SZ_LONG;
		if(nn == &amp;nod2)
			regfree(&amp;nod2);
		if(n == &amp;nod1)
			regfree(&amp;nod1);
		regfree(&amp;nod0);
		return;
	}
	nodreg(&amp;nod3, n, D_CX);
	if(reg[D_CX]) {
		gins(APUSHL, &amp;nod3, Z);
		c |= 4;
		reg[D_CX]++;
	}
	gins(AMOVL, nodconst(w/SZ_LONG), &amp;nod3);
	gins(ACLD, Z, Z);
	gins(AREP, Z, Z);
	gins(AMOVSL, Z, Z);
	if(c &amp; 4) {
		gins(APOPL, Z, &amp;nod3);
		reg[D_CX]--;
	}
	if(c &amp; 2) {
		gins(APOPL, Z, &amp;nod2);
		reg[nod2.reg]--;
	}
	if(c &amp; 1) {
		gins(APOPL, Z, &amp;nod1);
		reg[nod1.reg]--;
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/8g/cgen64.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/8g/cgen64.c</h1>

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
 * attempt to generate 64-bit
 *	res = n
 * return 1 on success, 0 if op not handled.
 */
void
cgen64(Node *n, Node *res)
{
	Node t1, t2, ax, dx, cx, ex, fx, *l, *r;
	Node lo1, lo2, hi1, hi2;
	Prog *p1, *p2;
	uint64 v;
	uint32 lv, hv;

	if(res-&gt;op != OINDREG &amp;&amp; res-&gt;op != ONAME) {
		dump(&#34;n&#34;, n);
		dump(&#34;res&#34;, res);
		fatal(&#34;cgen64 %O of %O&#34;, n-&gt;op, res-&gt;op);
	}
	switch(n-&gt;op) {
	default:
		fatal(&#34;cgen64 %O&#34;, n-&gt;op);

	case OMINUS:
		cgen(n-&gt;left, res);
		split64(res, &amp;lo1, &amp;hi1);
		gins(ANEGL, N, &amp;lo1);
		gins(AADCL, ncon(0), &amp;hi1);
		gins(ANEGL, N, &amp;hi1);
		splitclean();
		return;

	case OCOM:
		cgen(n-&gt;left, res);
		split64(res, &amp;lo1, &amp;hi1);
		gins(ANOTL, N, &amp;lo1);
		gins(ANOTL, N, &amp;hi1);
		splitclean();
		return;

	case OADD:
	case OSUB:
	case OMUL:
	case OLSH:
	case ORSH:
	case OAND:
	case OOR:
	case OXOR:
		// binary operators.
		// common setup below.
		break;
	}

	l = n-&gt;left;
	r = n-&gt;right;
	if(!l-&gt;addable) {
		tempalloc(&amp;t1, l-&gt;type);
		cgen(l, &amp;t1);
		l = &amp;t1;
	}
	if(r != N &amp;&amp; !r-&gt;addable) {
		tempalloc(&amp;t2, r-&gt;type);
		cgen(r, &amp;t2);
		r = &amp;t2;
	}

	nodreg(&amp;ax, types[TINT32], D_AX);
	nodreg(&amp;cx, types[TINT32], D_CX);
	nodreg(&amp;dx, types[TINT32], D_DX);

	// Setup for binary operation.
	split64(l, &amp;lo1, &amp;hi1);
	if(is64(r-&gt;type))
		split64(r, &amp;lo2, &amp;hi2);

	// Do op.  Leave result in DX:AX.
	switch(n-&gt;op) {
	case OADD:
		// TODO: Constants
		gins(AMOVL, &amp;lo1, &amp;ax);
		gins(AMOVL, &amp;hi1, &amp;dx);
		gins(AADDL, &amp;lo2, &amp;ax);
		gins(AADCL, &amp;hi2, &amp;dx);
		break;

	case OSUB:
		// TODO: Constants.
		gins(AMOVL, &amp;lo1, &amp;ax);
		gins(AMOVL, &amp;hi1, &amp;dx);
		gins(ASUBL, &amp;lo2, &amp;ax);
		gins(ASBBL, &amp;hi2, &amp;dx);
		break;

	case OMUL:
		// let&#39;s call the next two EX and FX.
		regalloc(&amp;ex, types[TPTR32], N);
		regalloc(&amp;fx, types[TPTR32], N);

		// load args into DX:AX and EX:CX.
		gins(AMOVL, &amp;lo1, &amp;ax);
		gins(AMOVL, &amp;hi1, &amp;dx);
		gins(AMOVL, &amp;lo2, &amp;cx);
		gins(AMOVL, &amp;hi2, &amp;ex);

		// if DX and EX are zero, use 32 x 32 -&gt; 64 unsigned multiply.
		gins(AMOVL, &amp;dx, &amp;fx);
		gins(AORL, &amp;ex, &amp;fx);
		p1 = gbranch(AJNE, T);
		gins(AMULL, &amp;cx, N);	// implicit &amp;ax
		p2 = gbranch(AJMP, T);
		patch(p1, pc);

		// full 64x64 -&gt; 64, from 32x32 -&gt; 64.
		gins(AIMULL, &amp;cx, &amp;dx);
		gins(AMOVL, &amp;ax, &amp;fx);
		gins(AIMULL, &amp;ex, &amp;fx);
		gins(AADDL, &amp;dx, &amp;fx);
		gins(AMOVL, &amp;cx, &amp;dx);
		gins(AMULL, &amp;dx, N);	// implicit &amp;ax
		gins(AADDL, &amp;fx, &amp;dx);
		patch(p2, pc);

		regfree(&amp;ex);
		regfree(&amp;fx);
		break;

	case OLSH:
		if(r-&gt;op == OLITERAL) {
			v = mpgetfix(r-&gt;val.u.xval);
			if(v &gt;= 64) {
				if(is64(r-&gt;type))
					splitclean();
				splitclean();
				split64(res, &amp;lo2, &amp;hi2);
				gins(AMOVL, ncon(0), &amp;lo2);
				gins(AMOVL, ncon(0), &amp;hi2);
				splitclean();
				goto out;
			}
			if(v &gt;= 32) {
				if(is64(r-&gt;type))
					splitclean();
				split64(res, &amp;lo2, &amp;hi2);
				gmove(&amp;lo1, &amp;hi2);
				if(v &gt; 32) {
					gins(ASHLL, ncon(v - 32), &amp;hi2);
				}
				gins(AMOVL, ncon(0), &amp;lo2);
				splitclean();
				splitclean();
				goto out;
			}

			// general shift
			gins(AMOVL, &amp;lo1, &amp;ax);
			gins(AMOVL, &amp;hi1, &amp;dx);
			p1 = gins(ASHLL, ncon(v), &amp;dx);
			p1-&gt;from.index = D_AX;	// double-width shift
			p1-&gt;from.scale = 0;
			gins(ASHLL, ncon(v), &amp;ax);
			break;
		}

		// load value into DX:AX.
		gins(AMOVL, &amp;lo1, &amp;ax);
		gins(AMOVL, &amp;hi1, &amp;dx);

		// load shift value into register.
		// if high bits are set, zero value.
		p1 = P;
		if(is64(r-&gt;type)) {
			gins(ACMPL, &amp;hi2, ncon(0));
			p1 = gbranch(AJNE, T);
			gins(AMOVL, &amp;lo2, &amp;cx);
		} else {
			cx.type = types[TUINT32];
			gmove(r, &amp;cx);
		}

		// if shift count is &gt;=64, zero value
		gins(ACMPL, &amp;cx, ncon(64));
		p2 = gbranch(optoas(OLT, types[TUINT32]), T);
		if(p1 != P)
			patch(p1, pc);
		gins(AXORL, &amp;dx, &amp;dx);
		gins(AXORL, &amp;ax, &amp;ax);
		patch(p2, pc);

		// if shift count is &gt;= 32, zero low.
		gins(ACMPL, &amp;cx, ncon(32));
		p1 = gbranch(optoas(OLT, types[TUINT32]), T);
		gins(AMOVL, &amp;ax, &amp;dx);
		gins(ASHLL, &amp;cx, &amp;dx);	// SHLL only uses bottom 5 bits of count
		gins(AXORL, &amp;ax, &amp;ax);
		p2 = gbranch(AJMP, T);
		patch(p1, pc);

		// general shift
		p1 = gins(ASHLL, &amp;cx, &amp;dx);
		p1-&gt;from.index = D_AX;	// double-width shift
		p1-&gt;from.scale = 0;
		gins(ASHLL, &amp;cx, &amp;ax);
		patch(p2, pc);
		break;

	case ORSH:
		if(r-&gt;op == OLITERAL) {
			v = mpgetfix(r-&gt;val.u.xval);
			if(v &gt;= 64) {
				if(is64(r-&gt;type))
					splitclean();
				splitclean();
				split64(res, &amp;lo2, &amp;hi2);
				if(hi1.type-&gt;etype == TINT32) {
					gmove(&amp;hi1, &amp;lo2);
					gins(ASARL, ncon(31), &amp;lo2);
					gmove(&amp;hi1, &amp;hi2);
					gins(ASARL, ncon(31), &amp;hi2);
				} else {
					gins(AMOVL, ncon(0), &amp;lo2);
					gins(AMOVL, ncon(0), &amp;hi2);
				}
				splitclean();
				goto out;
			}
			if(v &gt;= 32) {
				if(is64(r-&gt;type))
					splitclean();
				split64(res, &amp;lo2, &amp;hi2);
				gmove(&amp;hi1, &amp;lo2);
				if(v &gt; 32)
					gins(optoas(ORSH, hi1.type), ncon(v-32), &amp;lo2);
				if(hi1.type-&gt;etype == TINT32) {
					gmove(&amp;hi1, &amp;hi2);
					gins(ASARL, ncon(31), &amp;hi2);
				} else
					gins(AMOVL, ncon(0), &amp;hi2);
				splitclean();
				splitclean();
				goto out;
			}

			// general shift
			gins(AMOVL, &amp;lo1, &amp;ax);
			gins(AMOVL, &amp;hi1, &amp;dx);
			p1 = gins(ASHRL, ncon(v), &amp;ax);
			p1-&gt;from.index = D_DX;	// double-width shift
			p1-&gt;from.scale = 0;
			gins(optoas(ORSH, hi1.type), ncon(v), &amp;dx);
			break;
		}

		// load value into DX:AX.
		gins(AMOVL, &amp;lo1, &amp;ax);
		gins(AMOVL, &amp;hi1, &amp;dx);

		// load shift value into register.
		// if high bits are set, zero value.
		p1 = P;
		if(is64(r-&gt;type)) {
			gins(ACMPL, &amp;hi2, ncon(0));
			p1 = gbranch(AJNE, T);
			gins(AMOVL, &amp;lo2, &amp;cx);
		} else {
			cx.type = types[TUINT32];
			gmove(r, &amp;cx);
		}

		// if shift count is &gt;=64, zero or sign-extend value
		gins(ACMPL, &amp;cx, ncon(64));
		p2 = gbranch(optoas(OLT, types[TUINT32]), T);
		if(p1 != P)
			patch(p1, pc);
		if(hi1.type-&gt;etype == TINT32) {
			gins(ASARL, ncon(31), &amp;dx);
			gins(AMOVL, &amp;dx, &amp;ax);
		} else {
			gins(AXORL, &amp;dx, &amp;dx);
			gins(AXORL, &amp;ax, &amp;ax);
		}
		patch(p2, pc);

		// if shift count is &gt;= 32, sign-extend hi.
		gins(ACMPL, &amp;cx, ncon(32));
		p1 = gbranch(optoas(OLT, types[TUINT32]), T);
		gins(AMOVL, &amp;dx, &amp;ax);
		if(hi1.type-&gt;etype == TINT32) {
			gins(ASARL, &amp;cx, &amp;ax);	// SARL only uses bottom 5 bits of count
			gins(ASARL, ncon(31), &amp;dx);
		} else {
			gins(ASHRL, &amp;cx, &amp;ax);
			gins(AXORL, &amp;dx, &amp;dx);
		}
		p2 = gbranch(AJMP, T);
		patch(p1, pc);

		// general shift
		p1 = gins(ASHRL, &amp;cx, &amp;ax);
		p1-&gt;from.index = D_DX;	// double-width shift
		p1-&gt;from.scale = 0;
		gins(optoas(ORSH, hi1.type), &amp;cx, &amp;dx);
		patch(p2, pc);
		break;

	case OXOR:
	case OAND:
	case OOR:
		// make constant the right side (it usually is anyway).
		if(lo1.op == OLITERAL) {
			nswap(&amp;lo1, &amp;lo2);
			nswap(&amp;hi1, &amp;hi2);
		}
		if(lo2.op == OLITERAL) {
			// special cases for constants.
			lv = mpgetfix(lo2.val.u.xval);
			hv = mpgetfix(hi2.val.u.xval);
			splitclean();	// right side
			split64(res, &amp;lo2, &amp;hi2);
			switch(n-&gt;op) {
			case OXOR:
				gmove(&amp;lo1, &amp;lo2);
				gmove(&amp;hi1, &amp;hi2);
				switch(lv) {
				case 0:
					break;
				case 0xffffffffu:
					gins(ANOTL, N, &amp;lo2);
					break;
				default:
					gins(AXORL, ncon(lv), &amp;lo2);
					break;
				}
				switch(hv) {
				case 0:
					break;
				case 0xffffffffu:
					gins(ANOTL, N, &amp;hi2);
					break;
				default:
					gins(AXORL, ncon(hv), &amp;hi2);
					break;
				}
				break;

			case OAND:
				switch(lv) {
				case 0:
					gins(AMOVL, ncon(0), &amp;lo2);
					break;
				default:
					gmove(&amp;lo1, &amp;lo2);
					if(lv != 0xffffffffu)
						gins(AANDL, ncon(lv), &amp;lo2);
					break;
				}
				switch(hv) {
				case 0:
					gins(AMOVL, ncon(0), &amp;hi2);
					break;
				default:
					gmove(&amp;hi1, &amp;hi2);
					if(hv != 0xffffffffu)
						gins(AANDL, ncon(hv), &amp;hi2);
					break;
				}
				break;

			case OOR:
				switch(lv) {
				case 0:
					gmove(&amp;lo1, &amp;lo2);
					break;
				case 0xffffffffu:
					gins(AMOVL, ncon(0xffffffffu), &amp;lo2);
					break;
				default:
					gmove(&amp;lo1, &amp;lo2);
					gins(AORL, ncon(lv), &amp;lo2);
					break;
				}
				switch(hv) {
				case 0:
					gmove(&amp;hi1, &amp;hi2);
					break;
				case 0xffffffffu:
					gins(AMOVL, ncon(0xffffffffu), &amp;hi2);
					break;
				default:
					gmove(&amp;hi1, &amp;hi2);
					gins(AORL, ncon(hv), &amp;hi2);
					break;
				}
				break;
			}
			splitclean();
			splitclean();
			goto out;
		}
		gins(AMOVL, &amp;lo1, &amp;ax);
		gins(AMOVL, &amp;hi1, &amp;dx);
		gins(optoas(n-&gt;op, lo1.type), &amp;lo2, &amp;ax);
		gins(optoas(n-&gt;op, lo1.type), &amp;hi2, &amp;dx);
		break;
	}
	if(is64(r-&gt;type))
		splitclean();
	splitclean();

	split64(res, &amp;lo1, &amp;hi1);
	gins(AMOVL, &amp;ax, &amp;lo1);
	gins(AMOVL, &amp;dx, &amp;hi1);
	splitclean();

out:
	if(r == &amp;t2)
		tempfree(&amp;t2);
	if(l == &amp;t1)
		tempfree(&amp;t1);
}

/*
 * generate comparison of nl, nr, both 64-bit.
 * nl is memory; nr is constant or memory.
 */
void
cmp64(Node *nl, Node *nr, int op, Prog *to)
{
	Node lo1, hi1, lo2, hi2, rr;
	Prog *br;
	Type *t;

	split64(nl, &amp;lo1, &amp;hi1);
	split64(nr, &amp;lo2, &amp;hi2);

	// compare most significant word;
	// if they differ, we&#39;re done.
	t = hi1.type;
	if(nl-&gt;op == OLITERAL || nr-&gt;op == OLITERAL)
		gins(ACMPL, &amp;hi1, &amp;hi2);
	else {
		regalloc(&amp;rr, types[TINT32], N);
		gins(AMOVL, &amp;hi1, &amp;rr);
		gins(ACMPL, &amp;rr, &amp;hi2);
		regfree(&amp;rr);
	}
	br = P;
	switch(op) {
	default:
		fatal(&#34;cmp64 %O %T&#34;, op, t);
	case OEQ:
		// cmp hi
		// jne L
		// cmp lo
		// jeq to
		// L:
		br = gbranch(AJNE, T);
		break;
	case ONE:
		// cmp hi
		// jne to
		// cmp lo
		// jne to
		patch(gbranch(AJNE, T), to);
		break;
	case OGE:
	case OGT:
		// cmp hi
		// jgt to
		// jlt L
		// cmp lo
		// jge to (or jgt to)
		// L:
		patch(gbranch(optoas(OGT, t), T), to);
		br = gbranch(optoas(OLT, t), T);
		break;
	case OLE:
	case OLT:
		// cmp hi
		// jlt to
		// jgt L
		// cmp lo
		// jle to (or jlt to)
		// L:
		patch(gbranch(optoas(OLT, t), T), to);
		br = gbranch(optoas(OGT, t), T);
		break;
	}

	// compare least significant word
	t = lo1.type;
	if(nl-&gt;op == OLITERAL || nr-&gt;op == OLITERAL)
		gins(ACMPL, &amp;lo1, &amp;lo2);
	else {
		regalloc(&amp;rr, types[TINT32], N);
		gins(AMOVL, &amp;lo1, &amp;rr);
		gins(ACMPL, &amp;rr, &amp;lo2);
		regfree(&amp;rr);
	}

	// jump again
	patch(gbranch(optoas(op, t), T), to);

	// point first branch down here if appropriate
	if(br != P)
		patch(br, pc);

	splitclean();
	splitclean();
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

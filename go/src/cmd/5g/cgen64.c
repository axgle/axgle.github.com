<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5g/cgen64.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/5g/cgen64.c</h1>

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
	Node t1, t2, *l, *r;
	Node lo1, lo2, hi1, hi2;
	Node al, ah, bl, bh, cl, ch, s, n1, creg;
	Prog *p1, *p2, *p3, *p4, *p5, *p6;

	uint64 v;

	if(res-&gt;op != OINDREG &amp;&amp; res-&gt;op != ONAME) {
		dump(&#34;n&#34;, n);
		dump(&#34;res&#34;, res);
		fatal(&#34;cgen64 %O of %O&#34;, n-&gt;op, res-&gt;op);
	}

	l = n-&gt;left;
	if(!l-&gt;addable) {
		tempname(&amp;t1, l-&gt;type);
		cgen(l, &amp;t1);
		l = &amp;t1;
	}

	split64(l, &amp;lo1, &amp;hi1);
	switch(n-&gt;op) {
	default:
		fatal(&#34;cgen64 %O&#34;, n-&gt;op);

	case OMINUS:
		split64(res, &amp;lo2, &amp;hi2);

		regalloc(&amp;t1, lo1.type, N);
		regalloc(&amp;al, lo1.type, N);
		regalloc(&amp;ah, hi1.type, N);

		gins(AMOVW, &amp;lo1, &amp;al);
		gins(AMOVW, &amp;hi1, &amp;ah);

		gmove(ncon(0), &amp;t1);
		p1 = gins(ASUB, &amp;al, &amp;t1);
		p1-&gt;scond |= C_SBIT;
		gins(AMOVW, &amp;t1, &amp;lo2);

		gmove(ncon(0), &amp;t1);
		gins(ASBC, &amp;ah, &amp;t1);
		gins(AMOVW, &amp;t1, &amp;hi2);

		regfree(&amp;t1);
		regfree(&amp;al);
		regfree(&amp;ah);
		splitclean();
		splitclean();
		return;

	case OCOM:
		split64(res, &amp;lo2, &amp;hi2);
		regalloc(&amp;n1, lo1.type, N);

		gins(AMOVW, &amp;lo1, &amp;n1);
		gins(AMVN, &amp;n1, &amp;n1);
		gins(AMOVW, &amp;n1, &amp;lo2);

		gins(AMOVW, &amp;hi1, &amp;n1);
		gins(AMVN, &amp;n1, &amp;n1);
		gins(AMOVW, &amp;n1, &amp;hi2);

		regfree(&amp;n1);
		splitclean();
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

	// setup for binary operators
	r = n-&gt;right;
	if(r != N &amp;&amp; !r-&gt;addable) {
		tempname(&amp;t2, r-&gt;type);
		cgen(r, &amp;t2);
		r = &amp;t2;
	}
	if(is64(r-&gt;type))
		split64(r, &amp;lo2, &amp;hi2);

	regalloc(&amp;al, lo1.type, N);
	regalloc(&amp;ah, hi1.type, N);

	// Do op.  Leave result in ah:al.
	switch(n-&gt;op) {
	default:
		fatal(&#34;cgen64: not implemented: %N\n&#34;, n);

	case OADD:
		// TODO: Constants
		regalloc(&amp;bl, types[TPTR32], N);
		regalloc(&amp;bh, types[TPTR32], N);
		gins(AMOVW, &amp;hi1, &amp;ah);
		gins(AMOVW, &amp;lo1, &amp;al);
		gins(AMOVW, &amp;hi2, &amp;bh);
		gins(AMOVW, &amp;lo2, &amp;bl);
		p1 = gins(AADD, &amp;bl, &amp;al);
		p1-&gt;scond |= C_SBIT;
		gins(AADC, &amp;bh, &amp;ah);
		regfree(&amp;bl);
		regfree(&amp;bh);
		break;

	case OSUB:
		// TODO: Constants.
		regalloc(&amp;bl, types[TPTR32], N);
		regalloc(&amp;bh, types[TPTR32], N);
		gins(AMOVW, &amp;lo1, &amp;al);
		gins(AMOVW, &amp;hi1, &amp;ah);
		gins(AMOVW, &amp;lo2, &amp;bl);
		gins(AMOVW, &amp;hi2, &amp;bh);
		p1 = gins(ASUB, &amp;bl, &amp;al);
		p1-&gt;scond |= C_SBIT;
		gins(ASBC, &amp;bh, &amp;ah);
		regfree(&amp;bl);
		regfree(&amp;bh);
		break;

	case OMUL:
		// TODO(kaib): this can be done with 4 regs and does not need 6
		regalloc(&amp;bl, types[TPTR32], N);
		regalloc(&amp;bh, types[TPTR32], N);
		regalloc(&amp;cl, types[TPTR32], N);
		regalloc(&amp;ch, types[TPTR32], N);

		// load args into bh:bl and bh:bl.
		gins(AMOVW, &amp;hi1, &amp;bh);
		gins(AMOVW, &amp;lo1, &amp;bl);
		gins(AMOVW, &amp;hi2, &amp;ch);
		gins(AMOVW, &amp;lo2, &amp;cl);

		// bl * cl
		p1 = gins(AMULLU, N, N);
		p1-&gt;from.type = D_REG;
		p1-&gt;from.reg = bl.val.u.reg;
		p1-&gt;reg = cl.val.u.reg;
		p1-&gt;to.type = D_REGREG;
		p1-&gt;to.reg = ah.val.u.reg;
		p1-&gt;to.offset = al.val.u.reg;
//print(&#34;%P\n&#34;, p1);

		// bl * ch
		p1 = gins(AMULA, N, N);
		p1-&gt;from.type = D_REG;
		p1-&gt;from.reg = bl.val.u.reg;
		p1-&gt;reg = ch.val.u.reg;
		p1-&gt;to.type = D_REGREG;
		p1-&gt;to.reg = ah.val.u.reg;
		p1-&gt;to.offset = ah.val.u.reg;
//print(&#34;%P\n&#34;, p1);

		// bh * cl
		p1 = gins(AMULA, N, N);
		p1-&gt;from.type = D_REG;
		p1-&gt;from.reg = bh.val.u.reg;
		p1-&gt;reg = cl.val.u.reg;
		p1-&gt;to.type = D_REGREG;
		p1-&gt;to.reg = ah.val.u.reg;
		p1-&gt;to.offset = ah.val.u.reg;
//print(&#34;%P\n&#34;, p1);

		regfree(&amp;bh);
		regfree(&amp;bl);
		regfree(&amp;ch);
		regfree(&amp;cl);

		break;

	case OLSH:
		regalloc(&amp;bl, lo1.type, N);
		regalloc(&amp;bh, hi1.type, N);
		gins(AMOVW, &amp;hi1, &amp;bh);
		gins(AMOVW, &amp;lo1, &amp;bl);

		if(r-&gt;op == OLITERAL) {
			v = mpgetfix(r-&gt;val.u.xval);
			if(v &gt;= 64) {
				// TODO(kaib): replace with gins(AMOVW, nodintconst(0), &amp;al)
				// here and below (verify it optimizes to EOR)
				gins(AEOR, &amp;al, &amp;al);
				gins(AEOR, &amp;ah, &amp;ah);
			} else if(v &gt; 32) {
				gins(AEOR, &amp;al, &amp;al);
				//	MOVW	bl&lt;&lt;(v-32), ah
				gshift(AMOVW, &amp;bl, SHIFT_LL, (v-32), &amp;ah);
			} else if(v == 32) {
				gins(AEOR, &amp;al, &amp;al);
				gins(AMOVW, &amp;bl, &amp;ah);
			} else if(v &gt; 0) {
				//	MOVW	bl&lt;&lt;v, al
				gshift(AMOVW, &amp;bl, SHIFT_LL, v, &amp;al);

				//	MOVW	bh&lt;&lt;v, ah
				gshift(AMOVW, &amp;bh, SHIFT_LL, v, &amp;ah);

				//	OR		bl&gt;&gt;(32-v), ah
				gshift(AORR, &amp;bl, SHIFT_LR, 32-v, &amp;ah);
			} else {
				gins(AMOVW, &amp;bl, &amp;al);
				gins(AMOVW, &amp;bh, &amp;ah);
			}
			goto olsh_break;
		}

		regalloc(&amp;s, types[TUINT32], N);
		regalloc(&amp;creg, types[TUINT32], N);
		if (is64(r-&gt;type)) {
			// shift is &gt;= 1&lt;&lt;32
			split64(r, &amp;cl, &amp;ch);
			gmove(&amp;ch, &amp;s);
			p1 = gins(AMOVW, &amp;s, &amp;s);
			p1-&gt;scond |= C_SBIT;
			p6 = gbranch(ABNE, T);
			gmove(&amp;cl, &amp;s);
			splitclean();
		} else {
			gmove(r, &amp;s);
			p6 = P;
		}
		p1 = gins(AMOVW, &amp;s, &amp;s);
		p1-&gt;scond |= C_SBIT;

		// shift == 0
		p1 = gins(AMOVW, &amp;bl, &amp;al);
		p1-&gt;scond = C_SCOND_EQ;
		p1 = gins(AMOVW, &amp;bh, &amp;ah);
		p1-&gt;scond = C_SCOND_EQ;
		p2 = gbranch(ABEQ, T);

		// shift is &lt; 32
		nodconst(&amp;n1, types[TUINT32], 32);
		gmove(&amp;n1, &amp;creg);
		gcmp(ACMP, &amp;s, &amp;creg);

		//	MOVW.LO		bl&lt;&lt;s, al
		p1 = gregshift(AMOVW, &amp;bl, SHIFT_LL, &amp;s, &amp;al);
		p1-&gt;scond = C_SCOND_LO;

		//	MOVW.LO		bh&lt;&lt;s, ah
		p1 = gregshift(AMOVW, &amp;bh, SHIFT_LL, &amp;s, &amp;ah);
		p1-&gt;scond = C_SCOND_LO;

		//	SUB.LO		s, creg
		p1 = gins(ASUB, &amp;s, &amp;creg);
		p1-&gt;scond = C_SCOND_LO;

		//	OR.LO		bl&gt;&gt;creg, ah
		p1 = gregshift(AORR, &amp;bl, SHIFT_LR, &amp;creg, &amp;ah);
		p1-&gt;scond = C_SCOND_LO;

		//	BLO	end
		p3 = gbranch(ABLO, T);

		// shift == 32
		p1 = gins(AEOR, &amp;al, &amp;al);
		p1-&gt;scond = C_SCOND_EQ;
		p1 = gins(AMOVW, &amp;bl, &amp;ah);
		p1-&gt;scond = C_SCOND_EQ;
		p4 = gbranch(ABEQ, T);

		// shift is &lt; 64
		nodconst(&amp;n1, types[TUINT32], 64);
		gmove(&amp;n1, &amp;creg);
		gcmp(ACMP, &amp;s, &amp;creg);

		//	EOR.LO	al, al
		p1 = gins(AEOR, &amp;al, &amp;al);
		p1-&gt;scond = C_SCOND_LO;

		//	MOVW.LO		creg&gt;&gt;1, creg
		p1 = gshift(AMOVW, &amp;creg, SHIFT_LR, 1, &amp;creg);
		p1-&gt;scond = C_SCOND_LO;

		//	SUB.LO		creg, s
		p1 = gins(ASUB, &amp;creg, &amp;s);
		p1-&gt;scond = C_SCOND_LO;

		//	MOVW	bl&lt;&lt;s, ah
		p1 = gregshift(AMOVW, &amp;bl, SHIFT_LL, &amp;s, &amp;ah);
		p1-&gt;scond = C_SCOND_LO;

		p5 = gbranch(ABLO, T);

		// shift &gt;= 64
		if (p6 != P) patch(p6, pc);
		gins(AEOR, &amp;al, &amp;al);
		gins(AEOR, &amp;ah, &amp;ah);

		patch(p2, pc);
		patch(p3, pc);
		patch(p4, pc);
		patch(p5, pc);
		regfree(&amp;s);
		regfree(&amp;creg);

olsh_break:
		regfree(&amp;bl);
		regfree(&amp;bh);
		break;


	case ORSH:
		regalloc(&amp;bl, lo1.type, N);
		regalloc(&amp;bh, hi1.type, N);
		gins(AMOVW, &amp;hi1, &amp;bh);
		gins(AMOVW, &amp;lo1, &amp;bl);

		if(r-&gt;op == OLITERAL) {
			v = mpgetfix(r-&gt;val.u.xval);
			if(v &gt;= 64) {
				if(bh.type-&gt;etype == TINT32) {
					//	MOVW	bh-&gt;31, al
					gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;al);

					//	MOVW	bh-&gt;31, ah
					gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;ah);
				} else {
					gins(AEOR, &amp;al, &amp;al);
					gins(AEOR, &amp;ah, &amp;ah);
				}
			} else if(v &gt; 32) {
				if(bh.type-&gt;etype == TINT32) {
					//	MOVW	bh-&gt;(v-32), al
					gshift(AMOVW, &amp;bh, SHIFT_AR, v-32, &amp;al);

					//	MOVW	bh-&gt;31, ah
					gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;ah);
				} else {
					//	MOVW	bh&gt;&gt;(v-32), al
					gshift(AMOVW, &amp;bh, SHIFT_LR, v-32, &amp;al);
					gins(AEOR, &amp;ah, &amp;ah);
				}
			} else if(v == 32) {
				gins(AMOVW, &amp;bh, &amp;al);
				if(bh.type-&gt;etype == TINT32) {
					//	MOVW	bh-&gt;31, ah
					gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;ah);
				} else {
					gins(AEOR, &amp;ah, &amp;ah);
				}
			} else if( v &gt; 0) {
				//	MOVW	bl&gt;&gt;v, al
				gshift(AMOVW, &amp;bl, SHIFT_LR, v, &amp;al);
	
				//	OR		bh&lt;&lt;(32-v), al
				gshift(AORR, &amp;bh, SHIFT_LL, 32-v, &amp;al);

				if(bh.type-&gt;etype == TINT32) {
					//	MOVW	bh-&gt;v, ah
					gshift(AMOVW, &amp;bh, SHIFT_AR, v, &amp;ah);
				} else {
					//	MOVW	bh&gt;&gt;v, ah
					gshift(AMOVW, &amp;bh, SHIFT_LR, v, &amp;ah);
				}
			} else {
				gins(AMOVW, &amp;bl, &amp;al);
				gins(AMOVW, &amp;bh, &amp;ah);
			}
			goto orsh_break;
		}

		regalloc(&amp;s, types[TUINT32], N);
		regalloc(&amp;creg, types[TUINT32], N);
		if (is64(r-&gt;type)) {
			// shift is &gt;= 1&lt;&lt;32
			split64(r, &amp;cl, &amp;ch);
			gmove(&amp;ch, &amp;s);
			p1 = gins(AMOVW, &amp;s, &amp;s);
			p1-&gt;scond |= C_SBIT;
			p6 = gbranch(ABNE, T);
			gmove(&amp;cl, &amp;s);
			splitclean();
		} else {
			gmove(r, &amp;s);
			p6 = P;
		}
		p1 = gins(AMOVW, &amp;s, &amp;s);
		p1-&gt;scond |= C_SBIT;

		// shift == 0
		p1 = gins(AMOVW, &amp;bl, &amp;al);
		p1-&gt;scond = C_SCOND_EQ;
		p1 = gins(AMOVW, &amp;bh, &amp;ah);
		p1-&gt;scond = C_SCOND_EQ;
		p2 = gbranch(ABEQ, T);

		// check if shift is &lt; 32
		nodconst(&amp;n1, types[TUINT32], 32);
		gmove(&amp;n1, &amp;creg);
		gcmp(ACMP, &amp;s, &amp;creg);

		//	MOVW.LO		bl&gt;&gt;s, al
		p1 = gregshift(AMOVW, &amp;bl, SHIFT_LR, &amp;s, &amp;al);
		p1-&gt;scond = C_SCOND_LO;

		//	SUB.LO		s,creg
		p1 = gins(ASUB, &amp;s, &amp;creg);
		p1-&gt;scond = C_SCOND_LO;

		//	OR.LO		bh&lt;&lt;(32-s), al
		p1 = gregshift(AORR, &amp;bh, SHIFT_LL, &amp;creg, &amp;al);
		p1-&gt;scond = C_SCOND_LO;

		if(bh.type-&gt;etype == TINT32) {
			//	MOVW	bh-&gt;s, ah
			p1 = gregshift(AMOVW, &amp;bh, SHIFT_AR, &amp;s, &amp;ah);
		} else {
			//	MOVW	bh&gt;&gt;s, ah
			p1 = gregshift(AMOVW, &amp;bh, SHIFT_LR, &amp;s, &amp;ah);
		}
		p1-&gt;scond = C_SCOND_LO;

		//	BLO	end
		p3 = gbranch(ABLO, T);

		// shift == 32
		if(bh.type-&gt;etype == TINT32)
			p1 = gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;ah);
		else
			p1 = gins(AEOR, &amp;al, &amp;al);
		p1-&gt;scond = C_SCOND_EQ;
		p1 = gins(AMOVW, &amp;bh, &amp;al);
		p1-&gt;scond = C_SCOND_EQ;
		p4 = gbranch(ABEQ, T);

		// check if shift is &lt; 64
		nodconst(&amp;n1, types[TUINT32], 64);
		gmove(&amp;n1, &amp;creg);
		gcmp(ACMP, &amp;s, &amp;creg);

		//	MOVW.LO		creg&gt;&gt;1, creg
		p1 = gshift(AMOVW, &amp;creg, SHIFT_LR, 1, &amp;creg);
		p1-&gt;scond = C_SCOND_LO;

		//	SUB.LO		creg, s
		p1 = gins(ASUB, &amp;creg, &amp;s);
		p1-&gt;scond = C_SCOND_LO;

		if(bh.type-&gt;etype == TINT32) {
			//	MOVW	bh-&gt;(s-32), al
			p1 = gregshift(AMOVW, &amp;bh, SHIFT_AR, &amp;s, &amp;al);
			p1-&gt;scond = C_SCOND_LO;

			//	MOVW	bh-&gt;31, ah
			p1 = gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;ah);
			p1-&gt;scond = C_SCOND_LO;
		} else {
			//	MOVW	bh&gt;&gt;(v-32), al
			p1 = gregshift(AMOVW, &amp;bh, SHIFT_LR, &amp;s, &amp;al);
			p1-&gt;scond = C_SCOND_LO;

			p1 = gins(AEOR, &amp;ah, &amp;ah);
			p1-&gt;scond = C_SCOND_LO;
		}

		//	BLO	end
		p5 = gbranch(ABLO, T);

		// s &gt;= 64
		if (p6 != P) patch(p6, pc);
		if(bh.type-&gt;etype == TINT32) {
			//	MOVW	bh-&gt;31, al
			gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;al);

			//	MOVW	bh-&gt;31, ah
			gshift(AMOVW, &amp;bh, SHIFT_AR, 31, &amp;ah);
		} else {
			gins(AEOR, &amp;al, &amp;al);
			gins(AEOR, &amp;ah, &amp;ah);
		}

		patch(p2, pc);
		patch(p3, pc);
		patch(p4, pc);
		patch(p5, pc);
		regfree(&amp;s);
		regfree(&amp;creg);


orsh_break:
		regfree(&amp;bl);
		regfree(&amp;bh);
		break;

	case OXOR:
	case OAND:
	case OOR:
		// TODO(kaib): literal optimizations
		// make constant the right side (it usually is anyway).
//		if(lo1.op == OLITERAL) {
//			nswap(&amp;lo1, &amp;lo2);
//			nswap(&amp;hi1, &amp;hi2);
//		}
//		if(lo2.op == OLITERAL) {
//			// special cases for constants.
//			lv = mpgetfix(lo2.val.u.xval);
//			hv = mpgetfix(hi2.val.u.xval);
//			splitclean();	// right side
//			split64(res, &amp;lo2, &amp;hi2);
//			switch(n-&gt;op) {
//			case OXOR:
//				gmove(&amp;lo1, &amp;lo2);
//				gmove(&amp;hi1, &amp;hi2);
//				switch(lv) {
//				case 0:
//					break;
//				case 0xffffffffu:
//					gins(ANOTL, N, &amp;lo2);
//					break;
//				default:
//					gins(AXORL, ncon(lv), &amp;lo2);
//					break;
//				}
//				switch(hv) {
//				case 0:
//					break;
//				case 0xffffffffu:
//					gins(ANOTL, N, &amp;hi2);
//					break;
//				default:
//					gins(AXORL, ncon(hv), &amp;hi2);
//					break;
//				}
//				break;

//			case OAND:
//				switch(lv) {
//				case 0:
//					gins(AMOVL, ncon(0), &amp;lo2);
//					break;
//				default:
//					gmove(&amp;lo1, &amp;lo2);
//					if(lv != 0xffffffffu)
//						gins(AANDL, ncon(lv), &amp;lo2);
//					break;
//				}
//				switch(hv) {
//				case 0:
//					gins(AMOVL, ncon(0), &amp;hi2);
//					break;
//				default:
//					gmove(&amp;hi1, &amp;hi2);
//					if(hv != 0xffffffffu)
//						gins(AANDL, ncon(hv), &amp;hi2);
//					break;
//				}
//				break;

//			case OOR:
//				switch(lv) {
//				case 0:
//					gmove(&amp;lo1, &amp;lo2);
//					break;
//				case 0xffffffffu:
//					gins(AMOVL, ncon(0xffffffffu), &amp;lo2);
//					break;
//				default:
//					gmove(&amp;lo1, &amp;lo2);
//					gins(AORL, ncon(lv), &amp;lo2);
//					break;
//				}
//				switch(hv) {
//				case 0:
//					gmove(&amp;hi1, &amp;hi2);
//					break;
//				case 0xffffffffu:
//					gins(AMOVL, ncon(0xffffffffu), &amp;hi2);
//					break;
//				default:
//					gmove(&amp;hi1, &amp;hi2);
//					gins(AORL, ncon(hv), &amp;hi2);
//					break;
//				}
//				break;
//			}
//			splitclean();
//			splitclean();
//			goto out;
//		}
		regalloc(&amp;n1, lo1.type, N);
		gins(AMOVW, &amp;lo1, &amp;al);
		gins(AMOVW, &amp;hi1, &amp;ah);
		gins(AMOVW, &amp;lo2, &amp;n1);
		gins(optoas(n-&gt;op, lo1.type), &amp;n1, &amp;al);
		gins(AMOVW, &amp;hi2, &amp;n1);
		gins(optoas(n-&gt;op, lo1.type), &amp;n1, &amp;ah);
		regfree(&amp;n1);
		break;
	}
	if(is64(r-&gt;type))
		splitclean();
	splitclean();

	split64(res, &amp;lo1, &amp;hi1);
	gins(AMOVW, &amp;al, &amp;lo1);
	gins(AMOVW, &amp;ah, &amp;hi1);
	splitclean();

//out:
	regfree(&amp;al);
	regfree(&amp;ah);
}

/*
 * generate comparison of nl, nr, both 64-bit.
 * nl is memory; nr is constant or memory.
 */
void
cmp64(Node *nl, Node *nr, int op, Prog *to)
{
	Node lo1, hi1, lo2, hi2, r1, r2;
	Prog *br;
	Type *t;

	split64(nl, &amp;lo1, &amp;hi1);
	split64(nr, &amp;lo2, &amp;hi2);

	// compare most significant word;
	// if they differ, we&#39;re done.
	t = hi1.type;
	regalloc(&amp;r1, types[TINT32], N);
	regalloc(&amp;r2, types[TINT32], N);
	gins(AMOVW, &amp;hi1, &amp;r1);
	gins(AMOVW, &amp;hi2, &amp;r2);
	gcmp(ACMP, &amp;r1, &amp;r2);
	regfree(&amp;r1);
	regfree(&amp;r2);

	br = P;
	switch(op) {
	default:
		fatal(&#34;cmp64 %O %T&#34;, op, t);
	case OEQ:
		// cmp hi
		// bne L
		// cmp lo
		// beq to
		// L:
		br = gbranch(ABNE, T);
		break;
	case ONE:
		// cmp hi
		// bne to
		// cmp lo
		// bne to
		patch(gbranch(ABNE, T), to);
		break;
	case OGE:
	case OGT:
		// cmp hi
		// bgt to
		// blt L
		// cmp lo
		// bge to (or bgt to)
		// L:
		patch(gbranch(optoas(OGT, t), T), to);
		br = gbranch(optoas(OLT, t), T);
		break;
	case OLE:
	case OLT:
		// cmp hi
		// blt to
		// bgt L
		// cmp lo
		// ble to (or jlt to)
		// L:
		patch(gbranch(optoas(OLT, t), T), to);
		br = gbranch(optoas(OGT, t), T);
		break;
	}

	// compare least significant word
	t = lo1.type;
	regalloc(&amp;r1, types[TINT32], N);
	regalloc(&amp;r2, types[TINT32], N);
	gins(AMOVW, &amp;lo1, &amp;r1);
	gins(AMOVW, &amp;lo2, &amp;r2);
	gcmp(ACMP, &amp;r1, &amp;r2);
	regfree(&amp;r1);
	regfree(&amp;r2);

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

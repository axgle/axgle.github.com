<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/fmt/fltfmt.c</title>

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
	<li>Thu Nov 12 15:47:27 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/fmt/fltfmt.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
 * The authors of this software are Rob Pike and Ken Thompson,
 * with contributions from Mike Burrows and Sean Dorward.
 *
 *     Copyright (c) 2002-2006 by Lucent Technologies.
 *     Portions Copyright (c) 2004 Google Inc.
 * 
 * Permission to use, copy, modify, and distribute this software for any
 * purpose without fee is hereby granted, provided that this entire notice
 * is included in all copies of any software which is or includes a copy
 * or modification of this software and in all copies of the supporting
 * documentation for such software.
 * THIS SOFTWARE IS BEING PROVIDED &#34;AS IS&#34;, WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTY.  IN PARTICULAR, NEITHER THE AUTHORS NOR LUCENT TECHNOLOGIES 
 * NOR GOOGLE INC MAKE ANY REPRESENTATION OR WARRANTY OF ANY KIND CONCERNING 
 * THE MERCHANTABILITY OF THIS SOFTWARE OR ITS FITNESS FOR ANY PARTICULAR PURPOSE.
 */

/* Copyright (c) 2002-2006 Lucent Technologies; see LICENSE */
#include &lt;u.h&gt;
#include &lt;errno.h&gt;
#include &lt;libc.h&gt;
#include &#34;fmtdef.h&#34;

enum
{
	FDIGIT	= 30,
	FDEFLT	= 6,
	NSIGNIF	= 17
};

/*
 * first few powers of 10, enough for about 1/2 of the
 * total space for doubles.
 */
static double pows10[] =
{
	  1e0,   1e1,   1e2,   1e3,   1e4,   1e5,   1e6,   1e7,   1e8,   1e9,
	 1e10,  1e11,  1e12,  1e13,  1e14,  1e15,  1e16,  1e17,  1e18,  1e19,
	 1e20,  1e21,  1e22,  1e23,  1e24,  1e25,  1e26,  1e27,  1e28,  1e29,
	 1e30,  1e31,  1e32,  1e33,  1e34,  1e35,  1e36,  1e37,  1e38,  1e39,
	 1e40,  1e41,  1e42,  1e43,  1e44,  1e45,  1e46,  1e47,  1e48,  1e49,
	 1e50,  1e51,  1e52,  1e53,  1e54,  1e55,  1e56,  1e57,  1e58,  1e59,
	 1e60,  1e61,  1e62,  1e63,  1e64,  1e65,  1e66,  1e67,  1e68,  1e69,
	 1e70,  1e71,  1e72,  1e73,  1e74,  1e75,  1e76,  1e77,  1e78,  1e79,
	 1e80,  1e81,  1e82,  1e83,  1e84,  1e85,  1e86,  1e87,  1e88,  1e89,
	 1e90,  1e91,  1e92,  1e93,  1e94,  1e95,  1e96,  1e97,  1e98,  1e99,
	1e100, 1e101, 1e102, 1e103, 1e104, 1e105, 1e106, 1e107, 1e108, 1e109,
	1e110, 1e111, 1e112, 1e113, 1e114, 1e115, 1e116, 1e117, 1e118, 1e119,
	1e120, 1e121, 1e122, 1e123, 1e124, 1e125, 1e126, 1e127, 1e128, 1e129,
	1e130, 1e131, 1e132, 1e133, 1e134, 1e135, 1e136, 1e137, 1e138, 1e139,
	1e140, 1e141, 1e142, 1e143, 1e144, 1e145, 1e146, 1e147, 1e148, 1e149,
	1e150, 1e151, 1e152, 1e153, 1e154, 1e155, 1e156, 1e157, 1e158, 1e159,
};

#undef	pow10
#define	npows10 ((int)(sizeof(pows10)/sizeof(pows10[0])))
#define	pow10(x)  fmtpow10(x)

static double
pow10(int n)
{
	double d;
	int neg;

	neg = 0;
	if(n &lt; 0){
		neg = 1;
		n = -n;
	}

	if(n &lt; npows10)
		d = pows10[n];
	else{
		d = pows10[npows10-1];
		for(;;){
			n -= npows10 - 1;
			if(n &lt; npows10){
				d *= pows10[n];
				break;
			}
			d *= pows10[npows10 - 1];
		}
	}
	if(neg)
		return 1./d;
	return d;
}

/*
 * add 1 to the decimal integer string a of length n.
 * if 99999 overflows into 10000, return 1 to tell caller
 * to move the virtual decimal point.
 */
static int
xadd1(char *a, int n)
{
	char *b;
	int c;

	if(n &lt; 0 || n &gt; NSIGNIF)
		return 0;
	for(b = a+n-1; b &gt;= a; b--) {
		c = *b + 1;
		if(c &lt;= &#39;9&#39;) {
			*b = c;
			return 0;
		}
		*b = &#39;0&#39;;
	}
	/*
	 * need to overflow adding digit.
	 * shift number down and insert 1 at beginning.
	 * decimal is known to be 0s or we wouldn&#39;t
	 * have gotten this far.  (e.g., 99999+1 =&gt; 00000)
	 */
	a[0] = &#39;1&#39;;
	return 1;
}

/*
 * subtract 1 from the decimal integer string a.
 * if 10000 underflows into 09999, make it 99999
 * and return 1 to tell caller to move the virtual
 * decimal point.  this way, xsub1 is inverse of xadd1.
 */
static int
xsub1(char *a, int n)
{
	char *b;
	int c;

	if(n &lt; 0 || n &gt; NSIGNIF)
		return 0;
	for(b = a+n-1; b &gt;= a; b--) {
		c = *b - 1;
		if(c &gt;= &#39;0&#39;) {
			if(c == &#39;0&#39; &amp;&amp; b == a) {
				/*
				 * just zeroed the top digit; shift everyone up.
				 * decimal is known to be 9s or we wouldn&#39;t
				 * have gotten this far.  (e.g., 10000-1 =&gt; 09999)
				 */
				*b = &#39;9&#39;;
				return 1;
			}
			*b = c;
			return 0;
		}
		*b = &#39;9&#39;;
	}
	/*
	 * can&#39;t get here.  the number a is always normalized
	 * so that it has a nonzero first digit.
	 */
	abort();
}

/*
 * format exponent like sprintf(p, &#34;e%+02d&#34;, e)
 */
static void
xfmtexp(char *p, int e, int ucase)
{
	char se[9];
	int i;

	*p++ = ucase ? &#39;E&#39; : &#39;e&#39;;
	if(e &lt; 0) {
		*p++ = &#39;-&#39;;
		e = -e;
	} else
		*p++ = &#39;+&#39;;
	i = 0;
	while(e) {
		se[i++] = e % 10 + &#39;0&#39;;
		e /= 10;
	}
	while(i &lt; 2)
		se[i++] = &#39;0&#39;;
	while(i &gt; 0)
		*p++ = se[--i];
	*p++ = &#39;\0&#39;;
}

/*
 * compute decimal integer m, exp such that:
 *	f = m*10^exp
 *	m is as short as possible with losing exactness
 * assumes special cases (NaN, +Inf, -Inf) have been handled.
 */
static void
xdtoa(double f, char *s, int *exp, int *neg, int *ns)
{
	int c, d, e2, e, ee, i, ndigit, oerrno;
	char tmp[NSIGNIF+10];
	double g;

	oerrno = errno; /* in case strtod smashes errno */

	/*
	 * make f non-negative.
	 */
	*neg = 0;
	if(f &lt; 0) {
		f = -f;
		*neg = 1;
	}

	/*
	 * must handle zero specially.
	 */
	if(f == 0){
		*exp = 0;
		s[0] = &#39;0&#39;;
		s[1] = &#39;\0&#39;;
		*ns = 1;
		return;
	}

	/*
	 * find g,e such that f = g*10^e.
	 * guess 10-exponent using 2-exponent, then fine tune.
	 */
	frexp(f, &amp;e2);
	e = (int)(e2 * .301029995664);
	g = f * pow10(-e);
	while(g &lt; 1) {
		e--;
		g = f * pow10(-e);
	}
	while(g &gt;= 10) {
		e++;
		g = f * pow10(-e);
	}

	/*
	 * convert NSIGNIF digits as a first approximation.
	 */
	for(i=0; i&lt;NSIGNIF; i++) {
		d = (int)g;
		s[i] = d+&#39;0&#39;;
		g = (g-d) * 10;
	}
	s[i] = 0;

	/*
	 * adjust e because s is 314159... not 3.14159...
	 */
	e -= NSIGNIF-1;
	xfmtexp(s+NSIGNIF, e, 0);

	/*
	 * adjust conversion until strtod(s) == f exactly.
	 */
	for(i=0; i&lt;10; i++) {
		g = strtod(s, nil);
		if(f &gt; g) {
			if(xadd1(s, NSIGNIF)) {
				/* gained a digit */
				e--;
				xfmtexp(s+NSIGNIF, e, 0);
			}
			continue;
		}
		if(f &lt; g) {
			if(xsub1(s, NSIGNIF)) {
				/* lost a digit */
				e++;
				xfmtexp(s+NSIGNIF, e, 0);
			}
			continue;
		}
		break;
	}

	/*
	 * play with the decimal to try to simplify.
	 */

	/*
	 * bump last few digits up to 9 if we can
	 */
	for(i=NSIGNIF-1; i&gt;=NSIGNIF-3; i--) {
		c = s[i];
		if(c != &#39;9&#39;) {
			s[i] = &#39;9&#39;;
			g = strtod(s, nil);
			if(g != f) {
				s[i] = c;
				break;
			}
		}
	}

	/*
	 * add 1 in hopes of turning 9s to 0s
	 */
	if(s[NSIGNIF-1] == &#39;9&#39;) {
		strcpy(tmp, s);
		ee = e;
		if(xadd1(tmp, NSIGNIF)) {
			ee--;
			xfmtexp(tmp+NSIGNIF, ee, 0);
		}
		g = strtod(tmp, nil);
		if(g == f) {
			strcpy(s, tmp);
			e = ee;
		}
	}

	/*
	 * bump last few digits down to 0 as we can.
	 */
	for(i=NSIGNIF-1; i&gt;=NSIGNIF-3; i--) {
		c = s[i];
		if(c != &#39;0&#39;) {
			s[i] = &#39;0&#39;;
			g = strtod(s, nil);
			if(g != f) {
				s[i] = c;
				break;
			}
		}
	}

	/*
	 * remove trailing zeros.
	 */
	ndigit = NSIGNIF;
	while(ndigit &gt; 1 &amp;&amp; s[ndigit-1] == &#39;0&#39;){
		e++;
		--ndigit;
	}
	s[ndigit] = 0;
	*exp = e;
	*ns = ndigit;
	errno = oerrno;
}

#ifdef PLAN9PORT
static char *special[] = { &#34;NaN&#34;, &#34;NaN&#34;, &#34;+Inf&#34;, &#34;+Inf&#34;, &#34;-Inf&#34;, &#34;-Inf&#34; };
#else
static char *special[] = { &#34;nan&#34;, &#34;NAN&#34;, &#34;inf&#34;, &#34;INF&#34;, &#34;-inf&#34;, &#34;-INF&#34; };
#endif

int
__efgfmt(Fmt *fmt)
{
	char buf[NSIGNIF+10], *dot, *digits, *p, *s, suf[10], *t;
	double f;
	int c, chr, dotwid, e, exp, fl, ndigits, neg, newndigits;
	int pad, point, prec, realchr, sign, sufwid, ucase, wid, z1, z2;
	Rune r, *rs, *rt;

	if(fmt-&gt;flags&amp;FmtLong)
		f = va_arg(fmt-&gt;args, long double);
	else
		f = va_arg(fmt-&gt;args, double);

	/*
	 * extract formatting flags
	 */
	fl = fmt-&gt;flags;
	fmt-&gt;flags = 0;
	prec = FDEFLT;
	if(fl &amp; FmtPrec)
		prec = fmt-&gt;prec;
	chr = fmt-&gt;r;
	ucase = 0;
	switch(chr) {
	case &#39;A&#39;:
	case &#39;E&#39;:
	case &#39;F&#39;:
	case &#39;G&#39;:
		chr += &#39;a&#39;-&#39;A&#39;;
		ucase = 1;
		break;
	}

	/*
	 * pick off special numbers.
	 */
	if(__isNaN(f)) {
		s = special[0+ucase];
	special:
		fmt-&gt;flags = fl &amp; (FmtWidth|FmtLeft);
		return __fmtcpy(fmt, s, strlen(s), strlen(s));
	}
	if(__isInf(f, 1)) {
		s = special[2+ucase];
		goto special;
	}
	if(__isInf(f, -1)) {
		s = special[4+ucase];
		goto special;
	}

	/*
	 * get exact representation.
	 */
	digits = buf;
	xdtoa(f, digits, &amp;exp, &amp;neg, &amp;ndigits);

	/*
	 * get locale&#39;s decimal point.
	 */
	dot = fmt-&gt;decimal;
	if(dot == nil)
		dot = &#34;.&#34;;
	dotwid = utflen(dot);

	/*
	 * now the formatting fun begins.
	 * compute parameters for actual fmt:
	 *
	 *	pad: number of spaces to insert before/after field.
	 *	z1: number of zeros to insert before digits
	 *	z2: number of zeros to insert after digits
	 *	point: number of digits to print before decimal point
	 *	ndigits: number of digits to use from digits[]
	 *	suf: trailing suffix, like &#34;e-5&#34;
	 */
	realchr = chr;
	switch(chr){
	case &#39;g&#39;:
		/*
		 * convert to at most prec significant digits. (prec=0 means 1)
		 */
		if(prec == 0)
			prec = 1;
		if(ndigits &gt; prec) {
			if(digits[prec] &gt;= &#39;5&#39; &amp;&amp; xadd1(digits, prec))
				exp++;
			exp += ndigits-prec;
			ndigits = prec;
		}

		/*
		 * extra rules for %g (implemented below):
		 *	trailing zeros removed after decimal unless FmtSharp.
		 *	decimal point only if digit follows.
		 */

		/* fall through to %e */
	default:
	case &#39;e&#39;:
		/*
		 * one significant digit before decimal, no leading zeros.
		 */
		point = 1;
		z1 = 0;

		/*
		 * decimal point is after ndigits digits right now.
		 * slide to be after first.
		 */
		e  = exp + (ndigits-1);

		/*
		 * if this is %g, check exponent and convert prec
		 */
		if(realchr == &#39;g&#39;) {
			if(-4 &lt;= e &amp;&amp; e &lt; prec)
				goto casef;
			prec--;	/* one digit before decimal; rest after */
		}

		/*
		 * compute trailing zero padding or truncate digits.
		 */
		if(1+prec &gt;= ndigits)
			z2 = 1+prec - ndigits;
		else {
			/*
			 * truncate digits
			 */
			assert(realchr != &#39;g&#39;);
			newndigits = 1+prec;
			if(digits[newndigits] &gt;= &#39;5&#39; &amp;&amp; xadd1(digits, newndigits)) {
				/*
				 * had 999e4, now have 100e5
				 */
				e++;
			}
			ndigits = newndigits;
			z2 = 0;
		}
		xfmtexp(suf, e, ucase);
		sufwid = strlen(suf);
		break;

	casef:
	case &#39;f&#39;:
		/*
		 * determine where digits go with respect to decimal point
		 */
		if(ndigits+exp &gt; 0) {
			point = ndigits+exp;
			z1 = 0;
		} else {
			point = 1;
			z1 = 1 + -(ndigits+exp);
		}

		/*
		 * %g specifies prec = number of significant digits
		 * convert to number of digits after decimal point
		 */
		if(realchr == &#39;g&#39;)
			prec += z1 - point;

		/*
		 * compute trailing zero padding or truncate digits.
		 */
		if(point+prec &gt;= z1+ndigits)
			z2 = point+prec - (z1+ndigits);
		else {
			/*
			 * truncate digits
			 */
			assert(realchr != &#39;g&#39;);
			newndigits = point+prec - z1;
			if(newndigits &lt; 0) {
				z1 += newndigits;
				newndigits = 0;
			} else if(newndigits == 0) {
				/* perhaps round up */
				if(digits[0] &gt;= &#39;5&#39;){
					digits[0] = &#39;1&#39;;
					newndigits = 1;
					goto newdigit;
				}
			} else if(digits[newndigits] &gt;= &#39;5&#39; &amp;&amp; xadd1(digits, newndigits)) {
				/*
				 * digits was 999, is now 100; make it 1000
				 */
				digits[newndigits++] = &#39;0&#39;;
			newdigit:
				/*
				 * account for new digit
				 */
				if(z1)	/* 0.099 =&gt; 0.100 or 0.99 =&gt; 1.00*/
					z1--;
				else	/* 9.99 =&gt; 10.00 */
					point++;
			}
			z2 = 0;
			ndigits = newndigits;
		}
		sufwid = 0;
		break;
	}

	/*
	 * if %g is given without FmtSharp, remove trailing zeros.
	 * must do after truncation, so that e.g. print %.3g 1.001
	 * produces 1, not 1.00.  sorry, but them&#39;s the rules.
	 */
	if(realchr == &#39;g&#39; &amp;&amp; !(fl &amp; FmtSharp)) {
		if(z1+ndigits+z2 &gt;= point) {
			if(z1+ndigits &lt; point)
				z2 = point - (z1+ndigits);
			else{
				z2 = 0;
				while(z1+ndigits &gt; point &amp;&amp; digits[ndigits-1] == &#39;0&#39;)
					ndigits--;
			}
		}
	}

	/*
	 * compute width of all digits and decimal point and suffix if any
	 */
	wid = z1+ndigits+z2;
	if(wid &gt; point)
		wid += dotwid;
	else if(wid == point){
		if(fl &amp; FmtSharp)
			wid += dotwid;
		else
			point++;	/* do not print any decimal point */
	}
	wid += sufwid;

	/*
	 * determine sign
	 */
	sign = 0;
	if(neg)
		sign = &#39;-&#39;;
	else if(fl &amp; FmtSign)
		sign = &#39;+&#39;;
	else if(fl &amp; FmtSpace)
		sign = &#39; &#39;;
	if(sign)
		wid++;

	/*
	 * compute padding
	 */
	pad = 0;
	if((fl &amp; FmtWidth) &amp;&amp; fmt-&gt;width &gt; wid)
		pad = fmt-&gt;width - wid;
	if(pad &amp;&amp; !(fl &amp; FmtLeft) &amp;&amp; (fl &amp; FmtZero)){
		z1 += pad;
		point += pad;
		pad = 0;
	}

	/*
	 * format the actual field.  too bad about doing this twice.
	 */
	if(fmt-&gt;runes){
		if(pad &amp;&amp; !(fl &amp; FmtLeft) &amp;&amp; __rfmtpad(fmt, pad) &lt; 0)
			return -1;
		rt = (Rune*)fmt-&gt;to;
		rs = (Rune*)fmt-&gt;stop;
		if(sign)
			FMTRCHAR(fmt, rt, rs, sign);
		while(z1&gt;0 || ndigits&gt;0 || z2&gt;0) {
			if(z1 &gt; 0){
				z1--;
				c = &#39;0&#39;;
			}else if(ndigits &gt; 0){
				ndigits--;
				c = *digits++;
			}else{
				z2--;
				c = &#39;0&#39;;
			}
			FMTRCHAR(fmt, rt, rs, c);
			if(--point == 0) {
				for(p = dot; *p; ){
					p += chartorune(&amp;r, p);
					FMTRCHAR(fmt, rt, rs, r);
				}
			}
		}
		fmt-&gt;nfmt += rt - (Rune*)fmt-&gt;to;
		fmt-&gt;to = rt;
		if(sufwid &amp;&amp; __fmtcpy(fmt, suf, sufwid, sufwid) &lt; 0)
			return -1;
		if(pad &amp;&amp; (fl &amp; FmtLeft) &amp;&amp; __rfmtpad(fmt, pad) &lt; 0)
			return -1;
	}else{
		if(pad &amp;&amp; !(fl &amp; FmtLeft) &amp;&amp; __fmtpad(fmt, pad) &lt; 0)
			return -1;
		t = (char*)fmt-&gt;to;
		s = (char*)fmt-&gt;stop;
		if(sign)
			FMTCHAR(fmt, t, s, sign);
		while(z1&gt;0 || ndigits&gt;0 || z2&gt;0) {
			if(z1 &gt; 0){
				z1--;
				c = &#39;0&#39;;
			}else if(ndigits &gt; 0){
				ndigits--;
				c = *digits++;
			}else{
				z2--;
				c = &#39;0&#39;;
			}
			FMTCHAR(fmt, t, s, c);
			if(--point == 0)
				for(p=dot; *p; p++)
					FMTCHAR(fmt, t, s, *p);
		}
		fmt-&gt;nfmt += t - (char*)fmt-&gt;to;
		fmt-&gt;to = t;
		if(sufwid &amp;&amp; __fmtcpy(fmt, suf, sufwid, sufwid) &lt; 0)
			return -1;
		if(pad &amp;&amp; (fl &amp; FmtLeft) &amp;&amp; __fmtpad(fmt, pad) &lt; 0)
			return -1;
	}
	return 0;
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

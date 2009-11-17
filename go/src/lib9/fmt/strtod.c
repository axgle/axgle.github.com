<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/fmt/strtod.c</title>

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
  <h1 id="generatedHeader">Text file src/lib9/fmt/strtod.c</h1>

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

#include &lt;u.h&gt;
#include &lt;errno.h&gt;
#include &lt;libc.h&gt;
#include &#34;fmtdef.h&#34;

static ulong
umuldiv(ulong a, ulong b, ulong c)
{
	double d;

	d = ((double)a * (double)b) / (double)c;
	if(d &gt;= 4294967295.)
		d = 4294967295.;
	return (ulong)d;
}

/*
 * This routine will convert to arbitrary precision
 * floating point entirely in multi-precision fixed.
 * The answer is the closest floating point number to
 * the given decimal number. Exactly half way are
 * rounded ala ieee rules.
 * Method is to scale input decimal between .500 and .999...
 * with external power of 2, then binary search for the
 * closest mantissa to this decimal number.
 * Nmant is is the required precision. (53 for ieee dp)
 * Nbits is the max number of bits/word. (must be &lt;= 28)
 * Prec is calculated - the number of words of fixed mantissa.
 */
enum
{
	Nbits	= 28,				/* bits safely represented in a ulong */
	Nmant	= 53,				/* bits of precision required */
	Prec	= (Nmant+Nbits+1)/Nbits,	/* words of Nbits each to represent mantissa */
	Sigbit	= 1&lt;&lt;(Prec*Nbits-Nmant),	/* first significant bit of Prec-th word */
	Ndig	= 1500,
	One	= (ulong)(1&lt;&lt;Nbits),
	Half	= (ulong)(One&gt;&gt;1),
	Maxe	= 310,

	Fsign	= 1&lt;&lt;0,		/* found - */
	Fesign	= 1&lt;&lt;1,		/* found e- */
	Fdpoint	= 1&lt;&lt;2,		/* found . */

	S0	= 0,		/* _		_S0	+S1	#S2	.S3 */
	S1,			/* _+		#S2	.S3 */
	S2,			/* _+#		#S2	.S4	eS5 */
	S3,			/* _+.		#S4 */
	S4,			/* _+#.#	#S4	eS5 */
	S5,			/* _+#.#e	+S6	#S7 */
	S6,			/* _+#.#e+	#S7 */
	S7			/* _+#.#e+#	#S7 */
};

static	int	xcmp(char*, char*);
static	int	fpcmp(char*, ulong*);
static	void	frnorm(ulong*);
static	void	divascii(char*, int*, int*, int*);
static	void	mulascii(char*, int*, int*, int*);

typedef	struct	Tab	Tab;
struct	Tab
{
	int	bp;
	int	siz;
	char*	cmp;
};

double
fmtstrtod(const char *as, char **aas)
{
	int na, ex, dp, bp, c, i, flag, state;
	ulong low[Prec], hig[Prec], mid[Prec];
	double d;
	char *s, a[Ndig];

	flag = 0;	/* Fsign, Fesign, Fdpoint */
	na = 0;		/* number of digits of a[] */
	dp = 0;		/* na of decimal point */
	ex = 0;		/* exonent */

	state = S0;
	for(s=(char*)as;; s++) {
		c = *s;
		if(c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;) {
			switch(state) {
			case S0:
			case S1:
			case S2:
				state = S2;
				break;
			case S3:
			case S4:
				state = S4;
				break;

			case S5:
			case S6:
			case S7:
				state = S7;
				ex = ex*10 + (c-&#39;0&#39;);
				continue;
			}
			if(na == 0 &amp;&amp; c == &#39;0&#39;) {
				dp--;
				continue;
			}
			if(na &lt; Ndig-50)
				a[na++] = c;
			continue;
		}
		switch(c) {
		case &#39;\t&#39;:
		case &#39;\n&#39;:
		case &#39;\v&#39;:
		case &#39;\f&#39;:
		case &#39;\r&#39;:
		case &#39; &#39;:
			if(state == S0)
				continue;
			break;
		case &#39;-&#39;:
			if(state == S0)
				flag |= Fsign;
			else
				flag |= Fesign;
		case &#39;+&#39;:
			if(state == S0)
				state = S1;
			else
			if(state == S5)
				state = S6;
			else
				break;	/* syntax */
			continue;
		case &#39;.&#39;:
			flag |= Fdpoint;
			dp = na;
			if(state == S0 || state == S1) {
				state = S3;
				continue;
			}
			if(state == S2) {
				state = S4;
				continue;
			}
			break;
		case &#39;e&#39;:
		case &#39;E&#39;:
			if(state == S2 || state == S4) {
				state = S5;
				continue;
			}
			break;
		}
		break;
	}

	/*
	 * clean up return char-pointer
	 */
	switch(state) {
	case S0:
		if(xcmp(s, &#34;nan&#34;) == 0) {
			if(aas != nil)
				*aas = s+3;
			goto retnan;
		}
	case S1:
		if(xcmp(s, &#34;infinity&#34;) == 0) {
			if(aas != nil)
				*aas = s+8;
			goto retinf;
		}
		if(xcmp(s, &#34;inf&#34;) == 0) {
			if(aas != nil)
				*aas = s+3;
			goto retinf;
		}
	case S3:
		if(aas != nil)
			*aas = (char*)as;
		goto ret0;	/* no digits found */
	case S6:
		s--;		/* back over +- */
	case S5:
		s--;		/* back over e */
		break;
	}
	if(aas != nil)
		*aas = s;

	if(flag &amp; Fdpoint)
	while(na &gt; 0 &amp;&amp; a[na-1] == &#39;0&#39;)
		na--;
	if(na == 0)
		goto ret0;	/* zero */
	a[na] = 0;
	if(!(flag &amp; Fdpoint))
		dp = na;
	if(flag &amp; Fesign)
		ex = -ex;
	dp += ex;
	if(dp &lt; -Maxe){
		errno = ERANGE;
		goto ret0;	/* underflow by exp */
	} else
	if(dp &gt; +Maxe)
		goto retinf;	/* overflow by exp */

	/*
	 * normalize the decimal ascii number
	 * to range .[5-9][0-9]* e0
	 */
	bp = 0;		/* binary exponent */
	while(dp &gt; 0)
		divascii(a, &amp;na, &amp;dp, &amp;bp);
	while(dp &lt; 0 || a[0] &lt; &#39;5&#39;)
		mulascii(a, &amp;na, &amp;dp, &amp;bp);

	/* close approx by naive conversion */
	mid[0] = 0;
	mid[1] = 1;
	for(i=0; (c=a[i]) != &#39;\0&#39;; i++) {
		mid[0] = mid[0]*10 + (c-&#39;0&#39;);
		mid[1] = mid[1]*10;
		if(i &gt;= 8)
			break;
	}
	low[0] = umuldiv(mid[0], One, mid[1]);
	hig[0] = umuldiv(mid[0]+1, One, mid[1]);
	for(i=1; i&lt;Prec; i++) {
		low[i] = 0;
		hig[i] = One-1;
	}

	/* binary search for closest mantissa */
	for(;;) {
		/* mid = (hig + low) / 2 */
		c = 0;
		for(i=0; i&lt;Prec; i++) {
			mid[i] = hig[i] + low[i];
			if(c)
				mid[i] += One;
			c = mid[i] &amp; 1;
			mid[i] &gt;&gt;= 1;
		}
		frnorm(mid);

		/* compare */
		c = fpcmp(a, mid);
		if(c &gt; 0) {
			c = 1;
			for(i=0; i&lt;Prec; i++)
				if(low[i] != mid[i]) {
					c = 0;
					low[i] = mid[i];
				}
			if(c)
				break;	/* between mid and hig */
			continue;
		}
		if(c &lt; 0) {
			for(i=0; i&lt;Prec; i++)
				hig[i] = mid[i];
			continue;
		}

		/* only hard part is if even/odd roundings wants to go up */
		c = mid[Prec-1] &amp; (Sigbit-1);
		if(c == Sigbit/2 &amp;&amp; (mid[Prec-1]&amp;Sigbit) == 0)
			mid[Prec-1] -= c;
		break;	/* exactly mid */
	}

	/* normal rounding applies */
	c = mid[Prec-1] &amp; (Sigbit-1);
	mid[Prec-1] -= c;
	if(c &gt;= Sigbit/2) {
		mid[Prec-1] += Sigbit;
		frnorm(mid);
	}
	goto out;

ret0:
	return 0;

retnan:
	return __NaN();

retinf:
	/*
	 * Unix strtod requires these.  Plan 9 would return Inf(0) or Inf(-1). */
	errno = ERANGE;
	if(flag &amp; Fsign)
		return -HUGE_VAL;
	return HUGE_VAL;

out:
	d = 0;
	for(i=0; i&lt;Prec; i++)
		d = d*One + mid[i];
	if(flag &amp; Fsign)
		d = -d;
	d = ldexp(d, bp - Prec*Nbits);
	if(d == 0){	/* underflow */
		errno = ERANGE;
	}
	return d;
}

static void
frnorm(ulong *f)
{
	int i, c;

	c = 0;
	for(i=Prec-1; i&gt;0; i--) {
		f[i] += c;
		c = f[i] &gt;&gt; Nbits;
		f[i] &amp;= One-1;
	}
	f[0] += c;
}

static int
fpcmp(char *a, ulong* f)
{
	ulong tf[Prec];
	int i, d, c;

	for(i=0; i&lt;Prec; i++)
		tf[i] = f[i];

	for(;;) {
		/* tf *= 10 */
		for(i=0; i&lt;Prec; i++)
			tf[i] = tf[i]*10;
		frnorm(tf);
		d = (tf[0] &gt;&gt; Nbits) + &#39;0&#39;;
		tf[0] &amp;= One-1;

		/* compare next digit */
		c = *a;
		if(c == 0) {
			if(&#39;0&#39; &lt; d)
				return -1;
			if(tf[0] != 0)
				goto cont;
			for(i=1; i&lt;Prec; i++)
				if(tf[i] != 0)
					goto cont;
			return 0;
		}
		if(c &gt; d)
			return +1;
		if(c &lt; d)
			return -1;
		a++;
	cont:;
	}
}

static void
divby(char *a, int *na, int b)
{
	int n, c;
	char *p;

	p = a;
	n = 0;
	while(n&gt;&gt;b == 0) {
		c = *a++;
		if(c == 0) {
			while(n) {
				c = n*10;
				if(c&gt;&gt;b)
					break;
				n = c;
			}
			goto xx;
		}
		n = n*10 + c-&#39;0&#39;;
		(*na)--;
	}
	for(;;) {
		c = n&gt;&gt;b;
		n -= c&lt;&lt;b;
		*p++ = c + &#39;0&#39;;
		c = *a++;
		if(c == 0)
			break;
		n = n*10 + c-&#39;0&#39;;
	}
	(*na)++;
xx:
	while(n) {
		n = n*10;
		c = n&gt;&gt;b;
		n -= c&lt;&lt;b;
		*p++ = c + &#39;0&#39;;
		(*na)++;
	}
	*p = 0;
}

static	Tab	tab1[] =
{
	 1,  0, &#34;&#34;,
	 3,  1, &#34;7&#34;,
	 6,  2, &#34;63&#34;,
	 9,  3, &#34;511&#34;,
	13,  4, &#34;8191&#34;,
	16,  5, &#34;65535&#34;,
	19,  6, &#34;524287&#34;,
	23,  7, &#34;8388607&#34;,
	26,  8, &#34;67108863&#34;,
	27,  9, &#34;134217727&#34;,
};

static void
divascii(char *a, int *na, int *dp, int *bp)
{
	int b, d;
	Tab *t;

	d = *dp;
	if(d &gt;= (int)(nelem(tab1)))
		d = (int)(nelem(tab1))-1;
	t = tab1 + d;
	b = t-&gt;bp;
	if(memcmp(a, t-&gt;cmp, t-&gt;siz) &gt; 0)
		d--;
	*dp -= d;
	*bp += b;
	divby(a, na, b);
}

static void
mulby(char *a, char *p, char *q, int b)
{
	int n, c;

	n = 0;
	*p = 0;
	for(;;) {
		q--;
		if(q &lt; a)
			break;
		c = *q - &#39;0&#39;;
		c = (c&lt;&lt;b) + n;
		n = c/10;
		c -= n*10;
		p--;
		*p = c + &#39;0&#39;;
	}
	while(n) {
		c = n;
		n = c/10;
		c -= n*10;
		p--;
		*p = c + &#39;0&#39;;
	}
}

static	Tab	tab2[] =
{
	 1,  1, &#34;&#34;,				/* dp = 0-0 */
	 3,  3, &#34;125&#34;,
	 6,  5, &#34;15625&#34;,
	 9,  7, &#34;1953125&#34;,
	13, 10, &#34;1220703125&#34;,
	16, 12, &#34;152587890625&#34;,
	19, 14, &#34;19073486328125&#34;,
	23, 17, &#34;11920928955078125&#34;,
	26, 19, &#34;1490116119384765625&#34;,
	27, 19, &#34;7450580596923828125&#34;,		/* dp 8-9 */
};

static void
mulascii(char *a, int *na, int *dp, int *bp)
{
	char *p;
	int d, b;
	Tab *t;

	d = -*dp;
	if(d &gt;= (int)(nelem(tab2)))
		d = (int)(nelem(tab2))-1;
	t = tab2 + d;
	b = t-&gt;bp;
	if(memcmp(a, t-&gt;cmp, t-&gt;siz) &lt; 0)
		d--;
	p = a + *na;
	*bp -= b;
	*dp += d;
	*na += d;
	mulby(a, p+d, p, b);
}

static int
xcmp(char *a, char *b)
{
	int c1, c2;

	while((c1 = *b++) != &#39;\0&#39;) {
		c2 = *a++;
		if(isupper(c2))
			c2 = tolower(c2);
		if(c1 != c2)
			return 1;
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

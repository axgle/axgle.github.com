<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/utf/mkrunetype.c</title>

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
  <h1 id="generatedHeader">Text file src/lib9/utf/mkrunetype.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * make is(upper|lower|title|space|alpha)rune and
 * to(upper|lower|title)rune from a UnicodeData.txt file.
 * these can be found at unicode.org
 *
 * with -c, runs a check of the existing runetype functions vs.
 * those extracted from UnicodeData.
 *
 * with -p, generates tables for pairs of chars, as well as for ranges
 * and singletons.
 *
 * UnicodeData defines 4 fields of interest:
 * 1) a category
 * 2) an upper case mapping
 * 3) a lower case mapping
 * 4) a title case mapping
 *
 * toupper, tolower, and totitle are defined directly from the mapping.
 *
 * isalpharune(c) is true iff c is a &#34;letter&#34; category
 * isupperrune(c) is true iff c is the target of toupperrune,
 *	or is in the uppercase letter category
 * similarly for islowerrune and istitlerune.
 * isspacerune is true for space category chars, &#34;C&#34; locale white space chars,
 *	and two additions:
 *	0085	&#34;next line&#34; control char
 *	feff]	&#34;zero-width non-break space&#34;
 * isdigitrune is true iff c is a numeric-digit category.
 */

#include &lt;stdio.h&gt;
#include &lt;stdlib.h&gt;
#include &lt;stdarg.h&gt;
#include &lt;string.h&gt;
#include &lt;libgen.h&gt;
#include &#34;utf.h&#34;
#include &#34;utfdef.h&#34;

enum {
	/*
	 * fields in the unicode data file
	 */
	FIELD_CODE,
	FIELD_NAME,
	FIELD_CATEGORY,
	FIELD_COMBINING,
	FIELD_BIDIR,
	FIELD_DECOMP,
	FIELD_DECIMAL_DIG,
	FIELD_DIG,
	FIELD_NUMERIC_VAL,
	FIELD_MIRRORED,
	FIELD_UNICODE_1_NAME,
	FIELD_COMMENT,
	FIELD_UPPER,
	FIELD_LOWER,
	FIELD_TITLE,
	NFIELDS,

	MAX_LINE	= 1024,

	TO_OFFSET	= 1 &lt;&lt; 20,

	NRUNES		= 1 &lt;&lt; 21,
};

#define TO_DELTA(xmapped,x)	(TO_OFFSET + (xmapped) - (x))

static char	myisspace[NRUNES];
static char	myisalpha[NRUNES];
static char	myisdigit[NRUNES];
static char	myisupper[NRUNES];
static char	myislower[NRUNES];
static char	myistitle[NRUNES];

static int	mytoupper[NRUNES];
static int	mytolower[NRUNES];
static int	mytotitle[NRUNES];

static void	check(void);
static void	mktables(char *src, int usepairs);
static void	fatal(const char *fmt, ...);
static int	mygetfields(char **fields, int nfields, char *str, const char *delim);
static int	getunicodeline(FILE *in, char **fields, char *buf);
static int	getcode(char *s);

static void
usage(void)
{
	fprintf(stderr, &#34;usage: mktables [-cp] &lt;UnicodeData.txt&gt;\n&#34;);
	exit(1);
}

int
main(int argc, char *argv[]){
	FILE *in;
	char buf[MAX_LINE], buf2[MAX_LINE];
	char *fields[NFIELDS + 1], *fields2[NFIELDS + 1];
	char *p;
	int i, code, last, docheck, usepairs;

	docheck = 0;
	usepairs = 0;
	ARGBEGIN{
	case &#39;c&#39;:
		docheck = 1;
		break;
	case &#39;p&#39;:
		usepairs = 1;
		break;
	default:
		usage();
	}ARGEND

	if(argc != 1){
		usage();
	}

	in = fopen(argv[0], &#34;r&#34;);
	if(in == NULL){
		fatal(&#34;can&#39;t open %s&#34;, argv[0]);
	}

	for(i = 0; i &lt; NRUNES; i++){
		mytoupper[i] = i;
		mytolower[i] = i;
		mytotitle[i] = i;
	}

	/*
	 * make sure isspace has all of the &#34;C&#34; locale whitespace chars
	 */
	myisspace[&#39;\t&#39;] = 1;
	myisspace[&#39;\n&#39;] = 1;
	myisspace[&#39;\r&#39;] = 1;
	myisspace[&#39;\f&#39;] = 1;
	myisspace[&#39;\v&#39;] = 1;

	/*
	 * a couple of other exceptions
	 */
	myisspace[0x85] = 1;	/* control char, &#34;next line&#34; */
	myisspace[0xfeff] = 1;	/* zero-width non-break space */

	last = -1;
	while(getunicodeline(in, fields, buf)){
		code = getcode(fields[FIELD_CODE]);
                if (code &gt;= NRUNES)
                  fatal(&#34;code-point value too big: %x&#34;, code);
		if(code &lt;= last)
			fatal(&#34;bad code sequence: %x then %x&#34;, last, code);
		last = code;

		/*
		 * check for ranges
		 */
		p = fields[FIELD_CATEGORY];
		if(strstr(fields[FIELD_NAME], &#34;, First&gt;&#34;) != NULL){
			if(!getunicodeline(in, fields2, buf2))
				fatal(&#34;range start at eof&#34;);
			if (strstr(fields2[FIELD_NAME], &#34;, Last&gt;&#34;) == NULL)
				fatal(&#34;range start not followed by range end&#34;);
			last = getcode(fields2[FIELD_CODE]);
			if(last &lt;= code)
				fatal(&#34;range out of sequence: %x then %x&#34;, code, last);
			if(strcmp(p, fields2[FIELD_CATEGORY]) != 0)
				fatal(&#34;range with mismatched category&#34;);
		}

		/*
		 * set properties and conversions
		 */
		for (; code &lt;= last; code++){
			if(p[0] == &#39;L&#39;)
				myisalpha[code] = 1;
			if(p[0] == &#39;Z&#39;)
				myisspace[code] = 1;

			if(strcmp(p, &#34;Lu&#34;) == 0)
				myisupper[code] = 1;
			if(strcmp(p, &#34;Ll&#34;) == 0)
				myislower[code] = 1;

			if(strcmp(p, &#34;Lt&#34;) == 0)
				myistitle[code] = 1;

			if(strcmp(p, &#34;Nd&#34;) == 0)
				myisdigit[code] = 1;

			/*
			 * when finding conversions, also need to mark
			 * upper/lower case, since some chars, like
			 * &#34;III&#34; (0x2162), aren&#39;t defined as letters but have a
			 * lower case mapping (&#34;iii&#34; (0x2172)).
			 */
			if(fields[FIELD_UPPER][0] != &#39;\0&#39;){
				mytoupper[code] = getcode(fields[FIELD_UPPER]);
			}
			if(fields[FIELD_LOWER][0] != &#39;\0&#39;){
				mytolower[code] = getcode(fields[FIELD_LOWER]);
			}
			if(fields[FIELD_TITLE][0] != &#39;\0&#39;){
				mytotitle[code] = getcode(fields[FIELD_TITLE]);
			}
		}
	}

	fclose(in);

	/*
	 * check for codes with no totitle mapping but a toupper mapping.
	 * these appear in UnicodeData-2.0.14.txt, but are almost certainly
	 * erroneous.
	 */
	for(i = 0; i &lt; NRUNES; i++){
		if(mytotitle[i] == i
		&amp;&amp; mytoupper[i] != i
		&amp;&amp; !myistitle[i])
			fprintf(stderr, &#34;warning: code=%.4x not istitle, totitle is same, toupper=%.4x\n&#34;, i, mytoupper[i]);
	}

	/*
	 * make sure isupper[c] is true if for some x toupper[x]  == c
	 * ditto for islower and istitle
	 */
	for(i = 0; i &lt; NRUNES; i++) {
		if(mytoupper[i] != i)
			myisupper[mytoupper[i]] = 1;
		if(mytolower[i] != i)
			myislower[mytolower[i]] = 1;
		if(mytotitle[i] != i)
			myistitle[mytotitle[i]] = 1;
	}

	if(docheck){
		check();
	}else{
		mktables(argv[0], usepairs);
	}
	return 0;
}

/*
 * generate a properties array for ranges, clearing those cases covered.
 * if force, generate one-entry ranges for singletons.
 */
static int
mkisrange(const char* label, char* prop, int force)
{
	int start, stop, some;

	/*
	 * first, the ranges
	 */
	some = 0;
	for(start = 0; start &lt; NRUNES; ) {
		if(!prop[start]){
			start++;
			continue;
		}

		for(stop = start + 1; stop &lt; NRUNES; stop++){
			if(!prop[stop]){
				break;
			}
			prop[stop] = 0;
		}
		if(force || stop != start + 1){
			if(!some){
				printf(&#34;static Rune __is%sr[] = {\n&#34;, label);
				some = 1;
			}
			prop[start] = 0;
			printf(&#34;\t0x%.4x, 0x%.4x,\n&#34;, start, stop - 1);
		}

		start = stop;
	}
	if(some)
		printf(&#34;};\n\n&#34;);
	return some;
}

/*
 * generate a mapping array for pairs with a skip between,
 * clearing those entries covered.
 */
static int
mkispair(const char *label, char *prop)
{
	int start, stop, some;

	some = 0;
	for(start = 0; start + 2 &lt; NRUNES; ) {
		if(!prop[start]){
			start++;
			continue;
		}

		for(stop = start + 2; stop &lt; NRUNES; stop += 2){
			if(!prop[stop]){
				break;
			}
			prop[stop] = 0;
		}
		if(stop != start + 2){
			if(!some){
				printf(&#34;static Rune __is%sp[] = {\n&#34;, label);
				some = 1;
			}
			prop[start] = 0;
			printf(&#34;\t0x%.4x, 0x%.4x,\n&#34;, start, stop - 2);
		}

		start = stop;
	}
	if(some)
		printf(&#34;};\n\n&#34;);
	return some;
}

/*
 * generate a properties array for singletons, clearing those cases covered.
 */
static int
mkissingle(const char *label, char *prop)
{
	int start, some;

	some = 0;
	for(start = 0; start &lt; NRUNES; start++) {
		if(!prop[start]){
			continue;
		}

		if(!some){
			printf(&#34;static Rune __is%ss[] = {\n&#34;, label);
			some = 1;
		}
		prop[start] = 0;
		printf(&#34;\t0x%.4x,\n&#34;, start);
	}
	if(some)
		printf(&#34;};\n\n&#34;);
	return some;
}

/*
 * generate tables and a function for is&lt;label&gt;rune
 */
static void
mkis(const char* label, char* prop, int usepairs)
{
	int isr, isp, iss;

	isr = mkisrange(label, prop, 0);
	isp = 0;
	if(usepairs)
		isp = mkispair(label, prop);
	iss = mkissingle(label, prop);

	printf(
		&#34;int\n&#34;
		&#34;is%srune(Rune c)\n&#34;
		&#34;{\n&#34;
		&#34;	Rune *p;\n&#34;
		&#34;\n&#34;,
		label);

	if(isr)
		printf(
			&#34;	p = rbsearch(c, __is%sr, nelem(__is%sr)/2, 2);\n&#34;
			&#34;	if(p &amp;&amp; c &gt;= p[0] &amp;&amp; c &lt;= p[1])\n&#34;
			&#34;		return 1;\n&#34;,
			label, label);

	if(isp)
		printf(
			&#34;	p = rbsearch(c, __is%sp, nelem(__is%sp)/2, 2);\n&#34;
			&#34;	if(p &amp;&amp; c &gt;= p[0] &amp;&amp; c &lt;= p[1] &amp;&amp; !((c - p[0]) &amp; 1))\n&#34;
			&#34;		return 1;\n&#34;,
			label, label);

	if(iss)
		printf(
			&#34;	p = rbsearch(c, __is%ss, nelem(__is%ss), 1);\n&#34;
			&#34;	if(p &amp;&amp; c == p[0])\n&#34;
			&#34;		return 1;\n&#34;,
			label, label);


	printf(
		&#34;	return 0;\n&#34;
		&#34;}\n&#34;
		&#34;\n&#34;
	);
}

/*
 * generate a mapping array for ranges, clearing those entries covered.
 * if force, generate one-entry ranges for singletons.
 */
static int
mktorange(const char* label, int* map, int force)
{
	int start, stop, delta, some;

	some = 0;
	for(start = 0; start &lt; NRUNES; ) {
		if(map[start] == start){
			start++;
			continue;
		}

		delta = TO_DELTA(map[start], start);
		if(delta != (Rune)delta)
			fatal(&#34;bad map delta %d&#34;, delta);
		for(stop = start + 1; stop &lt; NRUNES; stop++){
			if(TO_DELTA(map[stop], stop) != delta){
				break;
			}
			map[stop] = stop;
		}
		if(stop != start + 1){
			if(!some){
				printf(&#34;static Rune __to%sr[] = {\n&#34;, label);
				some = 1;
			}
			map[start] = start;
			printf(&#34;\t0x%.4x, 0x%.4x, %d,\n&#34;, start, stop - 1, delta);
		}

		start = stop;
	}
	if(some)
		printf(&#34;};\n\n&#34;);
	return some;
}

/*
 * generate a mapping array for pairs with a skip between,
 * clearing those entries covered.
 */
static int
mktopair(const char* label, int* map)
{
	int start, stop, delta, some;

	some = 0;
	for(start = 0; start + 2 &lt; NRUNES; ) {
		if(map[start] == start){
			start++;
			continue;
		}

		delta = TO_DELTA(map[start], start);
		if(delta != (Rune)delta)
			fatal(&#34;bad map delta %d&#34;, delta);
		for(stop = start + 2; stop &lt; NRUNES; stop += 2){
			if(TO_DELTA(map[stop], stop) != delta){
				break;
			}
			map[stop] = stop;
		}
		if(stop != start + 2){
			if(!some){
				printf(&#34;static Rune __to%sp[] = {\n&#34;, label);
				some = 1;
			}
			map[start] = start;
			printf(&#34;\t0x%.4x, 0x%.4x, %d,\n&#34;, start, stop - 2, delta);
		}

		start = stop;
	}
	if(some)
		printf(&#34;};\n\n&#34;);
	return some;
}

/*
 * generate a mapping array for singletons, clearing those entries covered.
 */
static int
mktosingle(const char* label, int* map)
{
	int start, delta, some;

	some = 0;
	for(start = 0; start &lt; NRUNES; start++) {
		if(map[start] == start){
			continue;
		}

		delta = TO_DELTA(map[start], start);
		if(delta != (Rune)delta)
			fatal(&#34;bad map delta %d&#34;, delta);
		if(!some){
			printf(&#34;static Rune __to%ss[] = {\n&#34;, label);
			some = 1;
		}
		map[start] = start;
		printf(&#34;\t0x%.4x, %d,\n&#34;, start, delta);
	}
	if(some)
		printf(&#34;};\n\n&#34;);
	return some;
}

/*
 * generate tables and a function for to&lt;label&gt;rune
 */
static void
mkto(const char* label, int* map, int usepairs)
{
	int tor, top, tos;

	tor = mktorange(label, map, 0);
	top = 0;
	if(usepairs)
		top = mktopair(label, map);
	tos = mktosingle(label, map);

	printf(
		&#34;Rune\n&#34;
		&#34;to%srune(Rune c)\n&#34;
		&#34;{\n&#34;
		&#34;	Rune *p;\n&#34;
		&#34;\n&#34;,
		label);

	if(tor)
		printf(
			&#34;	p = rbsearch(c, __to%sr, nelem(__to%sr)/3, 3);\n&#34;
			&#34;	if(p &amp;&amp; c &gt;= p[0] &amp;&amp; c &lt;= p[1])\n&#34;
			&#34;		return c + p[2] - %d;\n&#34;,
			label, label, TO_OFFSET);

	if(top)
		printf(
			&#34;	p = rbsearch(c, __to%sp, nelem(__to%sp)/3, 3);\n&#34;
			&#34;	if(p &amp;&amp; c &gt;= p[0] &amp;&amp; c &lt;= p[1] &amp;&amp; !((c - p[0]) &amp; 1))\n&#34;
			&#34;		return c + p[2] - %d;\n&#34;,
			label, label, TO_OFFSET);

	if(tos)
		printf(
			&#34;	p = rbsearch(c, __to%ss, nelem(__to%ss)/2, 2);\n&#34;
			&#34;	if(p &amp;&amp; c == p[0])\n&#34;
			&#34;		return c + p[1] - %d;\n&#34;,
			label, label, TO_OFFSET);


	printf(
		&#34;	return c;\n&#34;
		&#34;}\n&#34;
		&#34;\n&#34;
	);
}

// Make only range tables and a function for is&lt;label&gt;rune.
static void
mkisronly(const char* label, char* prop) {
	mkisrange(label, prop, 1);
	printf(
		&#34;int\n&#34;
		&#34;is%srune(Rune c)\n&#34;
		&#34;{\n&#34;
		&#34;	Rune *p;\n&#34;
		&#34;\n&#34;
		&#34;	p = rbsearch(c, __is%sr, nelem(__is%sr)/2, 2);\n&#34;
		&#34;	if(p &amp;&amp; c &gt;= p[0] &amp;&amp; c &lt;= p[1])\n&#34;
		&#34;		return 1;\n&#34;
		&#34;	return 0;\n&#34;
		&#34;}\n&#34;
		&#34;\n&#34;,
	        label, label, label);
}

/*
 * generate the body of runetype.
 * assumes there is a function Rune* rbsearch(Rune c, Rune *t, int n, int ne);
 */
static void
mktables(char *src, int usepairs)
{
	printf(&#34;/* generated automatically by mkrunetype.c from %s */\n\n&#34;,
		basename(src));

	/*
	 * we special case the space and digit tables, since they are assumed
	 * to be small with several ranges.
	 */
	mkisronly(&#34;space&#34;, myisspace);
	mkisronly(&#34;digit&#34;, myisdigit);

	mkis(&#34;alpha&#34;, myisalpha, 0);
	mkis(&#34;upper&#34;, myisupper, usepairs);
	mkis(&#34;lower&#34;, myislower, usepairs);
	mkis(&#34;title&#34;, myistitle, usepairs);

	mkto(&#34;upper&#34;, mytoupper, usepairs);
	mkto(&#34;lower&#34;, mytolower, usepairs);
	mkto(&#34;title&#34;, mytotitle, usepairs);
}

/*
 * find differences between the newly generated tables and current runetypes.
 */
static void
check(void)
{
	int i;

	for(i = 0; i &lt; NRUNES; i++){
		if(isdigitrune(i) != myisdigit[i])
			fprintf(stderr, &#34;isdigit diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, isdigitrune(i), myisdigit[i]);

		if(isspacerune(i) != myisspace[i])
			fprintf(stderr, &#34;isspace diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, isspacerune(i), myisspace[i]);

		if(isupperrune(i) != myisupper[i])
			fprintf(stderr, &#34;isupper diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, isupperrune(i), myisupper[i]);

		if(islowerrune(i) != myislower[i])
			fprintf(stderr, &#34;islower diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, islowerrune(i), myislower[i]);

		if(isalpharune(i) != myisalpha[i])
			fprintf(stderr, &#34;isalpha diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, isalpharune(i), myisalpha[i]);

		if(toupperrune(i) != mytoupper[i])
			fprintf(stderr, &#34;toupper diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, toupperrune(i), mytoupper[i]);

		if(tolowerrune(i) != mytolower[i])
			fprintf(stderr, &#34;tolower diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, tolowerrune(i), mytolower[i]);

		if(istitlerune(i) != myistitle[i])
			fprintf(stderr, &#34;istitle diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, istitlerune(i), myistitle[i]);

		if(totitlerune(i) != mytotitle[i])
			fprintf(stderr, &#34;totitle diff at %x: runetype=%x, unicode=%x\n&#34;,
				i, totitlerune(i), mytotitle[i]);


	}
}

static int
mygetfields(char **fields, int nfields, char *str, const char *delim)
{
	int nf;

	fields[0] = str;
	nf = 1;
	if(nf &gt;= nfields)
		return nf;

	for(; *str; str++){
		if(strchr(delim, *str) != NULL){
			*str = &#39;\0&#39;;
			fields[nf++] = str + 1;
			if(nf &gt;= nfields)
				break;
		}
	}
	return nf;
}

static int
getunicodeline(FILE *in, char **fields, char *buf)
{
	char *p;

	if(fgets(buf, MAX_LINE, in) == NULL)
		return 0;

	p = strchr(buf, &#39;\n&#39;);
	if (p == NULL)
		fatal(&#34;line too long&#34;);
	*p = &#39;\0&#39;;

	if (mygetfields(fields, NFIELDS + 1, buf, &#34;;&#34;) != NFIELDS)
		fatal(&#34;bad number of fields&#34;);

	return 1;
}

static int
getcode(char *s)
{
	int i, code;

	code = 0;
        i = 0;
        /* Parse a hex number */
	while(s[i]) {
		code &lt;&lt;= 4;
		if(s[i] &gt;= &#39;0&#39; &amp;&amp; s[i] &lt;= &#39;9&#39;)
			code += s[i] - &#39;0&#39;;
		else if(s[i] &gt;= &#39;A&#39; &amp;&amp; s[i] &lt;= &#39;F&#39;)
			code += s[i] - &#39;A&#39; + 10;
		else
			fatal(&#34;bad code char &#39;%c&#39;&#34;, s[i]);
                i++;
	}
	return code;
}

static void
fatal(const char *fmt, ...)
{
	va_list arg;

	fprintf(stderr, &#34;%s: fatal error: &#34;, argv0);
	va_start(arg, fmt);
	vfprintf(stderr, fmt, arg);
	va_end(arg);
	fprintf(stderr, &#34;\n&#34;);

	exit(1);
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

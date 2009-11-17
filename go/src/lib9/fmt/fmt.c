<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/fmt/fmt.c</title>

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
	<li>Thu Nov 12 15:58:03 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/fmt/fmt.c</h1>

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
#include &lt;libc.h&gt;
#include &#34;fmtdef.h&#34;

enum
{
	Maxfmt = 64
};

typedef struct Convfmt Convfmt;
struct Convfmt
{
	int	c;
	volatile	Fmts	fmt;	/* for spin lock in fmtfmt; avoids race due to write order */
};

static struct
{
	/* lock by calling __fmtlock, __fmtunlock */
	int	nfmt;
	Convfmt	fmt[Maxfmt];
} fmtalloc;

static Convfmt knownfmt[] = {
	&#39; &#39;,	__flagfmt,
	&#39;#&#39;,	__flagfmt,
	&#39;%&#39;,	__percentfmt,
	&#39;\&#39;&#39;,	__flagfmt,
	&#39;+&#39;,	__flagfmt,
	&#39;,&#39;,	__flagfmt,
	&#39;-&#39;,	__flagfmt,
	&#39;C&#39;,	__runefmt,	/* Plan 9 addition */
	&#39;E&#39;,	__efgfmt,
#ifndef PLAN9PORT
	&#39;F&#39;,	__efgfmt,	/* ANSI only */
#endif
	&#39;G&#39;,	__efgfmt,
#ifndef PLAN9PORT
	&#39;L&#39;,	__flagfmt,	/* ANSI only */
#endif
	&#39;S&#39;,	__runesfmt,	/* Plan 9 addition */
	&#39;X&#39;,	__ifmt,
	&#39;b&#39;,	__ifmt,		/* Plan 9 addition */
	&#39;c&#39;,	__charfmt,
	&#39;d&#39;,	__ifmt,
	&#39;e&#39;,	__efgfmt,
	&#39;f&#39;,	__efgfmt,
	&#39;g&#39;,	__efgfmt,
	&#39;h&#39;,	__flagfmt,
#ifndef PLAN9PORT
	&#39;i&#39;,	__ifmt,		/* ANSI only */
#endif
	&#39;l&#39;,	__flagfmt,
	&#39;n&#39;,	__countfmt,
	&#39;o&#39;,	__ifmt,
	&#39;p&#39;,	__ifmt,
	&#39;r&#39;,	__errfmt,
	&#39;s&#39;,	__strfmt,
#ifdef PLAN9PORT
	&#39;u&#39;,	__flagfmt,
#else
	&#39;u&#39;,	__ifmt,
#endif
	&#39;x&#39;,	__ifmt,
	0,	nil,
};


int	(*fmtdoquote)(int);

/*
 * __fmtlock() must be set
 */
static int
__fmtinstall(int c, Fmts f)
{
	Convfmt *p, *ep;

	if(c&lt;=0 || c&gt;=65536)
		return -1;
	if(!f)
		f = __badfmt;

	ep = &amp;fmtalloc.fmt[fmtalloc.nfmt];
	for(p=fmtalloc.fmt; p&lt;ep; p++)
		if(p-&gt;c == c)
			break;

	if(p == &amp;fmtalloc.fmt[Maxfmt])
		return -1;

	p-&gt;fmt = f;
	if(p == ep){	/* installing a new format character */
		fmtalloc.nfmt++;
		p-&gt;c = c;
	}

	return 0;
}

int
fmtinstall(int c, int (*f)(Fmt*))
{
	int ret;

	__fmtlock();
	ret = __fmtinstall(c, f);
	__fmtunlock();
	return ret;
}

static Fmts
fmtfmt(int c)
{
	Convfmt *p, *ep;

	ep = &amp;fmtalloc.fmt[fmtalloc.nfmt];
	for(p=fmtalloc.fmt; p&lt;ep; p++)
		if(p-&gt;c == c){
			while(p-&gt;fmt == nil)	/* loop until value is updated */
				;
			return p-&gt;fmt;
		}

	/* is this a predefined format char? */
	__fmtlock();
	for(p=knownfmt; p-&gt;c; p++)
		if(p-&gt;c == c){
			__fmtinstall(p-&gt;c, p-&gt;fmt);
			__fmtunlock();
			return p-&gt;fmt;
		}
	__fmtunlock();

	return __badfmt;
}

void*
__fmtdispatch(Fmt *f, void *fmt, int isrunes)
{
	Rune rune, r;
	int i, n;

	f-&gt;flags = 0;
	f-&gt;width = f-&gt;prec = 0;

	for(;;){
		if(isrunes){
			r = *(Rune*)fmt;
			fmt = (Rune*)fmt + 1;
		}else{
			fmt = (char*)fmt + chartorune(&amp;rune, (char*)fmt);
			r = rune;
		}
		f-&gt;r = r;
		switch(r){
		case &#39;\0&#39;:
			return nil;
		case &#39;.&#39;:
			f-&gt;flags |= FmtWidth|FmtPrec;
			continue;
		case &#39;0&#39;:
			if(!(f-&gt;flags &amp; FmtWidth)){
				f-&gt;flags |= FmtZero;
				continue;
			}
			/* fall through */
		case &#39;1&#39;: case &#39;2&#39;: case &#39;3&#39;: case &#39;4&#39;:
		case &#39;5&#39;: case &#39;6&#39;: case &#39;7&#39;: case &#39;8&#39;: case &#39;9&#39;:
			i = 0;
			while(r &gt;= &#39;0&#39; &amp;&amp; r &lt;= &#39;9&#39;){
				i = i * 10 + r - &#39;0&#39;;
				if(isrunes){
					r = *(Rune*)fmt;
					fmt = (Rune*)fmt + 1;
				}else{
					r = *(char*)fmt;
					fmt = (char*)fmt + 1;
				}
			}
			if(isrunes)
				fmt = (Rune*)fmt - 1;
			else
				fmt = (char*)fmt - 1;
		numflag:
			if(f-&gt;flags &amp; FmtWidth){
				f-&gt;flags |= FmtPrec;
				f-&gt;prec = i;
			}else{
				f-&gt;flags |= FmtWidth;
				f-&gt;width = i;
			}
			continue;
		case &#39;*&#39;:
			i = va_arg(f-&gt;args, int);
			if(i &lt; 0){
				/*
				 * negative precision =&gt;
				 * ignore the precision.
				 */
				if(f-&gt;flags &amp; FmtPrec){
					f-&gt;flags &amp;= ~FmtPrec;
					f-&gt;prec = 0;
					continue;
				}
				i = -i;
				f-&gt;flags |= FmtLeft;
			}
			goto numflag;
		}
		n = (*fmtfmt(r))(f);
		if(n &lt; 0)
			return nil;
		if(n == 0)
			return fmt;
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

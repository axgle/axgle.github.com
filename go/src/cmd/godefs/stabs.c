<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/godefs/stabs.c</title>

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
	<li>Thu Nov 12 16:01:31 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/godefs/stabs.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse stabs debug info.

#include &#34;a.h&#34;

int stabsdebug = 1;

// Hash table for type lookup by number.
Type *hash[1024];

// Look up type by number pair.
// TODO(rsc): Iant points out that n1 and n2 are always small and dense,
// so an array of arrays would be a better representation.
Type*
typebynum(uint n1, uint n2)
{
	uint h;
	Type *t;

	h = (n1*53+n2) % nelem(hash);
	for(t=hash[h]; t; t=t-&gt;next)
		if(t-&gt;n1 == n1 &amp;&amp; t-&gt;n2 == n2)
			return t;
	t = emalloc(sizeof *t);
	t-&gt;next = hash[h];
	hash[h] = t;
	t-&gt;n1 = n1;
	t-&gt;n2 = n2;
	return t;
}

// Parse name and colon from *pp, leaving copy in *sp.
static int
parsename(char **pp, char **sp)
{
	char *p;
	char *s;

	p = *pp;
	while(*p != &#39;\0&#39; &amp;&amp; *p != &#39;:&#39;)
		p++;
	if(*p == &#39;\0&#39;) {
		fprint(2, &#34;parsename expected colon\n&#34;);
		return -1;
	}
	s = emalloc(p - *pp + 1);
	memmove(s, *pp, p - *pp);
	*sp = s;
	*pp = p+1;
	return 0;
}

// Parse single number from *pp.
static int
parsenum1(char **pp, vlong *np)
{
	char *p;

	p = *pp;
	if(*p != &#39;-&#39; &amp;&amp; (*p &lt; &#39;0&#39; || *p &gt; &#39;9&#39;)) {
		fprint(2, &#34;parsenum expected minus or digit\n&#34;);
		return -1;
	}
	*np = strtoll(p, pp, 10);
	return 0;
}

// Parse type number - either single number or (n1, n2).
static int
parsetypenum(char **pp, vlong *n1p, vlong *n2p)
{
	char *p;

	p = *pp;
	if(*p == &#39;(&#39;) {
		p++;
		if(parsenum1(&amp;p, n1p) &lt; 0)
			return -1;
		if(*p++ != &#39;,&#39;) {
			if(stabsdebug)
				fprint(2, &#34;parsetypenum expected comma\n&#34;);
			return -1;
		}
		if(parsenum1(&amp;p, n2p) &lt; 0)
			return -1;
		if(*p++ != &#39;)&#39;) {
			if(stabsdebug)
				fprint(2, &#34;parsetypenum expected right paren\n&#34;);
			return -1;
		}
		*pp = p;
		return 0;
	}

	if(parsenum1(&amp;p, n1p) &lt; 0)
		return -1;
	*n2p = 0;
	*pp = p;
	return 0;
}

// Integer types are represented in stabs as a &#34;range&#34;
// type with a lo and a hi value.  The lo and hi used to
// be lo and hi for the type, but there are now odd
// extensions for floating point and 64-bit numbers.
//
// Have to keep signs separate from values because
// Int64&#39;s lo is -0.
typedef struct Intrange Intrange;
struct Intrange
{
	int signlo;	// sign of lo
	vlong lo;
	int signhi;	// sign of hi
	vlong hi;
	int kind;
};

// NOTE(rsc): Iant says that these might be different depending
// on the gcc mode, though I haven&#39;t observed this yet.
Intrange intranges[] = {
	&#39;+&#39;, 0, &#39;+&#39;, 127, Int8,	// char
	&#39;-&#39;, 128, &#39;+&#39;, 127, Int8,	// signed char
	&#39;+&#39;, 0, &#39;+&#39;, 255, Uint8,
	&#39;-&#39;, 32768, &#39;+&#39;, 32767, Int16,
	&#39;+&#39;, 0, &#39;+&#39;, 65535, Uint16,
	&#39;-&#39;, 2147483648LL, &#39;+&#39;, 2147483647LL, Int32,
	&#39;+&#39;, 0, &#39;+&#39;, 4294967295LL, Uint32,

	// abnormal cases
	&#39;-&#39;, 0, &#39;+&#39;, 4294967295LL, Int64,
	&#39;+&#39;, 0, &#39;-&#39;, 1, Uint64,

	&#39;+&#39;, 4, &#39;+&#39;, 0, Float32,
	&#39;+&#39;, 8, &#39;+&#39;, 0, Float64,
	&#39;+&#39;, 16, &#39;+&#39;, 0, Void,
};

static int kindsize[] = {
	0,
	0,
	8,
	8,
	16,
	16,
	32,
	32,
	64,
	64,
};

// Parse a single type definition from *pp.
static Type*
parsedef(char **pp, char *name)
{
	char *p;
	Type *t, *tt;
	int i, signlo, signhi;
	vlong n1, n2, lo, hi;
	Field *f;
	Intrange *r;

	p = *pp;

	// reference to another type?
	if(isdigit(*p) || *p == &#39;(&#39;) {
		if(parsetypenum(&amp;p, &amp;n1, &amp;n2) &lt; 0)
			return nil;
		t = typebynum(n1, n2);
		if(name &amp;&amp; t-&gt;name == nil) {
			t-&gt;name = name;
			// save definitions of names beginning with $
			if(name[0] == &#39;$&#39; &amp;&amp; !t-&gt;saved) {
				typ = erealloc(typ, (ntyp+1)*sizeof typ[0]);
				typ[ntyp] = t;
				ntyp++;
			}
		}

		// is there an =def suffix?
		if(*p == &#39;=&#39;) {
			p++;
			tt = parsedef(&amp;p, name);
			if(tt == nil)
				return nil;

			if(tt == t) {
				tt-&gt;kind = Void;
			} else {
				t-&gt;type = tt;
				t-&gt;kind = Typedef;
			}

			// assign given name, but do not record in typ.
			// assume the name came from a typedef
			// which will be recorded.
			if(name)
				tt-&gt;name = name;
		}

		*pp = p;
		return t;
	}

	// otherwise a type literal.  first letter identifies kind
	t = emalloc(sizeof *t);
	switch(*p) {
	default:
		fprint(2, &#34;unknown type char %c\n&#34;, *p);
		*pp = &#34;&#34;;
		return t;

	case &#39;*&#39;:	// pointer
		p++;
		t-&gt;kind = Ptr;
		tt = parsedef(&amp;p, nil);
		if(tt == nil)
			return nil;
		t-&gt;type = tt;
		break;

	case &#39;a&#39;:	// array
		p++;
		t-&gt;kind = Array;
		// index type
		tt = parsedef(&amp;p, nil);
		if(tt == nil)
			return nil;
		t-&gt;size = tt-&gt;size;
		// element type
		tt = parsedef(&amp;p, nil);
		if(tt == nil)
			return nil;
		t-&gt;type = tt;
		break;

	case &#39;e&#39;:	// enum type - record $names in con array.
		p++;
		for(;;) {
			if(*p == &#39;\0&#39;)
				return nil;
			if(*p == &#39;;&#39;) {
				p++;
				break;
			}
			if(parsename(&amp;p, &amp;name) &lt; 0)
				return nil;
			if(parsenum1(&amp;p, &amp;n1) &lt; 0)
				return nil;
			if(name[0] == &#39;$&#39;) {
				con = erealloc(con, (ncon+1)*sizeof con[0]);
				name++;
				con[ncon].name = name;
				con[ncon].value = n1;
				ncon++;
			}
			if(*p != &#39;,&#39;)
				return nil;
			p++;
		}
		break;

	case &#39;f&#39;:	// function
		p++;
		if(parsedef(&amp;p, nil) == nil)
			return nil;
		break;

	case &#39;r&#39;:	// sub-range (used for integers)
		p++;
		if(parsedef(&amp;p, nil) == nil)
			return nil;
		// usually, the return from parsedef == t, but not always.

		if(*p != &#39;;&#39; || *++p == &#39;;&#39;) {
			if(stabsdebug)
				fprint(2, &#34;range expected number: %s\n&#34;, p);
			return nil;
		}
		if(*p == &#39;-&#39;) {
			signlo = &#39;-&#39;;
			p++;
		} else
			signlo = &#39;+&#39;;
		lo = strtoll(p, &amp;p, 10);
		if(*p != &#39;;&#39; || *++p == &#39;;&#39;) {
			if(stabsdebug)
				fprint(2, &#34;range expected number: %s\n&#34;, p);
			return nil;
		}
		if(*p == &#39;-&#39;) {
			signhi = &#39;-&#39;;
			p++;
		} else
			signhi = &#39;+&#39;;
		hi = strtoll(p, &amp;p, 10);
		if(*p != &#39;;&#39;) {
			if(stabsdebug)
				fprint(2, &#34;range expected trailing semi: %s\n&#34;, p);
			return nil;
		}
		p++;
		t-&gt;size = hi+1;	// might be array size
		for(i=0; i&lt;nelem(intranges); i++) {
			r = &amp;intranges[i];
			if(r-&gt;signlo == signlo &amp;&amp; r-&gt;signhi == signhi &amp;&amp; r-&gt;lo == lo &amp;&amp; r-&gt;hi == hi) {
				t-&gt;kind = r-&gt;kind;
				break;
			}
		}
		break;

	case &#39;s&#39;:	// struct
	case &#39;u&#39;:	// union
		t-&gt;kind = Struct;
		if(*p == &#39;u&#39;)
			t-&gt;kind = Union;

		// assign given name, but do not record in typ.
		// assume the name came from a typedef
		// which will be recorded.
		if(name)
			t-&gt;name = name;
		p++;
		if(parsenum1(&amp;p, &amp;n1) &lt; 0)
			return nil;
		t-&gt;size = n1;
		for(;;) {
			if(*p == &#39;\0&#39;)
				return nil;
			if(*p == &#39;;&#39;) {
				p++;
				break;
			}
			t-&gt;f = erealloc(t-&gt;f, (t-&gt;nf+1)*sizeof t-&gt;f[0]);
			f = &amp;t-&gt;f[t-&gt;nf];
			if(parsename(&amp;p, &amp;f-&gt;name) &lt; 0)
				return nil;
			f-&gt;type = parsedef(&amp;p, nil);
			if(f-&gt;type == nil)
				return nil;
			if(*p != &#39;,&#39;) {
				fprint(2, &#34;expected comma after def of %s:\n%s\n&#34;, f-&gt;name, p);
				return nil;
			}
			p++;
			if(parsenum1(&amp;p, &amp;n1) &lt; 0)
				return nil;
			f-&gt;offset = n1;
			if(*p != &#39;,&#39;) {
				fprint(2, &#34;expected comma after offset of %s:\n%s\n&#34;, f-&gt;name, p);
				return nil;
			}
			p++;
			if(parsenum1(&amp;p, &amp;n1) &lt; 0)
				return nil;
			f-&gt;size = n1;
			if(*p != &#39;;&#39;) {
				fprint(2, &#34;expected semi after size of %s:\n%s\n&#34;, f-&gt;name, p);
				return nil;
			}

			// rewrite
			//	uint32 x : 8;
			// into
			//	uint8 x;
			// hooray for bitfields.
			while(f-&gt;type-&gt;kind == Typedef)
				f-&gt;type = f-&gt;type-&gt;type;
			while(Int16 &lt;= f-&gt;type-&gt;kind &amp;&amp; f-&gt;type-&gt;kind &lt;= Uint64 &amp;&amp; kindsize[f-&gt;type-&gt;kind] &gt; f-&gt;size) {
				tt = emalloc(sizeof *tt);
				*tt = *f-&gt;type;
				f-&gt;type = tt;
				f-&gt;type-&gt;kind -= 2;
			}
			p++;
			t-&gt;nf++;
		}
		break;

	case &#39;x&#39;:
		// reference to struct, union not yet defined.
		p++;
		switch(*p) {
		case &#39;s&#39;:
			t-&gt;kind = Struct;
			break;
		case &#39;u&#39;:
			t-&gt;kind = Union;
			break;
		default:
			fprint(2, &#34;unknown x type char x%c&#34;, *p);
			*pp = &#34;&#34;;
			return t;
		}
		if(parsename(&amp;p, &amp;t-&gt;name) &lt; 0)
			return nil;
		break;
	}
	*pp = p;
	return t;
}


// Parse a stab type in p, saving info in the type hash table
// and also in the list of recorded types if appropriate.
void
parsestabtype(char *p)
{
	char *p0, *name;

	p0 = p;

	// p is the quoted string output from gcc -gstabs on a .stabs line.
	//	name:t(1,2)
	//	name:t(1,2)=def
	if(parsename(&amp;p, &amp;name) &lt; 0) {
	Bad:
		// Use fprint instead of sysfatal to avoid
		// sysfatal&#39;s internal buffer size limit.
		fprint(2, &#34;cannot parse stabs type:\n%s\n(at %s)\n&#34;, p0, p);
		sysfatal(&#34;stabs parse&#34;);
	}
	if(*p != &#39;t&#39; &amp;&amp; *p != &#39;T&#39;)
		goto Bad;
	p++;

	// parse the definition.
	if(name[0] == &#39;\0&#39;)
		name = nil;
	if(parsedef(&amp;p, name) == nil)
		goto Bad;
	if(*p != &#39;\0&#39;)
		goto Bad;
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

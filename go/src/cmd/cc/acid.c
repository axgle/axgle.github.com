<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cc/acid.c</title>

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
	<li>Thu Nov 12 15:49:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cc/acid.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/cc/acid.c
// http://code.google.com/p/inferno-os/source/browse/utils/cc/acid.c
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

#include &#34;cc.h&#34;

static char *kwd[] =
{
	&#34;$adt&#34;, &#34;$aggr&#34;, &#34;$append&#34;, &#34;$complex&#34;, &#34;$defn&#34;,
	&#34;$delete&#34;, &#34;$do&#34;, &#34;$else&#34;, &#34;$eval&#34;, &#34;$head&#34;, &#34;$if&#34;,
	&#34;$local&#34;, &#34;$loop&#34;, &#34;$return&#34;, &#34;$tail&#34;, &#34;$then&#34;,
	&#34;$union&#34;, &#34;$whatis&#34;, &#34;$while&#34;,
};

char*
amap(char *s)
{
	int i, bot, top, new;

	bot = 0;
	top = bot + nelem(kwd) - 1;
	while(bot &lt;= top){
		new = bot + (top - bot)/2;
		i = strcmp(kwd[new]+1, s);
		if(i == 0)
			return kwd[new];

		if(i &lt; 0)
			bot = new + 1;
		else
			top = new - 1;
	}
	return s;
}

Sym*
acidsue(Type *t)
{
	int h;
	Sym *s;

	if(t != T)
	for(h=0; h&lt;nelem(hash); h++)
		for(s = hash[h]; s != S; s = s-&gt;link)
			if(s-&gt;suetag &amp;&amp; s-&gt;suetag-&gt;link == t)
				return s;
	return 0;
}

Sym*
acidfun(Type *t)
{
	int h;
	Sym *s;

	for(h=0; h&lt;nelem(hash); h++)
		for(s = hash[h]; s != S; s = s-&gt;link)
			if(s-&gt;type == t)
				return s;
	return 0;
}

char	acidchar[NTYPE];
Init	acidcinit[] =
{
	TCHAR,		&#39;C&#39;,	0,
	TUCHAR,		&#39;b&#39;,	0,
	TSHORT,		&#39;d&#39;,	0,
	TUSHORT,	&#39;u&#39;,	0,
	TLONG,		&#39;D&#39;,	0,
	TULONG,		&#39;U&#39;,	0,
	TVLONG,		&#39;V&#39;,	0,
	TUVLONG,	&#39;W&#39;,	0,
	TFLOAT,		&#39;f&#39;,	0,
	TDOUBLE,	&#39;F&#39;,	0,
	TARRAY,		&#39;a&#39;,	0,
	TIND,		&#39;X&#39;,	0,
	-1,		0,	0,
};

static void
acidinit(void)
{
	Init *p;

	for(p=acidcinit; p-&gt;code &gt;= 0; p++)
		acidchar[p-&gt;code] = p-&gt;value;

	acidchar[TINT] = acidchar[TLONG];
	acidchar[TUINT] = acidchar[TULONG];
	if(types[TINT]-&gt;width != types[TLONG]-&gt;width) {
		acidchar[TINT] = acidchar[TSHORT];
		acidchar[TUINT] = acidchar[TUSHORT];
		if(types[TINT]-&gt;width != types[TSHORT]-&gt;width)
			warn(Z, &#34;acidmember int not long or short&#34;);
	}
	if(types[TIND]-&gt;width == types[TUVLONG]-&gt;width)
		acidchar[TIND] = &#39;Y&#39;;
	
}

void
acidmember(Type *t, int32 off, int flag)
{
	Sym *s, *s1;
	Type *l;
	static int acidcharinit = 0;

	if(acidcharinit == 0) {
		acidinit();
		acidcharinit = 1;
	}
	s = t-&gt;sym;
	switch(t-&gt;etype) {
	default:
		Bprint(&amp;outbuf, &#34;	T%d\n&#34;, t-&gt;etype);
		break;

	case TIND:
		if(s == S)
			break;
		l = t-&gt;link;
		if(flag) { 
			if(typesu[l-&gt;etype]) {
				s1 = acidsue(l-&gt;link);
				if(s1 != S) {
					Bprint(&amp;outbuf, &#34;	&#39;A&#39; %s %ld %s;\n&#34;,
						amap(s1-&gt;name),
						t-&gt;offset+off, amap(s-&gt;name));
					break;
				}
			}
		} else {
			l = t-&gt;link;
			s1 = S;
			if(typesu[l-&gt;etype])
				s1 = acidsue(l-&gt;link);
			if(s1 != S) {
				Bprint(&amp;outbuf,
					&#34;\tprint(indent, \&#34;%s\t(%s)\&#34;, addr.%s\\X, \&#34;\\n\&#34;);\n&#34;,
					amap(s-&gt;name), amap(s1-&gt;name), amap(s-&gt;name));
			} else {
				Bprint(&amp;outbuf,
					&#34;\tprint(indent, \&#34;%s\t\&#34;, addr.%s\\X, \&#34;\\n\&#34;);\n&#34;,
					amap(s-&gt;name), amap(s-&gt;name));
			}
			break;
		}

	case TINT:
	case TUINT:
	case TCHAR:
	case TUCHAR:
	case TSHORT:
	case TUSHORT:
	case TLONG:
	case TULONG:
	case TVLONG:
	case TUVLONG:
	case TFLOAT:
	case TDOUBLE:
	case TARRAY:
		if(s == S)
			break;
		if(flag) {
			Bprint(&amp;outbuf, &#34;	&#39;%c&#39; %ld %s;\n&#34;,
			acidchar[t-&gt;etype], t-&gt;offset+off, amap(s-&gt;name));
		} else {
			Bprint(&amp;outbuf, &#34;\tprint(indent, \&#34;%s\t\&#34;, addr.%s, \&#34;\\n\&#34;);\n&#34;,
				amap(s-&gt;name), amap(s-&gt;name));
		}
		break;

	case TSTRUCT:
	case TUNION:
		s1 = acidsue(t-&gt;link);
		if(s1 == S)
			break;
		if(flag) {
			if(s == S) {
				Bprint(&amp;outbuf, &#34;	{\n&#34;);
				for(l = t-&gt;link; l != T; l = l-&gt;down)
					acidmember(l, t-&gt;offset+off, flag);
				Bprint(&amp;outbuf, &#34;	};\n&#34;);
			} else {
				Bprint(&amp;outbuf, &#34;	%s %ld %s;\n&#34;,
					amap(s1-&gt;name),
					t-&gt;offset+off, amap(s-&gt;name));
			}
		} else {
			if(s != S) {
				Bprint(&amp;outbuf, &#34;\tprint(indent, \&#34;%s %s {\\n\&#34;);\n&#34;,
					amap(s1-&gt;name), amap(s-&gt;name));
				Bprint(&amp;outbuf, &#34;\tindent_%s(addr.%s, indent+\&#34;\\t\&#34;);\n&#34;,
					amap(s1-&gt;name), amap(s-&gt;name));
				Bprint(&amp;outbuf, &#34;\tprint(indent, \&#34;}\\n\&#34;);\n&#34;);
			} else {
				Bprint(&amp;outbuf, &#34;\tprint(indent, \&#34;%s {\\n\&#34;);\n&#34;,
					amap(s1-&gt;name));
				Bprint(&amp;outbuf, &#34;\tindent_%s(addr+%ld, indent+\&#34;\\t\&#34;);\n&#34;,
					amap(s1-&gt;name), t-&gt;offset+off);
				Bprint(&amp;outbuf, &#34;\tprint(indent, \&#34;}\\n\&#34;);\n&#34;);
			}
		}
		break;
	}
}

void
acidtype(Type *t)
{
	Sym *s;
	Type *l;
	Io *i;
	int n;
	char *an;

	if(!debug[&#39;a&#39;])
		return;
	if(debug[&#39;a&#39;] &gt; 1) {
		n = 0;
		for(i=iostack; i; i=i-&gt;link)
			n++;
		if(n &gt; 1)
			return;
	}
	s = acidsue(t-&gt;link);
	if(s == S)
		return;
	switch(t-&gt;etype) {
	default:
		Bprint(&amp;outbuf, &#34;T%d\n&#34;, t-&gt;etype);
		return;

	case TUNION:
	case TSTRUCT:
		if(debug[&#39;s&#39;])
			goto asmstr;
		an = amap(s-&gt;name);
		Bprint(&amp;outbuf, &#34;sizeof%s = %ld;\n&#34;, an, t-&gt;width);
		Bprint(&amp;outbuf, &#34;aggr %s\n{\n&#34;, an);
		for(l = t-&gt;link; l != T; l = l-&gt;down)
			acidmember(l, 0, 1);
		Bprint(&amp;outbuf, &#34;};\n\n&#34;);

		Bprint(&amp;outbuf, &#34;defn\n%s(addr) {\n\tindent_%s(addr, \&#34;\\t\&#34;);\n}\n&#34;, an, an);
		Bprint(&amp;outbuf, &#34;defn\nindent_%s(addr, indent) {\n\tcomplex %s addr;\n&#34;, an, an);
		for(l = t-&gt;link; l != T; l = l-&gt;down)
			acidmember(l, 0, 0);
		Bprint(&amp;outbuf, &#34;};\n\n&#34;);
		break;
	asmstr:
		if(s == S)
			break;
		for(l = t-&gt;link; l != T; l = l-&gt;down)
			if(l-&gt;sym != S)
				Bprint(&amp;outbuf, &#34;#define\t%s.%s\t%ld\n&#34;,
					s-&gt;name,
					l-&gt;sym-&gt;name,
					l-&gt;offset);
		break;
	}
}

void
acidvar(Sym *s)
{
	int n;
	Io *i;
	Type *t;
	Sym *s1, *s2;

	if(!debug[&#39;a&#39;] || debug[&#39;s&#39;])
		return;
	if(debug[&#39;a&#39;] &gt; 1) {
		n = 0;
		for(i=iostack; i; i=i-&gt;link)
			n++;
		if(n &gt; 1)
			return;
	}
	t = s-&gt;type;
	while(t &amp;&amp; t-&gt;etype == TIND)
		t = t-&gt;link;
	if(t == T)
		return;
	if(t-&gt;etype == TENUM) {
		Bprint(&amp;outbuf, &#34;%s = &#34;, amap(s-&gt;name));
		if(!typefd[t-&gt;etype])
			Bprint(&amp;outbuf, &#34;%lld;\n&#34;, s-&gt;vconst);
		else
			Bprint(&amp;outbuf, &#34;%f\n;&#34;, s-&gt;fconst);
		return;
	}
	if(!typesu[t-&gt;etype])
		return;
	s1 = acidsue(t-&gt;link);
	if(s1 == S)
		return;
	switch(s-&gt;class) {
	case CAUTO:
	case CPARAM:
		s2 = acidfun(thisfn);
		if(s2)
			Bprint(&amp;outbuf, &#34;complex %s %s:%s;\n&#34;,
				amap(s1-&gt;name), amap(s2-&gt;name), amap(s-&gt;name));
		break;
	
	case CSTATIC:
	case CEXTERN:
	case CGLOBL:
	case CLOCAL:
		Bprint(&amp;outbuf, &#34;complex %s %s;\n&#34;,
			amap(s1-&gt;name), amap(s-&gt;name));
		break;
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/ld/lib.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/ld/lib.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Inferno utils/6l/obj.c
// http://code.google.com/p/inferno-os/source/browse/utils/6l/obj.c
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

#include	&#34;l.h&#34;
#include	&#34;lib.h&#34;
#include	&lt;ar.h&gt;

char	symname[]	= SYMDEF;
char*	libdir[16] = { &#34;.&#34; };
int	nlibdir = 1;
int	cout = -1;

char*	goroot;
char*	goarch;
char*	goos;

void
Lflag(char *arg)
{
	if(nlibdir &gt;= nelem(libdir)-1) {
		print(&#34;too many -L&#39;s: %d\n&#34;, nlibdir);
		usage();
	}
	libdir[nlibdir++] = arg;
}

void
libinit(void)
{
	mywhatsys();	// get goroot, goarch, goos
	if(strcmp(goarch, thestring) != 0)
		print(&#34;goarch is not known: %s\n&#34;, goarch);

	// add goroot to the end of the libdir list.
	libdir[nlibdir++] = smprint(&#34;%s/pkg/%s_%s&#34;, goroot, goos, goarch);

	unlink(outfile);
	cout = create(outfile, 1, 0775);
	if(cout &lt; 0) {
		diag(&#34;cannot create %s&#34;, outfile);
		errorexit();
	}

	if(INITENTRY == nil) {
		INITENTRY = mal(strlen(goarch)+strlen(goos)+10);
		sprint(INITENTRY, &#34;_rt0_%s_%s&#34;, goarch, goos);
	}
	lookup(INITENTRY, 0)-&gt;type = SXREF;
}

void
errorexit(void)
{
	if(nerrors) {
		if(cout &gt;= 0)
			remove(outfile);
		exits(&#34;error&#34;);
	}
	exits(0);
}

void
addlib(char *src, char *obj)
{
	char name[1024], pname[1024], comp[256], *p;
	int i, search;

	if(histfrogp &lt;= 0)
		return;

	search = 0;
	if(histfrog[0]-&gt;name[1] == &#39;/&#39;) {
		sprint(name, &#34;&#34;);
		i = 1;
	} else
	if(histfrog[0]-&gt;name[1] == &#39;.&#39;) {
		sprint(name, &#34;.&#34;);
		i = 0;
	} else {
		sprint(name, &#34;&#34;);
		i = 0;
		search = 1;
	}

	for(; i&lt;histfrogp; i++) {
		snprint(comp, sizeof comp, histfrog[i]-&gt;name+1);
		for(;;) {
			p = strstr(comp, &#34;$O&#34;);
			if(p == 0)
				break;
			memmove(p+1, p+2, strlen(p+2)+1);
			p[0] = thechar;
		}
		for(;;) {
			p = strstr(comp, &#34;$M&#34;);
			if(p == 0)
				break;
			if(strlen(comp)+strlen(thestring)-2+1 &gt;= sizeof comp) {
				diag(&#34;library component too long&#34;);
				return;
			}
			memmove(p+strlen(thestring), p+2, strlen(p+2)+1);
			memmove(p, thestring, strlen(thestring));
		}
		if(strlen(name) + strlen(comp) + 3 &gt;= sizeof(name)) {
			diag(&#34;library component too long&#34;);
			return;
		}
		strcat(name, &#34;/&#34;);
		strcat(name, comp);
	}

	if(search) {
		// try dot, -L &#34;libdir&#34;, and then goroot.
		for(i=0; i&lt;nlibdir; i++) {
			snprint(pname, sizeof pname, &#34;%s/%s&#34;, libdir[i], name);
			if(access(pname, AEXIST) &gt;= 0)
				break;
		}
		strcpy(name, pname);
	}
	cleanname(name);
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f addlib: %s %s pulls in %s\n&#34;, cputime(), obj, src, name);

	for(i=0; i&lt;libraryp; i++)
		if(strcmp(name, library[i]) == 0)
			return;
	if(libraryp == nlibrary){
		nlibrary = 50 + 2*libraryp;
		library = realloc(library, sizeof library[0] * nlibrary);
		libraryobj = realloc(libraryobj, sizeof libraryobj[0] * nlibrary);
	}

	p = mal(strlen(name) + 1);
	strcpy(p, name);
	library[libraryp] = p;
	p = mal(strlen(obj) + 1);
	strcpy(p, obj);
	libraryobj[libraryp] = p;
	libraryp++;
}

void
loadlib(void)
{
	int i;
	int32 h;
	Sym *s;
	char *a;

loop:
	xrefresolv = 0;
	for(i=0; i&lt;libraryp; i++) {
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f autolib: %s (from %s)\n&#34;, cputime(), library[i], libraryobj[i]);
		objfile(library[i]);
	}
	if(xrefresolv)
	for(h=0; h&lt;nelem(hash); h++)
	for(s = hash[h]; s != S; s = s-&gt;link)
		if(s-&gt;type == SXREF)
			goto loop;

	i = strlen(goroot)+strlen(goarch)+strlen(goos)+20;
	a = mal(i);
	snprint(a, i, &#34;%s/pkg/%s_%s/runtime.a&#34;, goroot, goos, goarch);
	objfile(a);
}

void
objfile(char *file)
{
	int32 off, esym, cnt, l;
	int work;
	Biobuf *f;
	Sym *s;
	char magbuf[SARMAG];
	char name[100], pname[150];
	struct ar_hdr arhdr;
	char *e, *start, *stop;

	if(file[0] == &#39;-&#39; &amp;&amp; file[1] == &#39;l&#39;) {	// TODO: fix this
		if(debug[&#39;9&#39;])
			sprint(name, &#34;/%s/lib/lib&#34;, thestring);
		else
			sprint(name, &#34;/usr/%clib/lib&#34;, thechar);
		strcat(name, file+2);
		strcat(name, &#34;.a&#34;);
		file = name;
	}
	if(debug[&#39;v&#39;])
		Bprint(&amp;bso, &#34;%5.2f ldobj: %s\n&#34;, cputime(), file);
	Bflush(&amp;bso);
	f = Bopen(file, 0);
	if(f == nil) {
		diag(&#34;cannot open file: %s&#34;, file);
		errorexit();
	}
	l = Bread(f, magbuf, SARMAG);
	if(l != SARMAG || strncmp(magbuf, ARMAG, SARMAG)){
		/* load it as a regular file */
		l = Bseek(f, 0L, 2);
		Bseek(f, 0L, 0);
		ldobj(f, l, file);
		Bterm(f);
		return;
	}

	l = Bread(f, &amp;arhdr, SAR_HDR);
	if(l != SAR_HDR) {
		diag(&#34;%s: short read on archive file symbol header&#34;, file);
		goto out;
	}
	if(strncmp(arhdr.name, symname, strlen(symname))) {
		diag(&#34;%s: first entry not symbol header&#34;, file);
		goto out;
	}

	esym = SARMAG + SAR_HDR + atolwhex(arhdr.size);
	off = SARMAG + SAR_HDR;

	/*
	 * just bang the whole symbol file into memory
	 */
	Bseek(f, off, 0);
	cnt = esym - off;
	start = mal(cnt + 10);
	cnt = Bread(f, start, cnt);
	if(cnt &lt;= 0){
		Bterm(f);
		return;
	}
	stop = &amp;start[cnt];
	memset(stop, 0, 10);

	work = 1;
	while(work) {
		if(debug[&#39;v&#39;])
			Bprint(&amp;bso, &#34;%5.2f library pass: %s\n&#34;, cputime(), file);
		Bflush(&amp;bso);
		work = 0;
		for(e = start; e &lt; stop; e = strchr(e+5, 0) + 1) {
			s = lookup(e+5, 0);
			if(s-&gt;type != SXREF)
				continue;
			sprint(pname, &#34;%s(%s)&#34;, file, s-&gt;name);
			if(debug[&#39;v&#39;])
				Bprint(&amp;bso, &#34;%5.2f library: %s\n&#34;, cputime(), pname);
			Bflush(&amp;bso);
			l = e[1] &amp; 0xff;
			l |= (e[2] &amp; 0xff) &lt;&lt; 8;
			l |= (e[3] &amp; 0xff) &lt;&lt; 16;
			l |= (e[4] &amp; 0xff) &lt;&lt; 24;
			Bseek(f, l, 0);
			l = Bread(f, &amp;arhdr, SAR_HDR);
			if(l != SAR_HDR)
				goto bad;
			if(strncmp(arhdr.fmag, ARFMAG, sizeof(arhdr.fmag)))
				goto bad;
			l = SARNAME;
			while(l &gt; 0 &amp;&amp; arhdr.name[l-1] == &#39; &#39;)
				l--;
			sprint(pname, &#34;%s(%.*s)&#34;, file, l, arhdr.name);
			l = atolwhex(arhdr.size);
			ldobj(f, l, pname);
			if(s-&gt;type == SXREF) {
				diag(&#34;%s: failed to load: %s&#34;, file, s-&gt;name);
				errorexit();
			}
			work = 1;
			xrefresolv = 1;
		}
	}
	return;

bad:
	diag(&#34;%s: bad or out of date archive&#34;, file);
out:
	Bterm(f);
}

void
ldobj(Biobuf *f, int64 len, char *pn)
{
	static int files;
	static char **filen;
	char **nfilen, *line;
	int n, c1, c2, c3;
	vlong import0, import1, eof;
	char src[1024];

	eof = Boffset(f) + len;
	src[0] = &#39;\0&#39;;

	if((files&amp;15) == 0){
		nfilen = malloc((files+16)*sizeof(char*));
		memmove(nfilen, filen, files*sizeof(char*));
		free(filen);
		filen = nfilen;
	}
	pn = strdup(pn);
	filen[files++] = pn;


	/* check the header */
	line = Brdline(f, &#39;\n&#39;);
	if(line == nil) {
		if(Blinelen(f) &gt; 0) {
			diag(&#34;%s: malformed object file&#34;, pn);
			return;
		}
		goto eof;
	}
	n = Blinelen(f) - 1;
	if(n != strlen(thestring) || strncmp(line, thestring, n) != 0) {
		if(line)
			line[n] = &#39;\0&#39;;
		diag(&#34;file not %s [%s]\n&#34;, thestring, line);
		return;
	}

	/* skip over exports and other info -- ends with \n!\n */
	import0 = Boffset(f);
	c1 = &#39;\n&#39;;	// the last line ended in \n
	c2 = Bgetc(f);
	c3 = Bgetc(f);
	while(c1 != &#39;\n&#39; || c2 != &#39;!&#39; || c3 != &#39;\n&#39;) {
		c1 = c2;
		c2 = c3;
		c3 = Bgetc(f);
		if(c3 == Beof)
			goto eof;
	}
	import1 = Boffset(f);

	Bseek(f, import0, 0);
	ldpkg(f, import1 - import0 - 2, pn);	// -2 for !\n
	Bseek(f, import1, 0);

	ldobj1(f, eof - Boffset(f), pn);
	return;

eof:
	diag(&#34;truncated object file: %s&#34;, pn);
}

Sym*
lookup(char *symb, int v)
{
	Sym *s;
	char *p;
	int32 h;
	int l, c;

	h = v;
	for(p=symb; c = *p; p++)
		h = h+h+h + c;
	l = (p - symb) + 1;
	// not if(h &lt; 0) h = ~h, because gcc 4.3 -O2 miscompiles it.
	h &amp;= 0xffffff;
	h %= NHASH;
	for(s = hash[h]; s != S; s = s-&gt;link)
		if(s-&gt;version == v)
		if(memcmp(s-&gt;name, symb, l) == 0)
			return s;

	s = mal(sizeof(*s));
	if(debug[&#39;v&#39;] &gt; 1)
		Bprint(&amp;bso, &#34;lookup %s\n&#34;, symb);

	s-&gt;name = mal(l + 1);
	memmove(s-&gt;name, symb, l);

	s-&gt;link = hash[h];
	s-&gt;type = 0;
	s-&gt;version = v;
	s-&gt;value = 0;
	s-&gt;sig = 0;
	hash[h] = s;
	nsymbol++;
	return s;
}

void
copyhistfrog(char *buf, int nbuf)
{
	char *p, *ep;
	int i;

	p = buf;
	ep = buf + nbuf;
	i = 0;
	for(i=0; i&lt;histfrogp; i++) {
		p = seprint(p, ep, &#34;%s&#34;, histfrog[i]-&gt;name+1);
		if(i+1&lt;histfrogp &amp;&amp; (p == buf || p[-1] != &#39;/&#39;))
			p = seprint(p, ep, &#34;/&#34;);
	}
}

void
addhist(int32 line, int type)
{
	Auto *u;
	Sym *s;
	int i, j, k;

	u = mal(sizeof(Auto));
	s = mal(sizeof(Sym));
	s-&gt;name = mal(2*(histfrogp+1) + 1);

	u-&gt;asym = s;
	u-&gt;type = type;
	u-&gt;aoffset = line;
	u-&gt;link = curhist;
	curhist = u;

	s-&gt;name[0] = 0;
	j = 1;
	for(i=0; i&lt;histfrogp; i++) {
		k = histfrog[i]-&gt;value;
		s-&gt;name[j+0] = k&gt;&gt;8;
		s-&gt;name[j+1] = k;
		j += 2;
	}
	s-&gt;name[j] = 0;
	s-&gt;name[j+1] = 0;
}

void
histtoauto(void)
{
	Auto *l;

	while(l = curhist) {
		curhist = l-&gt;link;
		l-&gt;link = curauto;
		curauto = l;
	}
}

void
collapsefrog(Sym *s)
{
	int i;

	/*
	 * bad encoding of path components only allows
	 * MAXHIST components. if there is an overflow,
	 * first try to collapse xxx/..
	 */
	for(i=1; i&lt;histfrogp; i++)
		if(strcmp(histfrog[i]-&gt;name+1, &#34;..&#34;) == 0) {
			memmove(histfrog+i-1, histfrog+i+1,
				(histfrogp-i-1)*sizeof(histfrog[0]));
			histfrogp--;
			goto out;
		}

	/*
	 * next try to collapse .
	 */
	for(i=0; i&lt;histfrogp; i++)
		if(strcmp(histfrog[i]-&gt;name+1, &#34;.&#34;) == 0) {
			memmove(histfrog+i, histfrog+i+1,
				(histfrogp-i-1)*sizeof(histfrog[0]));
			goto out;
		}

	/*
	 * last chance, just truncate from front
	 */
	memmove(histfrog+0, histfrog+1,
		(histfrogp-1)*sizeof(histfrog[0]));

out:
	histfrog[histfrogp-1] = s;
}

void
nuxiinit(void)
{
	int i, c;

	for(i=0; i&lt;4; i++) {
		c = find1(0x04030201L, i+1);
		if(i &lt; 2)
			inuxi2[i] = c;
		if(i &lt; 1)
			inuxi1[i] = c;
		inuxi4[i] = c;
		inuxi8[i] = c;
		inuxi8[i+4] = c+4;
		fnuxi4[i] = c;
		fnuxi8[i] = c;
		fnuxi8[i+4] = c+4;
	}
	if(debug[&#39;v&#39;]) {
		Bprint(&amp;bso, &#34;inuxi = &#34;);
		for(i=0; i&lt;1; i++)
			Bprint(&amp;bso, &#34;%d&#34;, inuxi1[i]);
		Bprint(&amp;bso, &#34; &#34;);
		for(i=0; i&lt;2; i++)
			Bprint(&amp;bso, &#34;%d&#34;, inuxi2[i]);
		Bprint(&amp;bso, &#34; &#34;);
		for(i=0; i&lt;4; i++)
			Bprint(&amp;bso, &#34;%d&#34;, inuxi4[i]);
		Bprint(&amp;bso, &#34; &#34;);
		for(i=0; i&lt;8; i++)
			Bprint(&amp;bso, &#34;%d&#34;, inuxi8[i]);
		Bprint(&amp;bso, &#34;\nfnuxi = &#34;);
		for(i=0; i&lt;4; i++)
			Bprint(&amp;bso, &#34;%d&#34;, fnuxi4[i]);
		Bprint(&amp;bso, &#34; &#34;);
		for(i=0; i&lt;8; i++)
			Bprint(&amp;bso, &#34;%d&#34;, fnuxi8[i]);
		Bprint(&amp;bso, &#34;\n&#34;);
	}
	Bflush(&amp;bso);
}

int
find1(int32 l, int c)
{
	char *p;
	int i;

	p = (char*)&amp;l;
	for(i=0; i&lt;4; i++)
		if(*p++ == c)
			return i;
	return 0;
}

int
find2(int32 l, int c)
{
	union {
		int32 l;
		short p[2];
	} u;
	short *p;
	int i;

	u.l = l;
	p = u.p;
	for(i=0; i&lt;4; i+=2) {
		if(((*p &gt;&gt; 8) &amp; 0xff) == c)
			return i;
		if((*p++ &amp; 0xff) == c)
			return i+1;
	}
	return 0;
}

int32
ieeedtof(Ieee *e)
{
	int exp;
	int32 v;

	if(e-&gt;h == 0)
		return 0;
	exp = (e-&gt;h&gt;&gt;20) &amp; ((1L&lt;&lt;11)-1L);
	exp -= (1L&lt;&lt;10) - 2L;
	v = (e-&gt;h &amp; 0xfffffL) &lt;&lt; 3;
	v |= (e-&gt;l &gt;&gt; 29) &amp; 0x7L;
	if((e-&gt;l &gt;&gt; 28) &amp; 1) {
		v++;
		if(v &amp; 0x800000L) {
			v = (v &amp; 0x7fffffL) &gt;&gt; 1;
			exp++;
		}
	}
	if(exp &lt;= -126 || exp &gt;= 130)
		diag(&#34;double fp to single fp overflow&#34;);
	v |= ((exp + 126) &amp; 0xffL) &lt;&lt; 23;
	v |= e-&gt;h &amp; 0x80000000L;
	return v;
}

double
ieeedtod(Ieee *ieeep)
{
	Ieee e;
	double fr;
	int exp;

	if(ieeep-&gt;h &amp; (1L&lt;&lt;31)) {
		e.h = ieeep-&gt;h &amp; ~(1L&lt;&lt;31);
		e.l = ieeep-&gt;l;
		return -ieeedtod(&amp;e);
	}
	if(ieeep-&gt;l == 0 &amp;&amp; ieeep-&gt;h == 0)
		return 0;
	fr = ieeep-&gt;l &amp; ((1L&lt;&lt;16)-1L);
	fr /= 1L&lt;&lt;16;
	fr += (ieeep-&gt;l&gt;&gt;16) &amp; ((1L&lt;&lt;16)-1L);
	fr /= 1L&lt;&lt;16;
	fr += (ieeep-&gt;h &amp; (1L&lt;&lt;20)-1L) | (1L&lt;&lt;20);
	fr /= 1L&lt;&lt;21;
	exp = (ieeep-&gt;h&gt;&gt;20) &amp; ((1L&lt;&lt;11)-1L);
	exp -= (1L&lt;&lt;10) - 2L;
	return ldexp(fr, exp);
}

void
undefsym(Sym *s)
{
	int n;

	n = imports;
	if(s-&gt;value != 0)
		diag(&#34;value != 0 on SXREF&#34;);
	if(n &gt;= 1&lt;&lt;Rindex)
		diag(&#34;import index %d out of range&#34;, n);
	s-&gt;value = n&lt;&lt;Roffset;
	s-&gt;type = SUNDEF;
	imports++;
}

void
zerosig(char *sp)
{
	Sym *s;

	s = lookup(sp, 0);
	s-&gt;sig = 0;
}

void
readundefs(char *f, int t)
{
	int i, n;
	Sym *s;
	Biobuf *b;
	char *l, buf[256], *fields[64];

	if(f == nil)
		return;
	b = Bopen(f, OREAD);
	if(b == nil){
		diag(&#34;could not open %s: %r&#34;, f);
		errorexit();
	}
	while((l = Brdline(b, &#39;\n&#39;)) != nil){
		n = Blinelen(b);
		if(n &gt;= sizeof(buf)){
			diag(&#34;%s: line too long&#34;, f);
			errorexit();
		}
		memmove(buf, l, n);
		buf[n-1] = &#39;\0&#39;;
		n = getfields(buf, fields, nelem(fields), 1, &#34; \t\r\n&#34;);
		if(n == nelem(fields)){
			diag(&#34;%s: bad format&#34;, f);
			errorexit();
		}
		for(i = 0; i &lt; n; i++){
			s = lookup(fields[i], 0);
			s-&gt;type = SXREF;
			s-&gt;subtype = t;
			if(t == SIMPORT)
				nimports++;
			else
				nexports++;
		}
	}
	Bterm(b);
}

int32
Bget4(Biobuf *f)
{
	uchar p[4];

	if(Bread(f, p, 4) != 4)
		return 0;
	return p[0] | (p[1] &lt;&lt; 8) | (p[2] &lt;&lt; 16) | (p[3] &lt;&lt; 24);
}

void
mywhatsys(void)
{
	char *s;

	goroot = getenv(&#34;GOROOT&#34;);
	goarch = getenv(&#34;GOARCH&#34;);
	goos = getenv(&#34;GOOS&#34;);

	if(goroot == nil) {
		s = getenv(&#34;HOME&#34;);
		if(s == nil)
			s = &#34;/home/ken&#34;;
		goroot = mal(strlen(s) + 10);
		strcpy(goroot, s);
		strcat(goroot, &#34;/go&#34;);
	}
	if(goarch == nil) {
		goarch = &#34;amd64&#34;;
	}
	if(goos == nil) {
		goos = &#34;linux&#34;;
	}
}

int
pathchar(void)
{
	return &#39;/&#39;;
}

static	uchar*	hunk;
static	uint32	nhunk;
#define	NHUNK	(10UL&lt;&lt;20)

void*
mal(uint32 n)
{
	void *v;

	while(n &amp; 7)
		n++;
	if(n &gt; NHUNK) {
		v = malloc(n);
		memset(v, 0, n);
		return v;
	}
	if(n &gt; nhunk) {
		hunk = malloc(NHUNK);
		nhunk = NHUNK;
	}

	v = hunk;
	nhunk -= n;
	hunk += n;

	memset(v, 0, n);
	return v;
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

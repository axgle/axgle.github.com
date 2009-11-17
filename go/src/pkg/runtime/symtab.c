<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/symtab.c</title>

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
	<li>Thu Nov 12 15:50:16 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/symtab.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Runtime symbol table access.  Work in progress.
// The Plan 9 symbol table is not in a particularly convenient form.
// The routines here massage it into a more usable form; eventually
// we&#39;ll change 6l to do this for us, but it is easier to experiment
// here than to change 6l and all the other tools.
//
// The symbol table also needs to be better integrated with the type
// strings table in the future.  This is just a quick way to get started
// and figure out exactly what we want.

#include &#34;runtime.h&#34;

// TODO(rsc): Move this *under* the text segment.
// Then define names for these addresses instead of hard-coding magic ones.
#ifdef _64BIT
#define SYMCOUNTS ((int32*)(0x99LL&lt;&lt;32))	// known to 6l
#define SYMDATA ((byte*)(0x99LL&lt;&lt;32) + 8)
#else
#define SYMCOUNTS ((int32*)(0x99LL&lt;&lt;24))	// known to 8l
#define SYMDATA ((byte*)(0x99LL&lt;&lt;24) + 8)
#endif


// Return a pointer to a byte array containing the symbol table segment.
void
runtime·symdat(Slice *symtab, Slice *pclntab)
{
	Slice *a;
	int32 *v;

	// TODO(rsc): Remove once TODO at top of file is done.
	if(goos != nil &amp;&amp; strcmp((uint8*)goos, (uint8*)&#34;nacl&#34;) == 0) {
		symtab = mal(sizeof *a);
		pclntab = mal(sizeof *a);
		FLUSH(&amp;symtab);
		FLUSH(&amp;pclntab);
		return;
	}

	v = SYMCOUNTS;

	a = mal(sizeof *a);
	a-&gt;len = v[0];
	a-&gt;cap = a-&gt;len;
	a-&gt;array = SYMDATA;
	symtab = a;
	FLUSH(&amp;symtab);

	a = mal(sizeof *a);
	a-&gt;len = v[1];
	a-&gt;cap = a-&gt;len;
	a-&gt;array = SYMDATA + v[0];
	pclntab = a;
	FLUSH(&amp;pclntab);
}

typedef struct Sym Sym;
struct Sym
{
	uintptr value;
	byte symtype;
	byte *name;
//	byte *gotype;
};

// Walk over symtab, calling fn(&amp;s) for each symbol.
static void
walksymtab(void (*fn)(Sym*))
{
	int32 *v;
	byte *p, *ep, *q;
	Sym s;

	// TODO(rsc): Remove once TODO at top of file is done.
	if(goos != nil &amp;&amp; strcmp((uint8*)goos, (uint8*)&#34;nacl&#34;) == 0)
		return;

	v = SYMCOUNTS;
	p = SYMDATA;
	ep = p + v[0];
	while(p &lt; ep) {
		if(p + 7 &gt; ep)
			break;
		s.value = ((uint32)p[0]&lt;&lt;24) | ((uint32)p[1]&lt;&lt;16) | ((uint32)p[2]&lt;&lt;8) | ((uint32)p[3]);
		if(!(p[4]&amp;0x80))
			break;
		s.symtype = p[4] &amp; ~0x80;
		p += 5;
		s.name = p;
		if(s.symtype == &#39;z&#39; || s.symtype == &#39;Z&#39;) {
			// path reference string - skip first byte,
			// then 2-byte pairs ending at two zeros.
			q = p+1;
			for(;;) {
				if(q+2 &gt; ep)
					return;
				if(q[0] == &#39;\0&#39; &amp;&amp; q[1] == &#39;\0&#39;)
					break;
				q += 2;
			}
			p = q+2;
		}else{
			q = mchr(p, &#39;\0&#39;, ep);
			if(q == nil)
				break;
			p = q+1;
		}
		p += 4;	// go type
		fn(&amp;s);
	}
}

// Symtab walker; accumulates info about functions.

static Func *func;
static int32 nfunc;

static byte **fname;
static int32 nfname;

static void
dofunc(Sym *sym)
{
	Func *f;

	switch(sym-&gt;symtype) {
	case &#39;t&#39;:
	case &#39;T&#39;:
		if(strcmp(sym-&gt;name, (byte*)&#34;etext&#34;) == 0)
			break;
		if(func == nil) {
			nfunc++;
			break;
		}
		f = &amp;func[nfunc++];
		f-&gt;name = gostring(sym-&gt;name);
		f-&gt;entry = sym-&gt;value;
		break;
	case &#39;m&#39;:
		if(nfunc &gt; 0 &amp;&amp; func != nil)
			func[nfunc-1].frame = sym-&gt;value;
		break;
	case &#39;p&#39;:
		if(nfunc &gt; 0 &amp;&amp; func != nil) {
			f = &amp;func[nfunc-1];
			// args counts 32-bit words.
			// sym-&gt;value is the arg&#39;s offset.
			// don&#39;t know width of this arg, so assume it is 64 bits.
			if(f-&gt;args &lt; sym-&gt;value/4 + 2)
				f-&gt;args = sym-&gt;value/4 + 2;
		}
		break;
	case &#39;f&#39;:
		if(fname == nil) {
			if(sym-&gt;value &gt;= nfname)
				nfname = sym-&gt;value+1;
			break;
		}
		fname[sym-&gt;value] = sym-&gt;name;
		break;
	}
}

// put together the path name for a z entry.
// the f entries have been accumulated into fname already.
static void
makepath(byte *buf, int32 nbuf, byte *path)
{
	int32 n, len;
	byte *p, *ep, *q;

	if(nbuf &lt;= 0)
		return;

	p = buf;
	ep = buf + nbuf;
	*p = &#39;\0&#39;;
	for(;;) {
		if(path[0] == 0 &amp;&amp; path[1] == 0)
			break;
		n = (path[0]&lt;&lt;8) | path[1];
		path += 2;
		if(n &gt;= nfname)
			break;
		q = fname[n];
		len = findnull(q);
		if(p+1+len &gt;= ep)
			break;
		if(p &gt; buf &amp;&amp; p[-1] != &#39;/&#39;)
			*p++ = &#39;/&#39;;
		mcpy(p, q, len+1);
		p += len;
	}
}

// walk symtab accumulating path names for use by pc/ln table.
// don&#39;t need the full generality of the z entry history stack because
// there are no includes in go (and only sensible includes in our c);
// assume code only appear in top-level files.
static void
dosrcline(Sym *sym)
{
	static byte srcbuf[1000];
	static struct {
		String srcstring;
		int32 aline;
		int32 delta;
	} files[200];
	static int32 incstart;
	static int32 nfunc, nfile, nhist;
	Func *f;
	int32 i;

	switch(sym-&gt;symtype) {
	case &#39;t&#39;:
	case &#39;T&#39;:
		if(strcmp(sym-&gt;name, (byte*)&#34;etext&#34;) == 0)
			break;
		f = &amp;func[nfunc++];
		// find source file
		for(i = 0; i &lt; nfile - 1; i++) {
			if (files[i+1].aline &gt; f-&gt;ln0)
				break;
		}
		f-&gt;src = files[i].srcstring;
		f-&gt;ln0 -= files[i].delta;
		break;
	case &#39;z&#39;:
		if(sym-&gt;value == 1) {
			// entry for main source file for a new object.
			makepath(srcbuf, sizeof srcbuf, sym-&gt;name+1);
			nhist = 0;
			nfile = 0;
			if(nfile == nelem(files))
				return;
			files[nfile].srcstring = gostring(srcbuf);
			files[nfile].aline = 0;
			files[nfile++].delta = 0;
		} else {
			// push or pop of included file.
			makepath(srcbuf, sizeof srcbuf, sym-&gt;name+1);
			if(srcbuf[0] != &#39;\0&#39;) {
				if(nhist++ == 0)
					incstart = sym-&gt;value;
				if(nhist == 0 &amp;&amp; nfile &lt; nelem(files)) {
					// new top-level file
					files[nfile].srcstring = gostring(srcbuf);
					files[nfile].aline = sym-&gt;value;
					// this is &#34;line 0&#34;
					files[nfile++].delta = sym-&gt;value - 1;
				}
			}else{
				if(--nhist == 0)
					files[nfile-1].delta += sym-&gt;value - incstart;
			}
		}
	}
}

enum { PcQuant = 1 };

// Interpret pc/ln table, saving the subpiece for each func.
static void
splitpcln(void)
{
	int32 line;
	uintptr pc;
	byte *p, *ep;
	Func *f, *ef;
	int32 *v;

	// TODO(rsc): Remove once TODO at top of file is done.
	if(goos != nil &amp;&amp; strcmp((uint8*)goos, (uint8*)&#34;nacl&#34;) == 0)
		return;

	// pc/ln table bounds
	v = SYMCOUNTS;
	p = SYMDATA;
	p += v[0];
	ep = p+v[1];

	f = func;
	ef = func + nfunc;
	pc = func[0].entry;	// text base
	f-&gt;pcln.array = p;
	f-&gt;pc0 = pc - PcQuant;
	line = 0;
	for(; p &lt; ep; p++) {
		if(f &lt; ef &amp;&amp; pc &gt; (f+1)-&gt;entry) {
			f-&gt;pcln.len = p - f-&gt;pcln.array;
			f-&gt;pcln.cap = f-&gt;pcln.len;
			f++;
			f-&gt;pcln.array = p;
			f-&gt;pc0 = pc;
			f-&gt;ln0 = line;
		}
		if(*p == 0) {
			// 4 byte add to line
			line += (p[1]&lt;&lt;24) | (p[2]&lt;&lt;16) | (p[3]&lt;&lt;8) | p[4];
			p += 4;
		} else if(*p &lt;= 64) {
			line += *p;
		} else if(*p &lt;= 128) {
			line -= *p - 64;
		} else {
			pc += PcQuant*(*p - 129);
		}
		pc += PcQuant;
	}
	if(f &lt; ef) {
		f-&gt;pcln.len = p - f-&gt;pcln.array;
		f-&gt;pcln.cap = f-&gt;pcln.len;
	}
}


// Return actual file line number for targetpc in func f.
// (Source file is f-&gt;src.)
int32
funcline(Func *f, uint64 targetpc)
{
	byte *p, *ep;
	uintptr pc;
	int32 line;

	p = f-&gt;pcln.array;
	ep = p + f-&gt;pcln.len;
	pc = f-&gt;pc0;
	line = f-&gt;ln0;
	for(; p &lt; ep &amp;&amp; pc &lt;= targetpc; p++) {
		if(*p == 0) {
			line += (p[1]&lt;&lt;24) | (p[2]&lt;&lt;16) | (p[3]&lt;&lt;8) | p[4];
			p += 4;
		} else if(*p &lt;= 64) {
			line += *p;
		} else if(*p &lt;= 128) {
			line -= *p - 64;
		} else {
			pc += PcQuant*(*p - 129);
		}
		pc += PcQuant;
	}
	return line;
}

static void
buildfuncs(void)
{
	extern byte etext[];

	if(func != nil)
		return;
	// count funcs, fnames
	nfunc = 0;
	nfname = 0;
	walksymtab(dofunc);

	// initialize tables
	func = mal((nfunc+1)*sizeof func[0]);
	func[nfunc].entry = (uint64)etext;
	fname = mal(nfname*sizeof fname[0]);
	nfunc = 0;
	walksymtab(dofunc);

	// split pc/ln table by func
	splitpcln();

	// record src file and line info for each func
	walksymtab(dosrcline);
}

Func*
findfunc(uintptr addr)
{
	Func *f;
	int32 nf, n;

	if(func == nil)
		buildfuncs();
	if(nfunc == 0)
		return nil;
	if(addr &lt; func[0].entry || addr &gt;= func[nfunc].entry)
		return nil;

	// binary search to find func with entry &lt;= addr.
	f = func;
	nf = nfunc;
	while(nf &gt; 0) {
		n = nf/2;
		if(f[n].entry &lt;= addr &amp;&amp; addr &lt; f[n+1].entry)
			return &amp;f[n];
		else if(addr &lt; f[n].entry)
			nf = n;
		else {
			f += n+1;
			nf -= n+1;
		}
	}

	// can&#39;t get here -- we already checked above
	// that the address was in the table bounds.
	// this can only happen if the table isn&#39;t sorted
	// by address or if the binary search above is buggy.
	prints(&#34;findfunc unreachable\n&#34;);
	return nil;
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

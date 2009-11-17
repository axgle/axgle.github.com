<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/godefs/main.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/godefs/main.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Godefs takes as input a host-compilable C file that includes
// standard system headers.  From that input file, it generates
// a standalone (no #includes) C or Go file containing equivalent
// definitions.
//
// The input C file is expected to define new types and enumerated
// constants whose names begin with $ (a legal identifier character
// in gcc).  The output is the standalone definitions of those names,
// with the $ removed.
//
// For example, if this is x.c:
//
//	#include &lt;sys/stat.h&gt;
//
//	typedef struct timespec $Timespec;
//	typedef struct stat $Stat;
//	enum {
//		$S_IFMT = S_IFMT,
//		$S_IFIFO = S_IFIFO,
//		$S_IFCHR = S_IFCHR,
//	};
//
// then &#34;godefs x.c&#34; generates:
//
//	// godefs x.c
//
//	// MACHINE GENERATED - DO NOT EDIT.
//
//	// Constants
//	enum {
//		S_IFMT = 0xf000,
//		S_IFIFO = 0x1000,
//		S_IFCHR = 0x2000,
//	};
//
//	// Types
//	#pragma pack on
//
//	typedef struct Timespec Timespec;
//	struct Timespec {
//		int32 tv_sec;
//		int32 tv_nsec;
//	};
//
//	typedef struct Stat Stat;
//	struct Stat {
//		int32 st_dev;
//		uint32 st_ino;
//		uint16 st_mode;
//		uint16 st_nlink;
//		uint32 st_uid;
//		uint32 st_gid;
//		int32 st_rdev;
//		Timespec st_atimespec;
//		Timespec st_mtimespec;
//		Timespec st_ctimespec;
//		int64 st_size;
//		int64 st_blocks;
//		int32 st_blksize;
//		uint32 st_flags;
//		uint32 st_gen;
//		int32 st_lspare;
//		int64 st_qspare[2];
//	};
//	#pragma pack off
//
// The -g flag to godefs causes it to generate Go output, not C.
// In the Go output, struct fields have leading xx_ prefixes removed
// and the first character capitalized (exported).
//
// Godefs works by invoking gcc to compile the given input file
// and then parses the debug info embedded in the assembly output.
// This is far easier than reading system headers on most machines.
//
// The -c flag sets the compiler (default &#34;gcc&#34;).
//
// The -f flag adds a flag to pass to the compiler (e.g., -f -m64).

#include &#34;a.h&#34;

void
usage(void)
{
	fprint(2, &#34;usage: godefs [-g package] [-c cc] [-f cc-arg] [defs.c ...]\n&#34;);
	exit(1);
}

int gotypefmt(Fmt*);
int ctypefmt(Fmt*);
int prefixlen(Type*);
int cutprefix(char*);

Lang go =
{
	&#34;const (\n&#34;,
	&#34;\t%s = %#llx;\n&#34;,
	&#34;)\n&#34;,

	&#34;type&#34;,

	&#34;type %s struct {\n&#34;,
	&#34;type %s struct {\n&#34;,
	&#34;\tPad%d [%d]byte;\n&#34;,
	&#34;}\n&#34;,

	gotypefmt,
};

Lang c =
{
	&#34;enum {\n&#34;,
	&#34;\t%s = %#llx,\n&#34;,
	&#34;};\n&#34;,

	&#34;typedef&#34;,

	&#34;typedef struct %s %s;\nstruct %s {\n&#34;,
	&#34;typedef union %s %s;\nunion %s {\n&#34;,
	&#34;\tbyte pad%d[%d];\n&#34;,
	&#34;};\n&#34;,

	ctypefmt,
};

char *pkg;

int oargc;
char **oargv;
Lang *lang = &amp;c;

Const *con;
int ncon;

Type **typ;
int ntyp;

void
waitforgcc(void)
{
	waitpid();
}

void
main(int argc, char **argv)
{
	int p[2], pid, i, j, n, off, npad, prefix;
	char **av, *q, *r, *tofree, *name;
	char nambuf[100];
	Biobuf *bin, *bout;
	Type *t;
	Field *f;

	quotefmtinstall();

	oargc = argc;
	oargv = argv;
	av = emalloc((30+argc)*sizeof av[0]);
	atexit(waitforgcc);

	n = 0;
	av[n++] = &#34;gcc&#34;;
	av[n++] = &#34;-c&#34;;
	av[n++] = &#34;-fdollars-in-identifiers&#34;;
	av[n++] = &#34;-S&#34;;	// write assembly
	av[n++] = &#34;-gstabs&#34;;	// include stabs info
	av[n++] = &#34;-o-&#34;;	// to stdout
	av[n++] = &#34;-xc&#34;;	// read C

	ARGBEGIN{
	case &#39;g&#39;:
		lang = &amp;go;
		pkg = EARGF(usage());
		break;
	case &#39;c&#39;:
		av[0] = EARGF(usage());
		break;
	case &#39;f&#39;:
		av[n++] = EARGF(usage());
		break;
	default:
		usage();
	}ARGEND

	if(argc == 0)
		av[n++] = &#34;-&#34;;
	else
		av[n++] = argv[0];
	av[n] = nil;

	// Run gcc writing assembly and stabs debugging to p[1].
	if(pipe(p) &lt; 0)
		sysfatal(&#34;pipe: %r&#34;);

	pid = fork();
	if(pid &lt; 0)
		sysfatal(&#34;fork: %r&#34;);
	if(pid == 0) {
		close(p[0]);
		dup(p[1], 1);
		if(argc == 0) {
			exec(av[0], av);
			fprint(2, &#34;exec gcc: %r\n&#34;);
			exit(1);
		}
		// Some versions of gcc do not accept -S with multiple files.
		// Run gcc once for each file.
		close(0);
		open(&#34;/dev/null&#34;, OREAD);
		for(i=0; i&lt;argc; i++) {
			pid = fork();
			if(pid &lt; 0)
				sysfatal(&#34;fork: %r&#34;);
			if(pid == 0) {
				av[n-1] = argv[i];
				exec(av[0], av);
				fprint(2, &#34;exec gcc: %r\n&#34;);
				exit(1);
			}
			waitpid();
		}
		exit(0);
	}
	close(p[1]);

	// Read assembly, pulling out .stabs lines.
	bin = Bfdopen(p[0], OREAD);
	while((q = Brdstr(bin, &#39;\n&#39;, 1)) != nil) {
		//	.stabs	&#34;float:t(0,12)=r(0,1);4;0;&#34;,128,0,0,0
		tofree = q;
		while(*q == &#39; &#39; || *q == &#39;\t&#39;)
			q++;
		if(strncmp(q, &#34;.stabs&#34;, 6) != 0)
			goto Continue;
		q += 6;
		while(*q == &#39; &#39; || *q == &#39;\t&#39;)
			q++;
		if(*q++ != &#39;\&#34;&#39;) {
		Bad:
			sysfatal(&#34;cannot parse .stabs line:\n%s&#34;, tofree);
		}

		r = strchr(q, &#39;\&#34;&#39;);
		if(r == nil)
			goto Bad;
		*r++ = &#39;\0&#39;;
		if(*r++ != &#39;,&#39;)
			goto Bad;
		if(*r &lt; &#39;0&#39; || *r &gt; &#39;9&#39;)
			goto Bad;
		if(atoi(r) != 128)	// stabs kind = local symbol
			goto Continue;

		parsestabtype(q);

	Continue:
		free(tofree);
	}
	Bterm(bin);
	waitpid();

	// Write defs to standard output.
	bout = Bfdopen(1, OWRITE);
	fmtinstall(&#39;T&#39;, lang-&gt;typefmt);

	// Echo original command line in header.
	Bprint(bout, &#34;//&#34;);
	for(i=0; i&lt;oargc; i++)
		Bprint(bout, &#34; %q&#34;, oargv[i]);
	Bprint(bout, &#34;\n&#34;);
	Bprint(bout, &#34;\n&#34;);
	Bprint(bout, &#34;// MACHINE GENERATED - DO NOT EDIT.\n&#34;);
	Bprint(bout, &#34;\n&#34;);

	if(pkg)
		Bprint(bout, &#34;package %s\n\n&#34;, pkg);

	// Constants.
	Bprint(bout, &#34;// Constants\n&#34;);
	if(ncon &gt; 0) {
		Bprint(bout, lang-&gt;constbegin);
		for(i=0; i&lt;ncon; i++)
			Bprint(bout, lang-&gt;constfmt, con[i].name, con[i].value);
		Bprint(bout, lang-&gt;constend);
	}
	Bprint(bout, &#34;\n&#34;);

	// Types

	// push our names down
	for(i=0; i&lt;ntyp; i++) {
		t = typ[i];
		name = t-&gt;name;
		while(t &amp;&amp; t-&gt;kind == Typedef)
			t = t-&gt;type;
		if(t)
			t-&gt;name = name;
	}

	Bprint(bout, &#34;// Types\n&#34;);

	// Have to turn off structure padding in Plan 9 compiler,
	// mainly because it is more aggressive than gcc tends to be.
	if(lang == &amp;c)
		Bprint(bout, &#34;#pragma pack on\n&#34;);

	for(i=0; i&lt;ntyp; i++) {
		Bprint(bout, &#34;\n&#34;);
		t = typ[i];
		name = t-&gt;name;
		while(t &amp;&amp; t-&gt;kind == Typedef) {
			if(name == nil &amp;&amp; t-&gt;name != nil) {
				name = t-&gt;name;
				if(t-&gt;printed)
					break;
			}
			t = t-&gt;type;
		}
		if(name == nil &amp;&amp; t-&gt;name != nil) {
			name = t-&gt;name;
			if(t-&gt;printed)
				continue;
			t-&gt;printed = 1;
		}
		if(name == nil) {
			fprint(2, &#34;unknown name for %T&#34;, typ[i]);
			continue;
		}
		if(name[0] == &#39;$&#39;)
			name++;
		npad = 0;
		off = 0;
		switch(t-&gt;kind) {
		case 0:
			fprint(2, &#34;unknown type definition for %s\n&#34;, name);
			break;
		default:	// numeric, array, or pointer
		case Array:
		case Ptr:
			Bprint(bout, &#34;%s %lT\n&#34;, lang-&gt;typdef, name, t);
			break;
		case Union:
			// In Go, print union as struct with only first element,
			// padded the rest of the way.
			Bprint(bout, lang-&gt;unionbegin, name, name, name);
			goto StructBody;
		case Struct:
			Bprint(bout, lang-&gt;structbegin, name, name, name);
		StructBody:
			prefix = 0;
			if(lang == &amp;go)
				prefix = prefixlen(t);
			for(j=0; j&lt;t-&gt;nf; j++) {
				f = &amp;t-&gt;f[j];
				// padding
				if(t-&gt;kind == Struct || lang == &amp;go) {
					if(f-&gt;offset%8 != 0 || f-&gt;size%8 != 0) {
						fprint(2, &#34;ignoring bitfield %s.%s\n&#34;, t-&gt;name, f-&gt;name);
						continue;
					}
					if(f-&gt;offset &lt; off)
						sysfatal(&#34;%s: struct fields went backward&#34;, t-&gt;name);
					if(off &lt; f-&gt;offset) {
						Bprint(bout, lang-&gt;structpadfmt, npad++, (f-&gt;offset - off) / 8);
						off = f-&gt;offset;
					}
					off += f-&gt;size;
				}
				name = f-&gt;name;
				if(cutprefix(name))
					name += prefix;
				if(strcmp(name, &#34;&#34;) == 0) {
					snprint(nambuf, sizeof nambuf, &#34;Pad%d&#34;, npad++);
					name = nambuf;
				}
				Bprint(bout, &#34;\t%#lT;\n&#34;, name, f-&gt;type);
				if(t-&gt;kind == Union &amp;&amp; lang == &amp;go)
					break;
			}
			// final padding
			if(t-&gt;kind == Struct || lang == &amp;go) {
				if(off/8 &lt; t-&gt;size)
					Bprint(bout, lang-&gt;structpadfmt, npad++, t-&gt;size - off/8);
			}
			Bprint(bout, lang-&gt;structend);
		}
	}
	if(lang == &amp;c)
		Bprint(bout, &#34;#pragma pack off\n&#34;);
	Bterm(bout);
	exit(0);
}

char *kindnames[] = {
	&#34;void&#34;,	// actually unknown, but byte is good for pointers
	&#34;void&#34;,
	&#34;int8&#34;,
	&#34;uint8&#34;,
	&#34;int16&#34;,
	&#34;uint16&#34;,
	&#34;int32&#34;,
	&#34;uint32&#34;,
	&#34;int64&#34;,
	&#34;uint64&#34;,
	&#34;float32&#34;,
	&#34;float64&#34;,
	&#34;ptr&#34;,
	&#34;struct&#34;,
	&#34;array&#34;,
	&#34;union&#34;,
	&#34;typedef&#34;,
};

int
ctypefmt(Fmt *f)
{
	char *name, *s;
	Type *t;

	name = nil;
	if(f-&gt;flags &amp; FmtLong) {
		name = va_arg(f-&gt;args, char*);
		if(name == nil || name[0] == &#39;\0&#39;)
			name = &#34;_anon_&#34;;
	}
	t = va_arg(f-&gt;args, Type*);
	while(t &amp;&amp; t-&gt;kind == Typedef)
		t = t-&gt;type;
	switch(t-&gt;kind) {
	case Struct:
	case Union:
		// must be named
		s = t-&gt;name;
		if(s == nil) {
			fprint(2, &#34;need name for anonymous struct\n&#34;);
			goto bad;
		}
		else if(s[0] != &#39;$&#39;)
			fprint(2, &#34;need name for struct %s\n&#34;, s);
		else
			s++;
		fmtprint(f, &#34;%s&#34;, s);
		if(name)
			fmtprint(f, &#34; %s&#34;, name);
		break;

	case Array:
		if(name)
			fmtprint(f, &#34;%T %s[%d]&#34;, t-&gt;type, name, t-&gt;size);
		else
			fmtprint(f, &#34;%T[%d]&#34;, t-&gt;type, t-&gt;size);
		break;

	case Ptr:
		if(name)
			fmtprint(f, &#34;%T *%s&#34;, t-&gt;type, name);
		else
			fmtprint(f, &#34;%T*&#34;, t-&gt;type);
		break;

	default:
		fmtprint(f, &#34;%s&#34;, kindnames[t-&gt;kind]);
		if(name)
			fmtprint(f, &#34; %s&#34;, name);
		break;

	bad:
		if(name)
			fmtprint(f, &#34;byte %s[%d]&#34;, name, t-&gt;size);
		else
			fmtprint(f, &#34;byte[%d]&#34;, t-&gt;size);
		break;
	}

	return 0;
}

int
gotypefmt(Fmt *f)
{
	char *name, *s;
	Type *t;

	if(f-&gt;flags &amp; FmtLong) {
		name = va_arg(f-&gt;args, char*);
		if(&#39;a&#39; &lt;= name[0] &amp;&amp; name[0] &lt;= &#39;z&#39;)
			name[0] += &#39;A&#39; - &#39;a&#39;;
		if(name[0] == &#39;_&#39; &amp;&amp; (f-&gt;flags &amp; FmtSharp))
			fmtprint(f, &#34;X&#34;);
		fmtprint(f, &#34;%s &#34;, name);
	}
	t = va_arg(f-&gt;args, Type*);
	while(t &amp;&amp; t-&gt;kind == Typedef)
		t = t-&gt;type;

	switch(t-&gt;kind) {
	case Struct:
	case Union:
		// must be named
		s = t-&gt;name;
		if(s == nil) {
			fprint(2, &#34;need name for anonymous struct\n&#34;);
			fmtprint(f, &#34;STRUCT&#34;);
		}
		else if(s[0] != &#39;$&#39;) {
			fprint(2, &#34;warning: missing name for struct %s\n&#34;, s);
			fmtprint(f, &#34;[%d]byte /* %s */&#34;, t-&gt;size, s);
		} else
			fmtprint(f, &#34;%s&#34;, s+1);
		break;

	case Array:
		fmtprint(f, &#34;[%d]%T&#34;, t-&gt;size, t-&gt;type);
		break;

	case Ptr:
		fmtprint(f, &#34;*%T&#34;, t-&gt;type);
		break;

	default:
		s = kindnames[t-&gt;kind];
		if(strcmp(s, &#34;void&#34;) == 0)
			s = &#34;byte&#34;;
		fmtprint(f, &#34;%s&#34;, s);
	}

	return 0;
}

// Is this the kind of name we should cut a prefix from?
// The rule is that the name cannot begin with underscore
// and must have an underscore eventually.
int
cutprefix(char *name)
{
	char *p;

	// special case: orig_ in register struct
	if(strncmp(name, &#34;orig_&#34;, 5) == 0)
		return 0;

	for(p=name; *p; p++) {
		if(*p == &#39;_&#39;)
			return p-name &gt; 0;
	}
	return 0;
}

// Figure out common struct prefix len
int
prefixlen(Type *t)
{
	int i;
	int len;
	char *p, *name;
	Field *f;

	len = 0;
	name = nil;
	for(i=0; i&lt;t-&gt;nf; i++) {
		f = &amp;t-&gt;f[i];
		if(!cutprefix(f-&gt;name))
			continue;
		p = strchr(f-&gt;name, &#39;_&#39;);
		if(p == nil)
			return 0;
		if(name == nil) {
			name = f-&gt;name;
			len = p+1 - name;
		}
		else if(strncmp(f-&gt;name, name, len) != 0)
			return 0;
	}
	return len;
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

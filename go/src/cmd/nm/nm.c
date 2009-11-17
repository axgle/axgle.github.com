<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/nm/nm.c</title>

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
  <h1 id="generatedHeader">Text file src/cmd/nm/nm.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/nm/nm.c
// http://code.google.com/p/inferno-os/source/browse/utils/nm/nm.c
//
//	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
//	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
//	Portions Copyright © 1997-1999 Vita Nuova Limited
//	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
//	Portions Copyright © 2004,2006 Bruce Ellis
//	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
//	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
//	Portions Copyright © 2009 The Go Authors. All rights reserved.
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

/*
 * nm.c -- drive nm
 */
#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &lt;ar.h&gt;
#include &lt;bio.h&gt;
#include &lt;mach.h&gt;

enum{
	CHUNK	=	256	/* must be power of 2 */
};

char	*errs;			/* exit status */
char	*filename;		/* current file */
char	symname[]=&#34;__.SYMDEF&#34;;	/* table of contents file name */
int	multifile;		/* processing multiple files */
int	aflag;
int	gflag;
int	hflag;
int	nflag;
int	sflag;
int	Sflag;
int	uflag;
int	Tflag;
int	tflag;

Sym	**fnames;		/* file path translation table */
Sym	**symptr;
int	nsym;
Biobuf	bout;

int	cmp(void*, void*);
void	error(char*, ...);
void	execsyms(int);
void	psym(Sym*, void*);
void	printsyms(Sym**, long);
void	doar(Biobuf*);
void	dofile(Biobuf*);
void	zenter(Sym*);

void
usage(void)
{
	fprint(2, &#34;usage: nm [-aghnsTu] file ...\n&#34;);
	exits(&#34;usage&#34;);
}

void
main(int argc, char *argv[])
{
	int i;
	Biobuf	*bin;

	Binit(&amp;bout, 1, OWRITE);
	argv0 = argv[0];
	ARGBEGIN {
	default:	usage();
	case &#39;a&#39;:	aflag = 1; break;
	case &#39;g&#39;:	gflag = 1; break;
	case &#39;h&#39;:	hflag = 1; break;
	case &#39;n&#39;:	nflag = 1; break;
	case &#39;s&#39;:	sflag = 1; break;
	case &#39;S&#39;:	nflag = Sflag = 1; break;
	case &#39;u&#39;:	uflag = 1; break;
	case &#39;t&#39;:	tflag = 1; break;
	case &#39;T&#39;:	Tflag = 1; break;
	} ARGEND
	if (argc == 0)
		usage();
	if (argc &gt; 1)
		multifile++;
	for(i=0; i&lt;argc; i++){
		filename = argv[i];
		bin = Bopen(filename, OREAD);
		if(bin == 0){
			error(&#34;cannot open %s&#34;, filename);
			continue;
		}
		if (isar(bin))
			doar(bin);
		else{
			Bseek(bin, 0, 0);
			dofile(bin);
		}
		Bterm(bin);
	}
	exits(errs);
}

/*
 * read an archive file,
 * processing the symbols for each intermediate file in it.
 */
void
doar(Biobuf *bp)
{
	int offset, size, obj;
	char membername[SARNAME];

	multifile = 1;
	for (offset = Boffset(bp);;offset += size) {
		size = nextar(bp, offset, membername);
		if (size &lt; 0) {
			error(&#34;phase error on ar header %ld&#34;, offset);
			return;
		}
		if (size == 0)
			return;
		if (strcmp(membername, symname) == 0)
			continue;
		obj = objtype(bp, 0);
		if (obj &lt; 0) {
			error(&#34;inconsistent file %s in %s&#34;,
					membername, filename);
			return;
		}
		if (!readar(bp, obj, offset+size, 1)) {
			error(&#34;invalid symbol reference in file %s&#34;,
					membername);
			return;
		}
		filename = membername;
		nsym=0;
		objtraverse(psym, 0);
		printsyms(symptr, nsym);
	}
}

/*
 * process symbols in a file
 */
void
dofile(Biobuf *bp)
{
	int obj;

	obj = objtype(bp, 0);
	if (obj &lt; 0)
		execsyms(Bfildes(bp));
	else
	if (readobj(bp, obj)) {
		nsym = 0;
		objtraverse(psym, 0);
		printsyms(symptr, nsym);
	}
}

/*
 * comparison routine for sorting the symbol table
 *	this screws up on &#39;z&#39; records when aflag == 1
 */
int
cmp(void *vs, void *vt)
{
	Sym **s, **t;

	s = vs;
	t = vt;
	if(nflag)	// sort on address (numeric) order
		if((*s)-&gt;value &lt; (*t)-&gt;value)
			return -1;
		else
			return (*s)-&gt;value &gt; (*t)-&gt;value;
	if(sflag)	// sort on file order (sequence)
		return (*s)-&gt;sequence - (*t)-&gt;sequence;
	return strcmp((*s)-&gt;name, (*t)-&gt;name);
}
/*
 * enter a symbol in the table of filename elements
 */
void
zenter(Sym *s)
{
	static int maxf = 0;

	if (s-&gt;value &gt; maxf) {
		maxf = (s-&gt;value+CHUNK-1) &amp;~ (CHUNK-1);
		fnames = realloc(fnames, (maxf+1)*sizeof(*fnames));
		if(fnames == 0) {
			error(&#34;out of memory&#34;, argv0);
			exits(&#34;memory&#34;);
		}
	}
	fnames[s-&gt;value] = s;
}

/*
 * get the symbol table from an executable file, if it has one
 */
void
execsyms(int fd)
{
	Fhdr f;
	Sym *s;
	int32 n;

	seek(fd, 0, 0);
	if (crackhdr(fd, &amp;f) == 0) {
		error(&#34;Can&#39;t read header for %s&#34;, filename);
		return;
	}
	if (syminit(fd, &amp;f) &lt; 0)
		return;
	s = symbase(&amp;n);
	nsym = 0;
	while(n--)
		psym(s++, 0);

	printsyms(symptr, nsym);
}

void
psym(Sym *s, void* p)
{
	USED(p);
	switch(s-&gt;type) {
	case &#39;T&#39;:
	case &#39;L&#39;:
	case &#39;D&#39;:
	case &#39;B&#39;:
		if (uflag)
			return;
		if (!aflag &amp;&amp; ((s-&gt;name[0] == &#39;.&#39; || s-&gt;name[0] == &#39;$&#39;)))
			return;
		break;
	case &#39;b&#39;:
	case &#39;d&#39;:
	case &#39;l&#39;:
	case &#39;t&#39;:
		if (uflag || gflag)
			return;
		if (!aflag &amp;&amp; ((s-&gt;name[0] == &#39;.&#39; || s-&gt;name[0] == &#39;$&#39;)))
			return;
		break;
	case &#39;U&#39;:
		if (gflag)
			return;
		break;
	case &#39;Z&#39;:
		if (!aflag)
			return;
		break;
	case &#39;m&#39;:
	case &#39;f&#39;:	/* we only see a &#39;z&#39; when the following is true*/
		if(!aflag || uflag || gflag)
			return;
		if (strcmp(s-&gt;name, &#34;.frame&#34;))
			zenter(s);
		break;
	case &#39;a&#39;:
	case &#39;p&#39;:
	case &#39;z&#39;:
	default:
		if(!aflag || uflag || gflag)
			return;
		break;
	}
	symptr = realloc(symptr, (nsym+1)*sizeof(Sym*));
	if (symptr == 0) {
		error(&#34;out of memory&#34;);
		exits(&#34;memory&#34;);
	}
	symptr[nsym++] = s;
}

void
printsyms(Sym **symptr, long nsym)
{
	int i, j, wid;
	Sym *s;
	char *cp;
	char path[512];

	qsort(symptr, nsym, sizeof(*symptr), (void*)cmp);

	wid = 0;
	for (i=0; i&lt;nsym; i++) {
		s = symptr[i];
		if (s-&gt;value &amp;&amp; wid == 0)
			wid = 8;
		else if (s-&gt;value &gt;= 0x100000000LL &amp;&amp; wid == 8)
			wid = 16;
	}
	for (i=0; i&lt;nsym; i++) {
		s = symptr[i];
		if (multifile &amp;&amp; !hflag)
			Bprint(&amp;bout, &#34;%s:&#34;, filename);
		if (s-&gt;type == &#39;z&#39;) {
			fileelem(fnames, (uchar *) s-&gt;name, path, 512);
			cp = path;
		} else
			cp = s-&gt;name;
		if (Tflag)
			Bprint(&amp;bout, &#34;%8ux &#34;, s-&gt;sig);
		if (s-&gt;value || s-&gt;type == &#39;a&#39; || s-&gt;type == &#39;p&#39;)
			Bprint(&amp;bout, &#34;%*llux &#34;, wid, s-&gt;value);
		else
			Bprint(&amp;bout, &#34;%*s &#34;, wid, &#34;&#34;);
		if(Sflag) {
			vlong siz;

			siz = 0;
			for(j=i+1; j&lt;nsym; j++) {
				if(symptr[j]-&gt;type != &#39;a&#39; &amp;&amp; symptr[j]-&gt;type != &#39;p&#39;) {
					siz = symptr[j]-&gt;value - s-&gt;value;
					break;
				}
			}
			if(siz &gt; 0)
				Bprint(&amp;bout, &#34;%*llud &#34;, wid, siz);
		}
		Bprint(&amp;bout, &#34;%c %s&#34;, s-&gt;type, cp);
		if(tflag &amp;&amp; s-&gt;gotype)
			Bprint(&amp;bout, &#34; %*llux&#34;, wid, s-&gt;gotype);
		Bprint(&amp;bout, &#34;\n&#34;);
	}
}

void
error(char *fmt, ...)
{
	Fmt f;
	char buf[128];
	va_list arg;

	fmtfdinit(&amp;f, 2, buf, sizeof buf);
	fmtprint(&amp;f, &#34;%s: &#34;, argv0);
	va_start(arg, fmt);
	fmtvprint(&amp;f, fmt, arg);
	va_end(arg);
	fmtprint(&amp;f, &#34;\n&#34;);
	fmtfdflush(&amp;f);
	errs = &#34;errors&#34;;
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

<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/libmach/obj.c</title>

  <link rel="stylesheet" type="text/css" href="../../doc/style.css">
  <script type="text/javascript" src="../../doc/godocs.js"></script>

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
        <a href="../../index.html"><img src="../../doc/logo-153x55.png" height="55" width="153" alt="Go Home Page" style="border:0" /></a>
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
    <li class="navhead"><a href="../../index.html">Home</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Documents</li>
    <li><a href="../../doc/go_tutorial.html">Tutorial</a></li>
    <li><a href="../../doc/effective_go.html">Effective Go</a></li>
    <li><a href="../../doc/go_faq.html">FAQ</a></li>
    <li><a href="../../doc/go_lang_faq.html">Language Design FAQ</a></li>
    <li><a href="http://www.youtube.com/watch?v=rKnDgT73v8s">Tech talk (1 hour)</a> (<a href="../../doc/go_talk-20091030.pdf">PDF</a>)</li>
    <li><a href="../../doc/go_spec.html">Language Specification</a></li>
    <li><a href="../../doc/go_mem.html">Memory Model</a></li>
    <li><a href="../../doc/go_for_cpp_programmers.html">Go for C++ Programmers</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">How To</li>
    <li><a href="../../doc/install.html">Install Go</a></li>
    <li><a href="../../doc/contribute.html">Contribute code</a></li>

    <li class="blank">&nbsp;</li>
    <li class="navhead">Programming</li>
    <li><a href="../../cmd/index.html">Command documentation</a></li>
    <li><a href="../../pkg/index.html">Package documentation</a></li>
    <li><a href="../index.html">Source files</a></li>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/libmach/obj.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno libmach/obj.c
// http://code.google.com/p/inferno-os/source/browse/utils/libmach/obj.c
//
// 	Copyright © 1994-1999 Lucent Technologies Inc.
// 	Power PC support Copyright © 1995-2004 C H Forsyth (forsyth@terzarima.net).
// 	Portions Copyright © 1997-1999 Vita Nuova Limited.
// 	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com).
// 	Revisions Copyright © 2000-2004 Lucent Technologies Inc. and others.
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

/*
 * obj.c
 * routines universal to all object files
 */
#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &lt;bio.h&gt;
#include &lt;ar.h&gt;
#include &lt;mach.h&gt;
#include &#34;obj.h&#34;

#define islocal(t)	((t)==&#39;a&#39; || (t)==&#39;p&#39;)

enum
{
	NNAMES	= 50,
	MAXIS	= 8,		/* max length to determine if a file is a .? file */
	MAXOFF	= 0x7fffffff,	/* larger than any possible local offset */
	NHASH	= 1024,		/* must be power of two */
	HASHMUL	= 79L,
};

int	_is2(char*),		/* in [$OS].c */
	_is5(char*),
	_is6(char*),
	_is7(char*),
	_is8(char*),
	_is9(char*),
	_isk(char*),
	_isq(char*),
	_isv(char*),
	_isu(char*),
	_read2(Biobuf*, Prog*),
	_read5(Biobuf*, Prog*),
	_read6(Biobuf*, Prog*),
	_read7(Biobuf*, Prog*),
	_read8(Biobuf*, Prog*),
	_read9(Biobuf*, Prog*),
	_readk(Biobuf*, Prog*),
	_readq(Biobuf*, Prog*),
	_readv(Biobuf*, Prog*),
	_readu(Biobuf*, Prog*);

typedef struct Obj	Obj;
typedef struct Symtab	Symtab;

struct	Obj		/* functions to handle each intermediate (.$O) file */
{
	char	*name;				/* name of each $O file */
	int	(*is)(char*);			/* test for each type of $O file */
	int	(*read)(Biobuf*, Prog*);	/* read for each type of $O file*/
};

static Obj	obj[] =
{			/* functions to identify and parse each type of obj */
	[Obj68020]	&#34;68020 .2&#34;,	_is2, _read2,
	[ObjAmd64]	&#34;amd64 .6&#34;,	_is6, _read6,
	[ObjArm]	&#34;arm .5&#34;,	_is5, _read5,
	[ObjAlpha]	&#34;alpha .7&#34;,	_is7, _read7,
	[Obj386]	&#34;386 .8&#34;,	_is8, _read8,
	[ObjSparc]	&#34;sparc .k&#34;,	_isk, _readk,
	[ObjPower]	&#34;power .q&#34;,	_isq, _readq,
	[ObjMips]	&#34;mips .v&#34;,	_isv, _readv,
	[ObjSparc64]	&#34;sparc64 .u&#34;,	_isu, _readu,
	[ObjPower64]	&#34;power64 .9&#34;,	_is9, _read9,
	[Maxobjtype]	0, 0
};

struct	Symtab
{
	struct	Sym 	s;
	struct	Symtab	*next;
};

static	Symtab *hash[NHASH];
static	Sym	*names[NNAMES];	/* working set of active names */

static	int	processprog(Prog*,int);	/* decode each symbol reference */
static	void	objreset(void);
static	void	objlookup(int, char *, int, uint);
static	void 	objupdate(int, int);

static	int	sequence;

int
objtype(Biobuf *bp, char **name)
{
	int i;
	char buf[MAXIS];
	int c;

Retry:
	if(Bread(bp, buf, MAXIS) &lt; MAXIS)
		return -1;
	Bseek(bp, -MAXIS, 1);
	for (i = 0; i &lt; Maxobjtype; i++) {
		if (obj[i].is &amp;&amp; (*obj[i].is)(buf)) {
			if (name)
				*name = obj[i].name;
			return i;
		}
	}

	/*
	 * Maybe there&#39;s an import block we need to skip
	 */
	for(i = 0; i &lt; MAXIS; i++) {
		if(isalpha(buf[i]) || isdigit(buf[i]))
			continue;
		if(i == 0 || buf[i] != &#39;\n&#39;)
			return -1;
		break;
	}

	/*
	 * Found one.  Skip until &#34;\n!\n&#34;
	 */
	while((c = Bgetc(bp)) != Beof) {
		if(c != &#39;\n&#39;)
			continue;
		c = Bgetc(bp);
		if(c != &#39;!&#39;){
			Bungetc(bp);
			continue;
		}
		c = Bgetc(bp);
		if(c != &#39;\n&#39;){
			Bungetc(bp);
			continue;
		}
		goto Retry;
	}
	return -1;
}

int
isar(Biobuf *bp)
{
	int n;
	char magbuf[SARMAG];

	n = Bread(bp, magbuf, SARMAG);
	if(n == SARMAG &amp;&amp; strncmp(magbuf, ARMAG, SARMAG) == 0)
		return 1;
	return 0;
}

/*
 * determine what kind of object file this is and process it.
 * return whether or not this was a recognized intermediate file.
 */
int
readobj(Biobuf *bp, int objtype)
{
	Prog p;

	if (objtype &lt; 0 || objtype &gt;= Maxobjtype || obj[objtype].is == 0)
		return 1;
	objreset();
	while ((*obj[objtype].read)(bp, &amp;p))
		if (!processprog(&amp;p, 1))
			return 0;
	return 1;
}

int
readar(Biobuf *bp, int objtype, vlong end, int doautos)
{
	Prog p;

	if (objtype &lt; 0 || objtype &gt;= Maxobjtype || obj[objtype].is == 0)
		return 1;
	objreset();
	while ((*obj[objtype].read)(bp, &amp;p) &amp;&amp; Boffset(bp) &lt; end)
		if (!processprog(&amp;p, doautos))
			return 0;
	return 1;
}

/*
 *	decode a symbol reference or definition
 */
static	int
processprog(Prog *p, int doautos)
{
	if(p-&gt;kind == aNone)
		return 1;
	if(p-&gt;sym &lt; 0 || p-&gt;sym &gt;= NNAMES)
		return 0;
	switch(p-&gt;kind)
	{
	case aName:
		if (!doautos)
		if(p-&gt;type != &#39;U&#39; &amp;&amp; p-&gt;type != &#39;b&#39;)
			break;
		objlookup(p-&gt;sym, p-&gt;id, p-&gt;type, p-&gt;sig);
		break;
	case aText:
		objupdate(p-&gt;sym, &#39;T&#39;);
		break;
	case aData:
		objupdate(p-&gt;sym, &#39;D&#39;);
		break;
	default:
		break;
	}
	return 1;
}

/*
 * find the entry for s in the symbol array.
 * make a new entry if it is not already there.
 */
static void
objlookup(int id, char *name, int type, uint sig)
{
	int32 h;
	char *cp;
	Sym *s;
	Symtab *sp;

	s = names[id];
	if(s &amp;&amp; strcmp(s-&gt;name, name) == 0) {
		s-&gt;type = type;
		s-&gt;sig = sig;
		return;
	}

	h = *name;
	for(cp = name+1; *cp; h += *cp++)
		h *= HASHMUL;
	if(h &lt; 0)
		h = ~h;
	h &amp;= (NHASH-1);
	if (type == &#39;U&#39; || type == &#39;b&#39; || islocal(type)) {
		for(sp = hash[h]; sp; sp = sp-&gt;next)
			if(strcmp(sp-&gt;s.name, name) == 0) {
				switch(sp-&gt;s.type) {
				case &#39;T&#39;:
				case &#39;D&#39;:
				case &#39;U&#39;:
					if (type == &#39;U&#39;) {
						names[id] = &amp;sp-&gt;s;
						return;
					}
					break;
				case &#39;t&#39;:
				case &#39;d&#39;:
				case &#39;b&#39;:
					if (type == &#39;b&#39;) {
						names[id] = &amp;sp-&gt;s;
						return;
					}
					break;
				case &#39;a&#39;:
				case &#39;p&#39;:
					if (islocal(type)) {
						names[id] = &amp;sp-&gt;s;
						return;
					}
					break;
				default:
					break;
				}
			}
	}
	sp = malloc(sizeof(Symtab));
	sp-&gt;s.name = name;
	sp-&gt;s.type = type;
	sp-&gt;s.sig = sig;
	sp-&gt;s.value = islocal(type) ? MAXOFF : 0;
	sp-&gt;s.sequence = sequence++;
	names[id] = &amp;sp-&gt;s;
	sp-&gt;next = hash[h];
	hash[h] = sp;
	return;
}
/*
 *	traverse the symbol lists
 */
void
objtraverse(void (*fn)(Sym*, void*), void *pointer)
{
	int i;
	Symtab *s;

	for(i = 0; i &lt; NHASH; i++)
		for(s = hash[i]; s; s = s-&gt;next)
			(*fn)(&amp;s-&gt;s, pointer);
}

/*
 * update the offset information for a &#39;a&#39; or &#39;p&#39; symbol in an intermediate file
 */
void
_offset(int id, vlong off)
{
	Sym *s;

	s = names[id];
	if (s &amp;&amp; s-&gt;name[0] &amp;&amp; islocal(s-&gt;type) &amp;&amp; s-&gt;value &gt; off)
		s-&gt;value = off;
}

/*
 * update the type of a global text or data symbol
 */
static void
objupdate(int id, int type)
{
	Sym *s;

	s = names[id];
	if (s &amp;&amp; s-&gt;name[0])
		if (s-&gt;type == &#39;U&#39;)
			s-&gt;type = type;
		else if (s-&gt;type == &#39;b&#39;)
			s-&gt;type = tolower(type);
}

/*
 * look for the next file in an archive
 */
int
nextar(Biobuf *bp, int offset, char *buf)
{
	struct ar_hdr a;
	int i, r;
	int32 arsize;

	if (offset&amp;01)
		offset++;
	Bseek(bp, offset, 0);
	r = Bread(bp, &amp;a, SAR_HDR);
	if(r != SAR_HDR)
		return 0;
	if(strncmp(a.fmag, ARFMAG, sizeof(a.fmag)))
		return -1;
	for(i=0; i&lt;sizeof(a.name) &amp;&amp; i&lt;SARNAME &amp;&amp; a.name[i] != &#39; &#39;; i++)
		buf[i] = a.name[i];
	buf[i] = 0;
	arsize = strtol(a.size, 0, 0);
	if (arsize&amp;1)
		arsize++;
	return arsize + SAR_HDR;
}

static void
objreset(void)
{
	int i;
	Symtab *s, *n;

	for(i = 0; i &lt; NHASH; i++) {
		for(s = hash[i]; s; s = n) {
			n = s-&gt;next;
			free(s-&gt;s.name);
			free(s);
		}
		hash[i] = 0;
	}
	memset(names, 0, sizeof names);
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

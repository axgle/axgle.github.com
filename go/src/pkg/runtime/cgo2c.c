<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/cgo2c.c</title>

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
	<li>Thu Nov 12 16:00:43 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/cgo2c.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* Translate a .cgo file into a .c file.  A .cgo file is a combination
   of a limited form of Go with C.  */

/*
   package PACKAGENAME
   {# line}
   func NAME([NAME TYPE { , NAME TYPE }]) [(NAME TYPE { , NAME TYPE })] \{
     C code with proper brace nesting
   \}
*/

/* We generate C code which implements the function such that it can
   be called from Go and executes the C code.  */

#include &lt;assert.h&gt;
#include &lt;ctype.h&gt;
#include &lt;stdio.h&gt;
#include &lt;stdlib.h&gt;
#include &lt;string.h&gt;
#include &lt;errno.h&gt;

/* Whether we&#39;re emitting for gcc */
static int gcc;

/* File and line number */
static const char *file;
static unsigned int lineno;

/* List of names and types.  */
struct params {
	struct params *next;
	char *name;
	char *type;
};

/* index into type_table */
enum {
	Bool,
	Float,
	Int,
	Uint,
	Uintptr,
	String,
	Slice,
};

static struct {
	char *name;
	int size;
} type_table[] = {
	/* variable sized first, for easy replacement */
	/* order matches enum above */
	/* default is 32-bit architecture sizes */
	&#34;bool&#34;,		1,
	&#34;float&#34;,	4,
	&#34;int&#34;,		4,
	&#34;uint&#34;,		4,
	&#34;uintptr&#34;,	4,
	&#34;String&#34;,	8,
	&#34;Slice&#34;,	12,

	/* fixed size */
	&#34;float32&#34;,	4,
	&#34;float64&#34;,	8,
	&#34;byte&#34;,		1,
	&#34;int8&#34;,		1,
	&#34;uint8&#34;,	1,
	&#34;int16&#34;,	2,
	&#34;uint16&#34;,	2,
	&#34;int32&#34;,	4,
	&#34;uint32&#34;,	4,
	&#34;int64&#34;,	8,
	&#34;uint64&#34;,	8,

	NULL,
};

/* Fixed structure alignment (non-gcc only) */
int structround = 4;

/* Unexpected EOF.  */
static void
bad_eof(void)
{
	fprintf(stderr, &#34;%s:%u: unexpected EOF\n&#34;, file, lineno);
	exit(1);
}

/* Out of memory.  */
static void
bad_mem(void)
{
	fprintf(stderr, &#34;%s:%u: out of memory\n&#34;, file, lineno);
	exit(1);
}

/* Allocate memory without fail.  */
static void *
xmalloc(unsigned int size)
{
	void *ret = malloc(size);
	if (ret == NULL)
		bad_mem();
	return ret;
}

/* Reallocate memory without fail.  */
static void*
xrealloc(void *buf, unsigned int size)
{
	void *ret = realloc(buf, size);
	if (ret == NULL)
		bad_mem();
	return ret;
}

/* Free a list of parameters.  */
static void
free_params(struct params *p)
{
	while (p != NULL) {
		struct params *next;

		next = p-&gt;next;
		free(p-&gt;name);
		free(p-&gt;type);
		free(p);
		p = next;
	}
}

/* Read a character, tracking lineno.  */
static int
getchar_update_lineno(void)
{
	int c;

	c = getchar();
	if (c == &#39;\n&#39;)
		++lineno;
	return c;
}

/* Read a character, giving an error on EOF, tracking lineno.  */
static int
getchar_no_eof(void)
{
	int c;

	c = getchar_update_lineno();
	if (c == EOF)
		bad_eof();
	return c;
}

/* Read a character, skipping comments.  */
static int
getchar_skipping_comments(void)
{
	int c;

	while (1) {
		c = getchar_update_lineno();
		if (c != &#39;/&#39;)
			return c;

		c = getchar();
		if (c == &#39;/&#39;) {
			do {
				c = getchar_update_lineno();
			} while (c != EOF &amp;&amp; c != &#39;\n&#39;);
			return c;
		} else if (c == &#39;*&#39;) {
			while (1) {
				c = getchar_update_lineno();
				if (c == EOF)
					return EOF;
				if (c == &#39;*&#39;) {
					do {
						c = getchar_update_lineno();
					} while (c == &#39;*&#39;);
					if (c == &#39;/&#39;)
						break;
				}
			}
		} else {
			ungetc(c, stdin);
			return &#39;/&#39;;
		}
	}
}

/* Read and return a token.  Tokens are delimited by whitespace or by
   [(),{}].  The latter are all returned as single characters.  */
static char *
read_token(void)
{
	int c;
	char *buf;
	unsigned int alc, off;
	const char* delims = &#34;(),{}&#34;;

	while (1) {
		c = getchar_skipping_comments();
		if (c == EOF)
			return NULL;
		if (!isspace(c))
			break;
	}
	alc = 16;
	buf = xmalloc(alc + 1);
	off = 0;
	if (strchr(delims, c) != NULL) {
		buf[off] = c;
		++off;
	} else {
		while (1) {
			if (off &gt;= alc) {
				alc *= 2;
				buf = xrealloc(buf, alc + 1);
			}
			buf[off] = c;
			++off;
			c = getchar_skipping_comments();
			if (c == EOF)
				break;
			if (isspace(c) || strchr(delims, c) != NULL) {
				if (c == &#39;\n&#39;)
					lineno--;
				ungetc(c, stdin);
				break;
			}
		}
	}
	buf[off] = &#39;\0&#39;;
	return buf;
}

/* Read a token, giving an error on EOF.  */
static char *
read_token_no_eof(void)
{
	char *token = read_token();
	if (token == NULL)
		bad_eof();
	return token;
}

/* Read the package clause, and return the package name.  */
static char *
read_package(void)
{
	char *token;

	token = read_token_no_eof();
	if (strcmp(token, &#34;package&#34;) != 0) {
		fprintf(stderr,
			&#34;%s:%u: expected \&#34;package\&#34;, got \&#34;%s\&#34;\n&#34;,
			file, lineno, token);
		exit(1);
	}
	return read_token_no_eof();
}

/* Read and copy preprocessor lines.  */
static void
read_preprocessor_lines(void)
{
	while (1) {
		int c;

		do {
			c = getchar_skipping_comments();
		} while (isspace(c));
		if (c != &#39;#&#39;) {
			ungetc(c, stdin);
			break;
		}
		putchar(c);
		do {
			c = getchar_update_lineno();
			putchar(c);
		} while (c != &#39;\n&#39;);
	}
}

/* Read a type in Go syntax and return a type in C syntax.  We only
   permit basic types and pointers.  */
static char *
read_type(void)
{
	char *p, *op, *q;
	int pointer_count;
	unsigned int len;

	p = read_token_no_eof();
	if (*p != &#39;*&#39;)
		return p;
	op = p;
	pointer_count = 0;
	while (*p == &#39;*&#39;) {
		++pointer_count;
		++p;
	}
	len = strlen(p);
	q = xmalloc(len + pointer_count + 1);
	memcpy(q, p, len);
	while (pointer_count &gt; 0) {
		q[len] = &#39;*&#39;;
		++len;
		--pointer_count;
	}
	q[len] = &#39;\0&#39;;
	free(op);
	return q;
}

/* Return the size of the given type. */
static int
type_size(char *p)
{
	int i;

	if(p[strlen(p)-1] == &#39;*&#39;)
		return type_table[Uintptr].size;

	for(i=0; type_table[i].name; i++)
		if(strcmp(type_table[i].name, p) == 0)
			return type_table[i].size;
	fprintf(stderr, &#34;%s:%u: unknown type %s\n&#34;, file, lineno, p);
	exit(1);
	return 0;
}

/* Read a list of parameters.  Each parameter is a name and a type.
   The list ends with a &#39;)&#39;.  We have already read the &#39;(&#39;.  */
static struct params *
read_params(int *poffset)
{
	char *token;
	struct params *ret, **pp, *p;
	int offset, size, rnd;

	ret = NULL;
	pp = &amp;ret;
	token = read_token_no_eof();
	offset = 0;
	if (strcmp(token, &#34;)&#34;) != 0) {
		while (1) {
			p = xmalloc(sizeof(struct params));
			p-&gt;name = token;
			p-&gt;type = read_type();
			p-&gt;next = NULL;
			*pp = p;
			pp = &amp;p-&gt;next;

			size = type_size(p-&gt;type);
			rnd = size;
			if(rnd &gt; structround)
				rnd = structround;
			if(offset%rnd)
				offset += rnd - offset%rnd;
			offset += size;

			token = read_token_no_eof();
			if (strcmp(token, &#34;,&#34;) != 0)
				break;
			token = read_token_no_eof();
		}
	}
	if (strcmp(token, &#34;)&#34;) != 0) {
		fprintf(stderr, &#34;%s:%u: expected &#39;(&#39;\n&#34;,
			file, lineno);
		exit(1);
	}
	if (poffset != NULL)
		*poffset = offset;
	return ret;
}

/* Read a function header.  This reads up to and including the initial
   &#39;{&#39; character.  Returns 1 if it read a header, 0 at EOF.  */
static int
read_func_header(char **name, struct params **params, int *paramwid, struct params **rets)
{
	int lastline;
	char *token;

	lastline = -1;
	while (1) {
		token = read_token();
		if (token == NULL)
			return 0;
		if (strcmp(token, &#34;func&#34;) == 0) {
			if(lastline != -1)
				printf(&#34;\n&#34;);
			break;
		}
		if (lastline != lineno) {
			if (lastline == lineno-1)
				printf(&#34;\n&#34;);
			else
				printf(&#34;\n#line %d \&#34;%s\&#34;\n&#34;, lineno, file);
			lastline = lineno;
		}
		printf(&#34;%s &#34;, token);
	}

	*name = read_token_no_eof();

	token = read_token();
	if (token == NULL || strcmp(token, &#34;(&#34;) != 0) {
		fprintf(stderr, &#34;%s:%u: expected \&#34;(\&#34;\n&#34;,
			file, lineno);
		exit(1);
	}
	*params = read_params(paramwid);

	token = read_token();
	if (token == NULL || strcmp(token, &#34;(&#34;) != 0)
		*rets = NULL;
	else {
		*rets = read_params(NULL);
		token = read_token();
	}
	if (token == NULL || strcmp(token, &#34;{&#34;) != 0) {
		fprintf(stderr, &#34;%s:%u: expected \&#34;{\&#34;\n&#34;,
			file, lineno);
		exit(1);
	}
	return 1;
}

/* Write out parameters.  */
static void
write_params(struct params *params, int *first)
{
	struct params *p;

	for (p = params; p != NULL; p = p-&gt;next) {
		if (*first)
			*first = 0;
		else
			printf(&#34;, &#34;);
		printf(&#34;%s %s&#34;, p-&gt;type, p-&gt;name);
	}
}

/* Write a 6g function header.  */
static void
write_6g_func_header(char *package, char *name, struct params *params,
		     int paramwid, struct params *rets)
{
	int first, n;

	printf(&#34;void\n%s·%s(&#34;, package, name);
	first = 1;
	write_params(params, &amp;first);

	/* insert padding to align output struct */
	if(rets != NULL &amp;&amp; paramwid%structround != 0) {
		n = structround - paramwid%structround;
		if(n &amp; 1)
			printf(&#34;, uint8&#34;);
		if(n &amp; 2)
			printf(&#34;, uint16&#34;);
		if(n &amp; 4)
			printf(&#34;, uint32&#34;);
	}

	write_params(rets, &amp;first);
	printf(&#34;)\n{\n&#34;);
}

/* Write a 6g function trailer.  */
static void
write_6g_func_trailer(struct params *rets)
{
	struct params *p;

	for (p = rets; p != NULL; p = p-&gt;next)
		printf(&#34;\tFLUSH(&amp;%s);\n&#34;, p-&gt;name);
	printf(&#34;}\n&#34;);
}

/* Define the gcc function return type if necessary.  */
static void
define_gcc_return_type(char *package, char *name, struct params *rets)
{
	struct params *p;

	if (rets == NULL || rets-&gt;next == NULL)
		return;
	printf(&#34;struct %s_%s_ret {\n&#34;, package, name);
	for (p = rets; p != NULL; p = p-&gt;next)
		printf(&#34;  %s %s;\n&#34;, p-&gt;type, p-&gt;name);
	printf(&#34;};\n&#34;);
}

/* Write out the gcc function return type.  */
static void
write_gcc_return_type(char *package, char *name, struct params *rets)
{
	if (rets == NULL)
		printf(&#34;void&#34;);
	else if (rets-&gt;next == NULL)
		printf(&#34;%s&#34;, rets-&gt;type);
	else
		printf(&#34;struct %s_%s_ret&#34;, package, name);
}

/* Write out a gcc function header.  */
static void
write_gcc_func_header(char *package, char *name, struct params *params,
		      struct params *rets)
{
	int first;
	struct params *p;

	define_gcc_return_type(package, name, rets);
	write_gcc_return_type(package, name, rets);
	printf(&#34; %s_%s(&#34;, package, name);
	first = 1;
	write_params(params, &amp;first);
	printf(&#34;) asm (\&#34;%s.%s\&#34;);\n&#34;, package, name);
	write_gcc_return_type(package, name, rets);
	printf(&#34; %s_%s(&#34;, package, name);
	first = 1;
	write_params(params, &amp;first);
	printf(&#34;)\n{\n&#34;);
	for (p = rets; p != NULL; p = p-&gt;next)
		printf(&#34;  %s %s;\n&#34;, p-&gt;type, p-&gt;name);
}

/* Write out a gcc function trailer.  */
static void
write_gcc_func_trailer(char *package, char *name, struct params *rets)
{
	if (rets == NULL)
		;
	else if (rets-&gt;next == NULL)
		printf(&#34;return %s;\n&#34;, rets-&gt;name);
	else {
		struct params *p;

		printf(&#34;  {\n    struct %s_%s_ret __ret;\n&#34;, package, name);
		for (p = rets; p != NULL; p = p-&gt;next)
			printf(&#34;    __ret.%s = %s;\n&#34;, p-&gt;name, p-&gt;name);
		printf(&#34;    return __ret;\n  }\n&#34;);
	}
	printf(&#34;}\n&#34;);
}

/* Write out a function header.  */
static void
write_func_header(char *package, char *name,
		  struct params *params, int paramwid,
		  struct params *rets)
{
	if (gcc)
		write_gcc_func_header(package, name, params, rets);
	else
		write_6g_func_header(package, name, params, paramwid, rets);
	printf(&#34;#line %d \&#34;%s\&#34;\n&#34;, lineno, file);
}

/* Write out a function trailer.  */
static void
write_func_trailer(char *package, char *name,
		   struct params *rets)
{
	if (gcc)
		write_gcc_func_trailer(package, name, rets);
	else
		write_6g_func_trailer(rets);
}

/* Read and write the body of the function, ending in an unnested }
   (which is read but not written).  */
static void
copy_body(void)
{
	int nesting = 0;
	while (1) {
		int c;

		c = getchar_no_eof();
		if (c == &#39;}&#39; &amp;&amp; nesting == 0)
			return;
		putchar(c);
		switch (c) {
		default:
			break;
		case &#39;{&#39;:
			++nesting;
			break;
		case &#39;}&#39;:
			--nesting;
			break;
		case &#39;/&#39;:
			c = getchar_update_lineno();
			putchar(c);
			if (c == &#39;/&#39;) {
				do {
					c = getchar_no_eof();
					putchar(c);
				} while (c != &#39;\n&#39;);
			} else if (c == &#39;*&#39;) {
				while (1) {
					c = getchar_no_eof();
					putchar(c);
					if (c == &#39;*&#39;) {
						do {
							c = getchar_no_eof();
							putchar(c);
						} while (c == &#39;*&#39;);
						if (c == &#39;/&#39;)
							break;
					}
				}
			}
			break;
		case &#39;&#34;&#39;:
		case &#39;\&#39;&#39;:
			{
				int delim = c;
				do {
					c = getchar_no_eof();
					putchar(c);
					if (c == &#39;\\&#39;) {
						c = getchar_no_eof();
						putchar(c);
						c = &#39;\0&#39;;
					}
				} while (c != delim);
			}
			break;
		}
	}
}

/* Process the entire file.  */
static void
process_file(void)
{
	char *package, *name;
	struct params *params, *rets;
	int paramwid;

	package = read_package();
	read_preprocessor_lines();
	while (read_func_header(&amp;name, &amp;params, &amp;paramwid, &amp;rets)) {
		write_func_header(package, name, params, paramwid, rets);
		copy_body();
		write_func_trailer(package, name, rets);
		free(name);
		free_params(params);
		free_params(rets);
	}
	free(package);
}

static void
usage(void)
{
	fprintf(stderr, &#34;Usage: cgo2c [--6g | --gc] [file]\n&#34;);
	exit(1);
}

int
main(int argc, char **argv)
{
	char *goarch;

	while(argc &gt; 1 &amp;&amp; argv[1][0] == &#39;-&#39;) {
		if(strcmp(argv[1], &#34;-&#34;) == 0)
			break;
		if(strcmp(argv[1], &#34;--6g&#34;) == 0)
			gcc = 0;
		else if(strcmp(argv[1], &#34;--gcc&#34;) == 0)
			gcc = 1;
		else
			usage();
		argc--;
		argv++;
	}

	if(argc &lt;= 1 || strcmp(argv[1], &#34;-&#34;) == 0) {
		file = &#34;&lt;stdin&gt;&#34;;
		process_file();
		return 0;
	}

	if(argc &gt; 2)
		usage();

	file = argv[1];
	if(freopen(file, &#34;r&#34;, stdin) == 0) {
		fprintf(stderr, &#34;open %s: %s\n&#34;, file, strerror(errno));
		exit(1);
	}

	if(!gcc) {
		// 6g etc; update size table
		goarch = getenv(&#34;GOARCH&#34;);
		if(goarch != NULL &amp;&amp; strcmp(goarch, &#34;amd64&#34;) == 0) {
			type_table[Uintptr].size = 8;
			type_table[String].size = 16;
			type_table[Slice].size = 8+4+4;
			structround = 8;
		}
	}

	process_file();
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

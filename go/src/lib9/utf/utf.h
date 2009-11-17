<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/lib9/utf/utf.h</title>

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
	<li>Thu Nov 12 15:51:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/lib9/utf/utf.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
/*
 * The authors of this software are Rob Pike and Ken Thompson.
 *              Copyright (c) 1998-2002 by Lucent Technologies.
 *              Portions Copyright (c) 2009 The Go Authors.  All rights reserved.
 * Permission to use, copy, modify, and distribute this software for any
 * purpose without fee is hereby granted, provided that this entire notice
 * is included in all copies of any software which is or includes a copy
 * or modification of this software and in all copies of the supporting
 * documentation for such software.
 * THIS SOFTWARE IS BEING PROVIDED &#34;AS IS&#34;, WITHOUT ANY EXPRESS OR IMPLIED
 * WARRANTY.  IN PARTICULAR, NEITHER THE AUTHORS NOR LUCENT TECHNOLOGIES MAKE ANY
 * REPRESENTATION OR WARRANTY OF ANY KIND CONCERNING THE MERCHANTABILITY
 * OF THIS SOFTWARE OR ITS FITNESS FOR ANY PARTICULAR PURPOSE.
 */

#ifndef _UTFH_
#define _UTFH_ 1

#include &lt;stdint.h&gt;

typedef unsigned int Rune;	/* Code-point values in Unicode 4.0 are 21 bits wide.*/

enum
{
  UTFmax	= 4,		/* maximum bytes per rune */
  Runesync	= 0x80,		/* cannot represent part of a UTF sequence (&lt;) */
  Runeself	= 0x80,		/* rune and UTF sequences are the same (&lt;) */
  Runeerror	= 0xFFFD,	/* decoding error in UTF */
  Runemax	= 0x10FFFF,	/* maximum rune value */
};

#ifdef	__cplusplus
extern &#34;C&#34; {
#endif

/*
 * rune routines
 */

/*
 * These routines were written by Rob Pike and Ken Thompson
 * and first appeared in Plan 9.
 * SEE ALSO
 * utf (7)
 * tcs (1)
*/

// runetochar copies (encodes) one rune, pointed to by r, to at most
// UTFmax bytes starting at s and returns the number of bytes generated.

int runetochar(char* s, const Rune* r);


// chartorune copies (decodes) at most UTFmax bytes starting at s to
// one rune, pointed to by r, and returns the number of bytes consumed.
// If the input is not exactly in UTF format, chartorune will set *r
// to Runeerror and return 1.
//
// Note: There is no special case for a &#34;null-terminated&#34; string. A
// string whose first byte has the value 0 is the UTF8 encoding of the
// Unicode value 0 (i.e., ASCII NULL). A byte value of 0 is illegal
// anywhere else in a UTF sequence.

int chartorune(Rune* r, const char* s);


// charntorune is like chartorune, except that it will access at most
// n bytes of s.  If the UTF sequence is incomplete within n bytes,
// charntorune will set *r to Runeerror and return 0. If it is complete
// but not in UTF format, it will set *r to Runeerror and return 1.
//
// Added 2004-09-24 by Wei-Hwa Huang

int charntorune(Rune* r, const char* s, int n);

// isvalidcharntorune(str, n, r, consumed)
// is a convenience function that calls &#34;*consumed = charntorune(r, str, n)&#34;
// and returns an int (logically boolean) indicating whether the first
// n bytes of str was a valid and complete UTF sequence.

int isvalidcharntorune(const char* str, int n, Rune* r, int* consumed);

// runelen returns the number of bytes required to convert r into UTF.

int runelen(Rune r);


// runenlen returns the number of bytes required to convert the n
// runes pointed to by r into UTF.

int runenlen(const Rune* r, int n);


// fullrune returns 1 if the string s of length n is long enough to be
// decoded by chartorune, and 0 otherwise. This does not guarantee
// that the string contains a legal UTF encoding. This routine is used
// by programs that obtain input one byte at a time and need to know
// when a full rune has arrived.

int fullrune(const char* s, int n);

// The following routines are analogous to the corresponding string
// routines with &#34;utf&#34; substituted for &#34;str&#34;, and &#34;rune&#34; substituted
// for &#34;chr&#34;.

// utflen returns the number of runes that are represented by the UTF
// string s. (cf. strlen)

int utflen(const char* s);


// utfnlen returns the number of complete runes that are represented
// by the first n bytes of the UTF string s. If the last few bytes of
// the string contain an incompletely coded rune, utfnlen will not
// count them; in this way, it differs from utflen, which includes
// every byte of the string. (cf. strnlen)

int utfnlen(const char* s, long n);


// utfrune returns a pointer to the first occurrence of rune r in the
// UTF string s, or 0 if r does not occur in the string.  The NULL
// byte terminating a string is considered to be part of the string s.
// (cf. strchr)

/*const*/ char* utfrune(const char* s, Rune r);


// utfrrune returns a pointer to the last occurrence of rune r in the
// UTF string s, or 0 if r does not occur in the string.  The NULL
// byte terminating a string is considered to be part of the string s.
// (cf. strrchr)

/*const*/ char* utfrrune(const char* s, Rune r);


// utfutf returns a pointer to the first occurrence of the UTF string
// s2 as a UTF substring of s1, or 0 if there is none. If s2 is the
// null string, utfutf returns s1. (cf. strstr)

const char* utfutf(const char* s1, const char* s2);


// utfecpy copies UTF sequences until a null sequence has been copied,
// but writes no sequences beyond es1.  If any sequences are copied,
// s1 is terminated by a null sequence, and a pointer to that sequence
// is returned.  Otherwise, the original s1 is returned. (cf. strecpy)

char* utfecpy(char *s1, char *es1, const char *s2);



// These functions are rune-string analogues of the corresponding
// functions in strcat (3).
//
// These routines first appeared in Plan 9.
// SEE ALSO
// memmove (3)
// rune (3)
// strcat (2)
//
// BUGS: The outcome of overlapping moves varies among implementations.

Rune* runestrcat(Rune* s1, const Rune* s2);
Rune* runestrncat(Rune* s1, const Rune* s2, long n);

const Rune* runestrchr(const Rune* s, Rune c);

int runestrcmp(const Rune* s1, const Rune* s2);
int runestrncmp(const Rune* s1, const Rune* s2, long n);

Rune* runestrcpy(Rune* s1, const Rune* s2);
Rune* runestrncpy(Rune* s1, const Rune* s2, long n);
Rune* runestrecpy(Rune* s1, Rune* es1, const Rune* s2);

Rune* runestrdup(const Rune* s);

const Rune* runestrrchr(const Rune* s, Rune c);
long runestrlen(const Rune* s);
const Rune* runestrstr(const Rune* s1, const Rune* s2);



// The following routines test types and modify cases for Unicode
// characters.  Unicode defines some characters as letters and
// specifies three cases: upper, lower, and title.  Mappings among the
// cases are also defined, although they are not exhaustive: some
// upper case letters have no lower case mapping, and so on.  Unicode
// also defines several character properties, a subset of which are
// checked by these routines.  These routines are based on Unicode
// version 3.0.0.
//
// NOTE: The routines are implemented in C, so the boolean functions
// (e.g., isupperrune) return 0 for false and 1 for true.
//
//
// toupperrune, tolowerrune, and totitlerune are the Unicode case
// mappings. These routines return the character unchanged if it has
// no defined mapping.

Rune toupperrune(Rune r);
Rune tolowerrune(Rune r);
Rune totitlerune(Rune r);


// isupperrune tests for upper case characters, including Unicode
// upper case letters and targets of the toupper mapping. islowerrune
// and istitlerune are defined analogously.

int isupperrune(Rune r);
int islowerrune(Rune r);
int istitlerune(Rune r);


// isalpharune tests for Unicode letters; this includes ideographs in
// addition to alphabetic characters.

int isalpharune(Rune r);


// isdigitrune tests for digits. Non-digit numbers, such as Roman
// numerals, are not included.

int isdigitrune(Rune r);


// isideographicrune tests for ideographic characters and numbers, as
// defined by the Unicode standard.

int isideographicrune(Rune r);


// isspacerune tests for whitespace characters, including &#34;C&#34; locale
// whitespace, Unicode defined whitespace, and the &#34;zero-width
// non-break space&#34; character.

int isspacerune(Rune r);


// (The comments in this file were copied from the manpage files rune.3,
// isalpharune.3, and runestrcat.3. Some formatting changes were also made
// to conform to Google style. /JRM 11/11/05)

#ifdef	__cplusplus
}
#endif

#endif
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

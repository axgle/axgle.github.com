<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/5l/optab.c</title>

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
	<li>Thu Nov 12 15:49:51 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/5l/optab.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Inferno utils/5l/optab.c
// http://code.google.com/p/inferno-os/source/browse/utils/5l/optab.c
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

Optab	optab[] =
{
	/* struct Optab:
	  OPCODE,	from, prog-&gt;reg, to,		 type,size,param,flag */
	{ ATEXT,	C_LEXT,	C_NONE,	C_LCON, 	 0, 0, 0 },
	{ ATEXT,	C_LEXT,	C_REG,	C_LCON, 	 0, 0, 0 },
	{ ATEXT,	C_ADDR,	C_NONE,	C_LCON, 	 0, 0, 0 },
	{ ATEXT,	C_ADDR,	C_REG,	C_LCON, 	 0, 0, 0 },

	{ AADD,		C_REG,	C_REG,	C_REG,		 1, 4, 0 },
	{ AADD,		C_REG,	C_NONE,	C_REG,		 1, 4, 0 },
	{ AMOVW,	C_REG,	C_NONE,	C_REG,		 1, 4, 0 },
	{ AMVN,		C_REG,	C_NONE,	C_REG,		 1, 4, 0 },
	{ ACMP,		C_REG,	C_REG,	C_NONE,		 1, 4, 0 },

	{ AADD,		C_RCON,	C_REG,	C_REG,		 2, 4, 0 },
	{ AADD,		C_RCON,	C_NONE,	C_REG,		 2, 4, 0 },
	{ AMOVW,	C_RCON,	C_NONE,	C_REG,		 2, 4, 0 },
	{ AMVN,		C_RCON,	C_NONE,	C_REG,		 2, 4, 0 },
	{ ACMP,		C_RCON,	C_REG,	C_NONE,		 2, 4, 0 },

	{ AADD,		C_SHIFT,C_REG,	C_REG,		 3, 4, 0 },
	{ AADD,		C_SHIFT,C_NONE,	C_REG,		 3, 4, 0 },
	{ AMVN,		C_SHIFT,C_NONE,	C_REG,		 3, 4, 0 },
	{ ACMP,		C_SHIFT,C_REG,	C_NONE,		 3, 4, 0 },

	{ AMOVW,	C_RECON,C_NONE,	C_REG,		 4, 4, REGSB },
	{ AMOVW,	C_RACON,C_NONE,	C_REG,		 4, 4, REGSP },

	{ AB,		C_NONE,	C_NONE,	C_SBRA,		 5, 4, 0,	LPOOL },
	{ ABL,		C_NONE,	C_NONE,	C_SBRA,		 5, 4, 0 },
	{ ABX,		C_NONE,	C_NONE,	C_SBRA,		 74, 20, 0 },
	{ ABEQ,		C_NONE,	C_NONE,	C_SBRA,		 5, 4, 0 },

	{ AB,		C_NONE,	C_NONE,	C_ROREG,	 6, 4, 0,	LPOOL },
	{ ABL,		C_NONE,	C_NONE,	C_ROREG,	 7, 8, 0 },
	{ ABX,		C_NONE,	C_NONE,	C_ROREG,	 75, 12, 0 },
	{ ABXRET,		C_NONE,	C_NONE,	C_ROREG,	 76, 4, 0 },

	{ ASLL,		C_RCON,	C_REG,	C_REG,		 8, 4, 0 },
	{ ASLL,		C_RCON,	C_NONE,	C_REG,		 8, 4, 0 },

	{ ASLL,		C_REG,	C_NONE,	C_REG,		 9, 4, 0 },
	{ ASLL,		C_REG,	C_REG,	C_REG,		 9, 4, 0 },

	{ ASWI,		C_NONE,	C_NONE,	C_NONE,		10, 4, 0 },
	{ ASWI,		C_NONE,	C_NONE,	C_LOREG,	10, 4, 0 },
	{ ASWI,		C_NONE,	C_NONE,	C_LCON,		10, 4, 0 },

	{ AWORD,	C_NONE,	C_NONE,	C_LCON,		11, 4, 0 },
	{ AWORD,	C_NONE,	C_NONE,	C_GCON,		11, 4, 0 },
	{ AWORD,	C_NONE,	C_NONE,	C_LEXT,		11, 4, 0 },
	{ AWORD,	C_NONE,	C_NONE,	C_ADDR,		11, 4, 0 },

	{ AMOVW,	C_NCON,	C_NONE,	C_REG,		12, 4, 0 },
	{ AMOVW,	C_LCON,	C_NONE,	C_REG,		12, 4, 0,	LFROM },

	{ AADD,		C_NCON,	C_REG,	C_REG,		13, 8, 0 },
	{ AADD,		C_NCON,	C_NONE,	C_REG,		13, 8, 0 },
	{ AMVN,		C_NCON,	C_NONE,	C_REG,		13, 8, 0 },
	{ ACMP,		C_NCON,	C_REG,	C_NONE,		13, 8, 0 },
	{ AADD,		C_LCON,	C_REG,	C_REG,		13, 8, 0,	LFROM },
	{ AADD,		C_LCON,	C_NONE,	C_REG,		13, 8, 0,	LFROM },
	{ AMVN,		C_LCON,	C_NONE,	C_REG,		13, 8, 0,	LFROM },
	{ ACMP,		C_LCON,	C_REG,	C_NONE,		13, 8, 0,	LFROM },

	{ AMOVB,	C_REG,	C_NONE,	C_REG,		14, 8, 0 },
	{ AMOVBU,	C_REG,	C_NONE,	C_REG,		58, 4, 0 },
	{ AMOVH,	C_REG,	C_NONE,	C_REG,		14, 8, 0 },
	{ AMOVHU,	C_REG,	C_NONE,	C_REG,		14, 8, 0 },

	{ AMUL,		C_REG,	C_REG,	C_REG,		15, 4, 0 },
	{ AMUL,		C_REG,	C_NONE,	C_REG,		15, 4, 0 },

	{ ADIV,		C_REG,	C_REG,	C_REG,		16, 4, 0 },
	{ ADIV,		C_REG,	C_NONE,	C_REG,		16, 4, 0 },

	{ AMULL,	C_REG,	C_REG,	C_REGREG,	17, 4, 0 },

	{ AMOVW,	C_REG,	C_NONE,	C_SEXT,		20, 4, REGSB },
	{ AMOVW,	C_REG,	C_NONE,	C_SAUTO,	20, 4, REGSP },
	{ AMOVW,	C_REG,	C_NONE,	C_SOREG,	20, 4, 0 },
	{ AMOVB,	C_REG,	C_NONE,	C_SEXT,		20, 4, REGSB },
	{ AMOVB,	C_REG,	C_NONE,	C_SAUTO,	20, 4, REGSP },
	{ AMOVB,	C_REG,	C_NONE,	C_SOREG,	20, 4, 0 },
	{ AMOVBU,	C_REG,	C_NONE,	C_SEXT,		20, 4, REGSB },
	{ AMOVBU,	C_REG,	C_NONE,	C_SAUTO,	20, 4, REGSP },
	{ AMOVBU,	C_REG,	C_NONE,	C_SOREG,	20, 4, 0 },

	{ AMOVW,	C_SEXT,	C_NONE,	C_REG,		21, 4, REGSB },
	{ AMOVW,	C_SAUTO,C_NONE,	C_REG,		21, 4, REGSP },
	{ AMOVW,	C_SOREG,C_NONE,	C_REG,		21, 4, 0 },
	{ AMOVBU,	C_SEXT,	C_NONE,	C_REG,		21, 4, REGSB },
	{ AMOVBU,	C_SAUTO,C_NONE,	C_REG,		21, 4, REGSP },
	{ AMOVBU,	C_SOREG,C_NONE,	C_REG,		21, 4, 0 },

	{ AMOVB,	C_SEXT,	C_NONE,	C_REG,		22, 12, REGSB },
	{ AMOVB,	C_SAUTO,C_NONE,	C_REG,		22, 12, REGSP },
	{ AMOVB,	C_SOREG,C_NONE,	C_REG,		22, 12, 0 },
	{ AMOVH,	C_SEXT,	C_NONE,	C_REG,		22, 12, REGSB },
	{ AMOVH,	C_SAUTO,C_NONE,	C_REG,		22, 12, REGSP },
	{ AMOVH,	C_SOREG,C_NONE,	C_REG,		22, 12, 0 },
	{ AMOVHU,	C_SEXT,	C_NONE,	C_REG,		22, 12, REGSB },
	{ AMOVHU,	C_SAUTO,C_NONE,	C_REG,		22, 12, REGSP },
	{ AMOVHU,	C_SOREG,C_NONE,	C_REG,		22, 12, 0 },

	{ AMOVH,	C_REG,	C_NONE,	C_SEXT,		23, 12, REGSB },
	{ AMOVH,	C_REG,	C_NONE,	C_SAUTO,	23, 12, REGSP },
	{ AMOVH,	C_REG,	C_NONE,	C_SOREG,	23, 12, 0 },
	{ AMOVHU,	C_REG,	C_NONE,	C_SEXT,		23, 12, REGSB },
	{ AMOVHU,	C_REG,	C_NONE,	C_SAUTO,	23, 12, REGSP },
	{ AMOVHU,	C_REG,	C_NONE,	C_SOREG,	23, 12, 0 },

	{ AMOVW,	C_REG,	C_NONE,	C_LEXT,		30, 8, REGSB,	LTO },
	{ AMOVW,	C_REG,	C_NONE,	C_LAUTO,	30, 8, REGSP,	LTO },
	{ AMOVW,	C_REG,	C_NONE,	C_LOREG,	30, 8, 0,	LTO },
	{ AMOVW,	C_REG,	C_NONE,	C_ADDR,		64, 8, 0,	LTO },
	{ AMOVB,	C_REG,	C_NONE,	C_LEXT,		30, 8, REGSB,	LTO },
	{ AMOVB,	C_REG,	C_NONE,	C_LAUTO,	30, 8, REGSP,	LTO },
	{ AMOVB,	C_REG,	C_NONE,	C_LOREG,	30, 8, 0,	LTO },
	{ AMOVB,	C_REG,	C_NONE,	C_ADDR,		64, 8, 0,	LTO },
	{ AMOVBU,	C_REG,	C_NONE,	C_LEXT,		30, 8, REGSB,	LTO },
	{ AMOVBU,	C_REG,	C_NONE,	C_LAUTO,	30, 8, REGSP,	LTO },
	{ AMOVBU,	C_REG,	C_NONE,	C_LOREG,	30, 8, 0,	LTO },
	{ AMOVBU,	C_REG,	C_NONE,	C_ADDR,		64, 8, 0,	LTO },

	{ AMOVW,	C_LEXT,	C_NONE,	C_REG,		31, 8, REGSB,	LFROM },
	{ AMOVW,	C_LAUTO,C_NONE,	C_REG,		31, 8, REGSP,	LFROM },
	{ AMOVW,	C_LOREG,C_NONE,	C_REG,		31, 8, 0,	LFROM },
	{ AMOVW,	C_ADDR,	C_NONE,	C_REG,		65, 8, 0,	LFROM },
	{ AMOVBU,	C_LEXT,	C_NONE,	C_REG,		31, 8, REGSB,	LFROM },
	{ AMOVBU,	C_LAUTO,C_NONE,	C_REG,		31, 8, REGSP,	LFROM },
	{ AMOVBU,	C_LOREG,C_NONE,	C_REG,		31, 8, 0,	LFROM },
	{ AMOVBU,	C_ADDR,	C_NONE,	C_REG,		65, 8, 0,	LFROM },

	{ AMOVB,	C_LEXT,	C_NONE,	C_REG,		32, 16, REGSB,	LFROM },
	{ AMOVB,	C_LAUTO,C_NONE,	C_REG,		32, 16, REGSP,	LFROM },
	{ AMOVB,	C_LOREG,C_NONE,	C_REG,		32, 16, 0,	LFROM },
	{ AMOVB,	C_ADDR,	C_NONE,	C_REG,		66, 16, 0,	LFROM },
	{ AMOVH,	C_LEXT,	C_NONE,	C_REG,		32, 16, REGSB,	LFROM },
	{ AMOVH,	C_LAUTO,C_NONE,	C_REG,		32, 16, REGSP,	LFROM },
	{ AMOVH,	C_LOREG,C_NONE,	C_REG,		32, 16, 0,	LFROM },
	{ AMOVH,	C_ADDR,	C_NONE,	C_REG,		66, 16, 0,	LFROM },
	{ AMOVHU,	C_LEXT,	C_NONE,	C_REG,		32, 16, REGSB,	LFROM },
	{ AMOVHU,	C_LAUTO,C_NONE,	C_REG,		32, 16, REGSP,	LFROM },
	{ AMOVHU,	C_LOREG,C_NONE,	C_REG,		32, 16, 0,	LFROM },
	{ AMOVHU,	C_ADDR,	C_NONE,	C_REG,		66, 16, 0,	LFROM },

	{ AMOVH,	C_REG,	C_NONE,	C_LEXT,		33, 24, REGSB,	LTO },
	{ AMOVH,	C_REG,	C_NONE,	C_LAUTO,	33, 24, REGSP,	LTO },
	{ AMOVH,	C_REG,	C_NONE,	C_LOREG,	33, 24, 0,	LTO },
	{ AMOVH,	C_REG,	C_NONE,	C_ADDR,		67, 24, 0,	LTO },
	{ AMOVHU,	C_REG,	C_NONE,	C_LEXT,		33, 24, REGSB,	LTO },
	{ AMOVHU,	C_REG,	C_NONE,	C_LAUTO,	33, 24, REGSP,	LTO },
	{ AMOVHU,	C_REG,	C_NONE,	C_LOREG,	33, 24, 0,	LTO },
	{ AMOVHU,	C_REG,	C_NONE,	C_ADDR,		67, 24, 0,	LTO },

	{ AMOVW,	C_LECON,C_NONE,	C_REG,		34, 8, REGSB,	LFROM },
	{ AMOVW,	C_LACON,C_NONE,	C_REG,		34, 8, REGSP,	LFROM },

	{ AMOVW,	C_PSR,	C_NONE,	C_REG,		35, 4, 0 },
	{ AMOVW,	C_REG,	C_NONE,	C_PSR,		36, 4, 0 },
	{ AMOVW,	C_RCON,	C_NONE,	C_PSR,		37, 4, 0 },

	{ AMOVM,	C_LCON,	C_NONE,	C_SOREG,	38, 4, 0 },
	{ AMOVM,	C_SOREG,C_NONE,	C_LCON,		39, 4, 0 },

	{ ASWPW,	C_SOREG,C_REG,	C_REG,		40, 4, 0 },

	{ ARFE,		C_NONE,	C_NONE,	C_NONE,		41, 4, 0 },

	{ AMOVF,	C_FREG,	C_NONE,	C_FEXT,		50, 4, REGSB },
	{ AMOVF,	C_FREG,	C_NONE,	C_FAUTO,	50, 4, REGSP },
	{ AMOVF,	C_FREG,	C_NONE,	C_FOREG,	50, 4, 0 },

	{ AMOVF,	C_FEXT,	C_NONE,	C_FREG,		51, 4, REGSB },
	{ AMOVF,	C_FAUTO,C_NONE,	C_FREG,		51, 4, REGSP },
	{ AMOVF,	C_FOREG,C_NONE,	C_FREG,		51, 4, 0 },

	{ AMOVF,	C_FREG,	C_NONE,	C_LEXT,		52, 12, REGSB,	LTO },
	{ AMOVF,	C_FREG,	C_NONE,	C_LAUTO,	52, 12, REGSP,	LTO },
	{ AMOVF,	C_FREG,	C_NONE,	C_LOREG,	52, 12, 0,	LTO },

	{ AMOVF,	C_LEXT,	C_NONE,	C_FREG,		53, 12, REGSB,	LFROM },
	{ AMOVF,	C_LAUTO,C_NONE,	C_FREG,		53, 12, REGSP,	LFROM },
	{ AMOVF,	C_LOREG,C_NONE,	C_FREG,		53, 12, 0,	LFROM },

	{ AMOVF,	C_FREG,	C_NONE,	C_ADDR,		68, 8, 0,	LTO },
	{ AMOVF,	C_ADDR,	C_NONE,	C_FREG,		69, 8, 0,	LFROM },

	{ AADDF,	C_FREG,	C_NONE,	C_FREG,		54, 4, 0 },
	{ AADDF,	C_FREG,	C_REG,	C_FREG,		54, 4, 0 },
	{ AADDF,	C_FCON,	C_NONE,	C_FREG,		54, 4, 0 },
	{ AADDF,	C_FCON,	C_REG,	C_FREG,		54, 4, 0 },
	{ AMOVF,	C_FCON,	C_NONE,	C_FREG,		54, 4, 0 },
	{ AMOVF,	C_FREG, C_NONE, C_FREG,		54, 4, 0 },

	{ ACMPF,	C_FREG,	C_REG,	C_NONE,		54, 4, 0 },
	{ ACMPF,	C_FCON,	C_REG,	C_NONE,		54, 4, 0 },

	{ AMOVFW,	C_FREG,	C_NONE,	C_REG,		55, 4, 0 },
	{ AMOVFW,	C_REG,	C_NONE,	C_FREG,		55, 4, 0 },

	{ AMOVW,	C_REG,	C_NONE,	C_FCR,		56, 4, 0 },
	{ AMOVW,	C_FCR,	C_NONE,	C_REG,		57, 4, 0 },

	{ AMOVW,	C_SHIFT,C_NONE,	C_REG,		59, 4, 0 },
	{ AMOVBU,	C_SHIFT,C_NONE,	C_REG,		59, 4, 0 },

	{ AMOVB,	C_SHIFT,C_NONE,	C_REG,		60, 4, 0 },

	{ AMOVW,	C_REG,	C_NONE,	C_SHIFT,	61, 4, 0 },
	{ AMOVB,	C_REG,	C_NONE,	C_SHIFT,	61, 4, 0 },
	{ AMOVBU,	C_REG,	C_NONE,	C_SHIFT,	61, 4, 0 },

	{ ACASE,	C_REG,	C_NONE,	C_NONE,		62, 4, 0 },
	{ ABCASE,	C_NONE, C_NONE, C_SBRA,		63, 4, 0 },

	{ AMOVH,	C_REG,	C_NONE,	C_HEXT,		70, 4, REGSB,	V4 },
	{ AMOVH,	C_REG,	C_NONE, C_HAUTO,	70, 4, REGSP,	V4 },
	{ AMOVH,	C_REG,	C_NONE,	C_HOREG,	70, 4, 0,	V4 },
	{ AMOVHU,	C_REG,	C_NONE,	C_HEXT,		70, 4, REGSB,	V4 },
	{ AMOVHU,	C_REG,	C_NONE, C_HAUTO,	70, 4, REGSP,	V4 },
	{ AMOVHU,	C_REG,	C_NONE,	C_HOREG,	70, 4, 0,	V4 },

	{ AMOVB,	C_HEXT,	C_NONE, C_REG,		71, 4, REGSB,	V4 },
	{ AMOVB,	C_HAUTO,C_NONE,	C_REG,		71, 4, REGSP,	V4 },
	{ AMOVB,	C_HOREG,C_NONE,	C_REG,		71, 4, 0,	V4 },
	{ AMOVH,	C_HEXT,	C_NONE,	C_REG,		71, 4, REGSB,	V4 },
	{ AMOVH,	C_HAUTO,C_NONE, C_REG,		71, 4, REGSP,	V4 },
	{ AMOVH,	C_HOREG,C_NONE,	C_REG,		71, 4, 0,	V4 },
	{ AMOVHU,	C_HEXT,	C_NONE,	C_REG,		71, 4, REGSB,	V4 },
	{ AMOVHU,	C_HAUTO,C_NONE, C_REG,		71, 4, REGSP,	V4 },
	{ AMOVHU,	C_HOREG,C_NONE,	C_REG,		71, 4, 0,	V4 },

	{ AMOVH,	C_REG,	C_NONE,	C_LEXT,		72, 8, REGSB,	LTO|V4 },
	{ AMOVH,	C_REG,	C_NONE, C_LAUTO,	72, 8, REGSP,	LTO|V4 },
	{ AMOVH,	C_REG,	C_NONE,	C_LOREG,	72, 8, 0,	LTO|V4 },
	{ AMOVHU,	C_REG,	C_NONE,	C_LEXT,		72, 8, REGSB,	LTO|V4 },
	{ AMOVHU,	C_REG,	C_NONE, C_LAUTO,	72, 8, REGSP,	LTO|V4 },
	{ AMOVHU,	C_REG,	C_NONE,	C_LOREG,	72, 8, 0,	LTO|V4 },

	{ AMOVB,	C_LEXT,	C_NONE, C_REG,		73, 8, REGSB,	LFROM|V4 },
	{ AMOVB,	C_LAUTO,C_NONE,	C_REG,		73, 8, REGSP,	LFROM|V4 },
	{ AMOVB,	C_LOREG,C_NONE,	C_REG,		73, 8, 0,	LFROM|V4 },
	{ AMOVH,	C_LEXT,	C_NONE,	C_REG,		73, 8, REGSB,	LFROM|V4 },
	{ AMOVH,	C_LAUTO,C_NONE, C_REG,		73, 8, REGSP,	LFROM|V4 },
	{ AMOVH,	C_LOREG,C_NONE,	C_REG,		73, 8, 0,	LFROM|V4 },
	{ AMOVHU,	C_LEXT,	C_NONE,	C_REG,		73, 8, REGSB,	LFROM|V4 },
	{ AMOVHU,	C_LAUTO,C_NONE, C_REG,		73, 8, REGSP,	LFROM|V4 },
	{ AMOVHU,	C_LOREG,C_NONE,	C_REG,		73, 8, 0,	LFROM|V4 },
	{ ALDREX,	C_SOREG,C_NONE,	C_REG,		77, 4, 0 },
	{ ASTREX,	C_SOREG,C_REG,	C_REG,		78, 4, 0 },

	{ AXXX,		C_NONE,	C_NONE,	C_NONE,		 0, 4, 0 },
};
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

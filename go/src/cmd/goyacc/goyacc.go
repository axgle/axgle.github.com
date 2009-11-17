<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/cmd/goyacc/goyacc.go</title>

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
  <h1 id="generatedHeader">Source file /src/cmd/goyacc/goyacc.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">/*</span>
<a id="L2"></a><span class="comment">Derived from Inferno&#39;s utils/iyacc/yacc.c</span>
<a id="L3"></a><span class="comment">http://code.google.com/p/inferno-os/source/browse/utils/iyacc/yacc.c</span>

<a id="L5"></a><span class="comment">This copyright NOTICE applies to all files in this directory and</span>
<a id="L6"></a><span class="comment">subdirectories, unless another copyright notice appears in a given</span>
<a id="L7"></a><span class="comment">file or subdirectory.  If you take substantial code from this software to use in</span>
<a id="L8"></a><span class="comment">other programs, you must somehow include with it an appropriate</span>
<a id="L9"></a><span class="comment">copyright notice that includes the copyright notice and the other</span>
<a id="L10"></a><span class="comment">notices below.  It is fine (and often tidier) to do that in a separate</span>
<a id="L11"></a><span class="comment">file such as NOTICE, LICENCE or COPYING.</span>

<a id="L13"></a><span class="comment">	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.</span>
<a id="L14"></a><span class="comment">	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)</span>
<a id="L15"></a><span class="comment">	Portions Copyright © 1997-1999 Vita Nuova Limited</span>
<a id="L16"></a><span class="comment">	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)</span>
<a id="L17"></a><span class="comment">	Portions Copyright © 2004,2006 Bruce Ellis</span>
<a id="L18"></a><span class="comment">	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)</span>
<a id="L19"></a><span class="comment">	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others</span>
<a id="L20"></a><span class="comment">	Portions Copyright © 2009 The Go Authors.  All rights reserved.</span>

<a id="L22"></a><span class="comment">Permission is hereby granted, free of charge, to any person obtaining a copy</span>
<a id="L23"></a><span class="comment">of this software and associated documentation files (the &#34;Software&#34;), to deal</span>
<a id="L24"></a><span class="comment">in the Software without restriction, including without limitation the rights</span>
<a id="L25"></a><span class="comment">to use, copy, modify, merge, publish, distribute, sublicense, and/or sell</span>
<a id="L26"></a><span class="comment">copies of the Software, and to permit persons to whom the Software is</span>
<a id="L27"></a><span class="comment">furnished to do so, subject to the following conditions:</span>

<a id="L29"></a><span class="comment">The above copyright notice and this permission notice shall be included in</span>
<a id="L30"></a><span class="comment">all copies or substantial portions of the Software.</span>

<a id="L32"></a><span class="comment">THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR</span>
<a id="L33"></a><span class="comment">IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,</span>
<a id="L34"></a><span class="comment">FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE</span>
<a id="L35"></a><span class="comment">AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER</span>
<a id="L36"></a><span class="comment">LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,</span>
<a id="L37"></a><span class="comment">OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN</span>
<a id="L38"></a><span class="comment">THE SOFTWARE.</span>
<a id="L39"></a><span class="comment">*/</span>

<a id="L41"></a>package main

<a id="L43"></a><span class="comment">// yacc</span>
<a id="L44"></a><span class="comment">// major difference is lack of stem (&#34;y&#34; variable)</span>
<a id="L45"></a><span class="comment">//</span>

<a id="L47"></a>import (
    <a id="L48"></a>&#34;flag&#34;;
    <a id="L49"></a>&#34;fmt&#34;;
    <a id="L50"></a>&#34;bufio&#34;;
    <a id="L51"></a>&#34;os&#34;;
<a id="L52"></a>)

<a id="L54"></a><span class="comment">// the following are adjustable</span>
<a id="L55"></a><span class="comment">// according to memory size</span>
<a id="L56"></a>const (
    <a id="L57"></a>ACTSIZE  = 30000;
    <a id="L58"></a>NSTATES  = 2000;
    <a id="L59"></a>TEMPSIZE = 2000;

    <a id="L61"></a>SYMINC   = 50;  <span class="comment">// increase for non-term or term</span>
    <a id="L62"></a>RULEINC  = 50;  <span class="comment">// increase for max rule length prodptr[i]</span>
    <a id="L63"></a>PRODINC  = 100; <span class="comment">// increase for productions     prodptr</span>
    <a id="L64"></a>WSETINC  = 50;  <span class="comment">// increase for working sets    wsets</span>
    <a id="L65"></a>STATEINC = 200; <span class="comment">// increase for states          statemem</span>

    <a id="L67"></a>NAMESIZE = 50;
    <a id="L68"></a>NTYPES   = 63;
    <a id="L69"></a>ISIZE    = 400;

    <a id="L71"></a>PRIVATE = 0xE000; <span class="comment">// unicode private use</span>

    <a id="L73"></a><span class="comment">// relationships which must hold:</span>
    <a id="L74"></a><span class="comment">//	TEMPSIZE &gt;= NTERMS + NNONTERM + 1;</span>
    <a id="L75"></a><span class="comment">//	TEMPSIZE &gt;= NSTATES;</span>
    <a id="L76"></a><span class="comment">//</span>

    <a id="L78"></a>NTBASE     = 010000;
    <a id="L79"></a>ERRCODE    = 8190;
    <a id="L80"></a>ACCEPTCODE = 8191;
    <a id="L81"></a>YYLEXUNK   = 3;
    <a id="L82"></a>TOKSTART   = 4; <span class="comment">//index of first defined token</span>
<a id="L83"></a>)

<a id="L85"></a><span class="comment">// no, left, right, binary assoc.</span>
<a id="L86"></a>const (
    <a id="L87"></a>NOASC = iota;
    <a id="L88"></a>LASC;
    <a id="L89"></a>RASC;
    <a id="L90"></a>BASC;
<a id="L91"></a>)

<a id="L93"></a><span class="comment">// flags for state generation</span>
<a id="L94"></a>const (
    <a id="L95"></a>DONE = iota;
    <a id="L96"></a>MUSTDO;
    <a id="L97"></a>MUSTLOOKAHEAD;
<a id="L98"></a>)

<a id="L100"></a><span class="comment">// flags for a rule having an action, and being reduced</span>
<a id="L101"></a>const (
    <a id="L102"></a>ACTFLAG = 1 &lt;&lt; (iota + 2);
    <a id="L103"></a>REDFLAG;
<a id="L104"></a>)

<a id="L106"></a><span class="comment">// output parser flags</span>
<a id="L107"></a>const YYFLAG = -1000

<a id="L109"></a><span class="comment">// parse tokens</span>
<a id="L110"></a>const (
    <a id="L111"></a>IDENTIFIER = PRIVATE + iota;
    <a id="L112"></a>MARK;
    <a id="L113"></a>TERM;
    <a id="L114"></a>LEFT;
    <a id="L115"></a>RIGHT;
    <a id="L116"></a>BINARY;
    <a id="L117"></a>PREC;
    <a id="L118"></a>LCURLY;
    <a id="L119"></a>IDENTCOLON;
    <a id="L120"></a>NUMBER;
    <a id="L121"></a>START;
    <a id="L122"></a>TYPEDEF;
    <a id="L123"></a>TYPENAME;
    <a id="L124"></a>UNION;
<a id="L125"></a>)

<a id="L127"></a>const ENDFILE = 0
<a id="L128"></a>const EMPTY = 1
<a id="L129"></a>const WHOKNOWS = 0
<a id="L130"></a>const OK = 1
<a id="L131"></a>const NOMORE = -1000

<a id="L133"></a><span class="comment">// macros for getting associativity and precedence levels</span>
<a id="L134"></a>func ASSOC(i int) int { return i &amp; 3 }

<a id="L136"></a>func PLEVEL(i int) int { return (i &gt;&gt; 4) &amp; 077 }

<a id="L138"></a>func TYPE(i int) int { return (i &gt;&gt; 10) &amp; 077 }

<a id="L140"></a><span class="comment">// macros for setting associativity and precedence levels</span>
<a id="L141"></a>func SETASC(i, j int) int { return i | j }

<a id="L143"></a>func SETPLEV(i, j int) int { return i | (j &lt;&lt; 4) }

<a id="L145"></a>func SETTYPE(i, j int) int { return i | (j &lt;&lt; 10) }

<a id="L147"></a><span class="comment">// I/O descriptors</span>
<a id="L148"></a>var finput *bufio.Reader <span class="comment">// input file</span>
<a id="L149"></a>var stderr *bufio.Writer
<a id="L150"></a>var ftable *bufio.Writer  <span class="comment">// y.go file</span>
<a id="L151"></a>var foutput *bufio.Writer <span class="comment">// y.output file</span>

<a id="L153"></a>var oflag string <span class="comment">// -o [y.go]		- y.go file</span>
<a id="L154"></a>var vflag string <span class="comment">// -v [y.output]	- y.output file</span>
<a id="L155"></a>var lflag bool   <span class="comment">// -l			- disable line directives</span>

<a id="L157"></a>var stacksize = 200

<a id="L159"></a><span class="comment">// communication variables between various I/O routines</span>
<a id="L160"></a>var infile string  <span class="comment">// input file name</span>
<a id="L161"></a>var numbval int    <span class="comment">// value of an input number</span>
<a id="L162"></a>var tokname string <span class="comment">// input token name, slop for runes and 0</span>
<a id="L163"></a>var tokflag = false

<a id="L165"></a><span class="comment">// structure declarations</span>
<a id="L166"></a>type Lkset []int

<a id="L168"></a>type Pitem struct {
    <a id="L169"></a>prod   []int;
    <a id="L170"></a>off    int; <span class="comment">// offset within the production</span>
    <a id="L171"></a>first  int; <span class="comment">// first term or non-term in item</span>
    <a id="L172"></a>prodno int; <span class="comment">// production number for sorting</span>
<a id="L173"></a>}

<a id="L175"></a>type Item struct {
    <a id="L176"></a>pitem Pitem;
    <a id="L177"></a>look  Lkset;
<a id="L178"></a>}

<a id="L180"></a>type Symb struct {
    <a id="L181"></a>name  string;
    <a id="L182"></a>value int;
<a id="L183"></a>}

<a id="L185"></a>type Wset struct {
    <a id="L186"></a>pitem Pitem;
    <a id="L187"></a>flag  int;
    <a id="L188"></a>ws    Lkset;
<a id="L189"></a>}

<a id="L191"></a><span class="comment">// storage of types</span>
<a id="L192"></a>var ntypes int             <span class="comment">// number of types defined</span>
<a id="L193"></a>var typeset [NTYPES]string <span class="comment">// pointers to type tags</span>

<a id="L195"></a><span class="comment">// token information</span>

<a id="L197"></a>var ntokens = 0 <span class="comment">// number of tokens</span>
<a id="L198"></a>var tokset []Symb
<a id="L199"></a>var toklev []int <span class="comment">// vector with the precedence of the terminals</span>

<a id="L201"></a><span class="comment">// nonterminal information</span>

<a id="L203"></a>var nnonter = -1 <span class="comment">// the number of nonterminals</span>
<a id="L204"></a>var nontrst []Symb
<a id="L205"></a>var start int <span class="comment">// start symbol</span>

<a id="L207"></a><span class="comment">// state information</span>

<a id="L209"></a>var nstate = 0                      <span class="comment">// number of states</span>
<a id="L210"></a>var pstate = make([]int, NSTATES+2) <span class="comment">// index into statemem to the descriptions of the states</span>
<a id="L211"></a>var statemem []Item
<a id="L212"></a>var tystate = make([]int, NSTATES) <span class="comment">// contains type information about the states</span>
<a id="L213"></a>var tstates []int                  <span class="comment">// states generated by terminal gotos</span>
<a id="L214"></a>var ntstates []int                 <span class="comment">// states generated by nonterminal gotos</span>
<a id="L215"></a>var mstates = make([]int, NSTATES) <span class="comment">// chain of overflows of term/nonterm generation lists</span>
<a id="L216"></a>var lastred int                    <span class="comment">// number of last reduction of a state</span>
<a id="L217"></a>var defact = make([]int, NSTATES)  <span class="comment">// default actions of states</span>

<a id="L219"></a><span class="comment">// lookahead set information</span>

<a id="L221"></a>var lkst []Lkset
<a id="L222"></a>var nolook = 0  <span class="comment">// flag to turn off lookahead computations</span>
<a id="L223"></a>var tbitset = 0 <span class="comment">// size of lookahead sets</span>
<a id="L224"></a>var clset Lkset <span class="comment">// temporary storage for lookahead computations</span>

<a id="L226"></a><span class="comment">// working set information</span>

<a id="L228"></a>var wsets []Wset
<a id="L229"></a>var cwp int

<a id="L231"></a><span class="comment">// storage for action table</span>

<a id="L233"></a>var amem []int                   <span class="comment">// action table storage</span>
<a id="L234"></a>var memp int                     <span class="comment">// next free action table position</span>
<a id="L235"></a>var indgo = make([]int, NSTATES) <span class="comment">// index to the stored goto table</span>

<a id="L237"></a><span class="comment">// temporary vector, indexable by states, terms, or ntokens</span>

<a id="L239"></a>var temp1 = make([]int, TEMPSIZE) <span class="comment">// temporary storage, indexed by terms + ntokens or states</span>
<a id="L240"></a>var lineno = 1                    <span class="comment">// current input line number</span>
<a id="L241"></a>var fatfl = 1                     <span class="comment">// if on, error is fatal</span>
<a id="L242"></a>var nerrors = 0                   <span class="comment">// number of errors</span>

<a id="L244"></a><span class="comment">// assigned token type values</span>

<a id="L246"></a>var extval = 0

<a id="L248"></a><span class="comment">// grammar rule information</span>

<a id="L250"></a>var nprod = 1      <span class="comment">// number of productions</span>
<a id="L251"></a>var prdptr [][]int <span class="comment">// pointers to descriptions of productions</span>
<a id="L252"></a>var levprd []int   <span class="comment">// precedence levels for the productions</span>
<a id="L253"></a>var rlines []int   <span class="comment">// line number for this rule</span>

<a id="L255"></a><span class="comment">// statistics collection variables</span>

<a id="L257"></a>var zzgoent = 0
<a id="L258"></a>var zzgobest = 0
<a id="L259"></a>var zzacent = 0
<a id="L260"></a>var zzexcp = 0
<a id="L261"></a>var zzclose = 0
<a id="L262"></a>var zzrrconf = 0
<a id="L263"></a>var zzsrconf = 0
<a id="L264"></a>var zzstate = 0

<a id="L266"></a><span class="comment">// optimizer arrays</span>

<a id="L268"></a>var yypgo [][]int
<a id="L269"></a>var optst [][]int
<a id="L270"></a>var ggreed []int
<a id="L271"></a>var pgo []int

<a id="L273"></a>var maxspr int <span class="comment">// maximum spread of any entry</span>
<a id="L274"></a>var maxoff int <span class="comment">// maximum offset into a array</span>
<a id="L275"></a>var maxa int

<a id="L277"></a><span class="comment">// storage for information about the nonterminals</span>

<a id="L279"></a>var pres [][][]int <span class="comment">// vector of pointers to productions yielding each nonterminal</span>
<a id="L280"></a>var pfirst []Lkset
<a id="L281"></a>var pempty []int <span class="comment">// vector of nonterminals nontrivially deriving e</span>

<a id="L283"></a><span class="comment">// random stuff picked out from between functions</span>

<a id="L285"></a>var indebug = 0 <span class="comment">// debugging flag for cpfir</span>
<a id="L286"></a>var pidebug = 0 <span class="comment">// debugging flag for putitem</span>
<a id="L287"></a>var gsdebug = 0 <span class="comment">// debugging flag for stagen</span>
<a id="L288"></a>var cldebug = 0 <span class="comment">// debugging flag for closure</span>
<a id="L289"></a>var pkdebug = 0 <span class="comment">// debugging flag for apack</span>
<a id="L290"></a>var g2debug = 0 <span class="comment">// debugging for go2gen</span>
<a id="L291"></a>var adb = 0     <span class="comment">// debugging for callopt</span>

<a id="L293"></a>type Resrv struct {
    <a id="L294"></a>name  string;
    <a id="L295"></a>value int;
<a id="L296"></a>}

<a id="L298"></a>var resrv = []Resrv{
    <a id="L299"></a>Resrv{&#34;binary&#34;, BINARY},
    <a id="L300"></a>Resrv{&#34;left&#34;, LEFT},
    <a id="L301"></a>Resrv{&#34;nonassoc&#34;, BINARY},
    <a id="L302"></a>Resrv{&#34;prec&#34;, PREC},
    <a id="L303"></a>Resrv{&#34;right&#34;, RIGHT},
    <a id="L304"></a>Resrv{&#34;start&#34;, START},
    <a id="L305"></a>Resrv{&#34;term&#34;, TERM},
    <a id="L306"></a>Resrv{&#34;token&#34;, TERM},
    <a id="L307"></a>Resrv{&#34;type&#34;, TYPEDEF},
    <a id="L308"></a>Resrv{&#34;union&#34;, UNION},
    <a id="L309"></a>Resrv{&#34;struct&#34;, UNION},
<a id="L310"></a>}

<a id="L312"></a>var zznewstate = 0

<a id="L314"></a>const EOF = -1
<a id="L315"></a>const UTFmax = 0x3f

<a id="L317"></a>func main() {

    <a id="L319"></a>setup(); <span class="comment">// initialize and read productions</span>

    <a id="L321"></a>tbitset = (ntokens + 32) / 32;
    <a id="L322"></a>cpres();  <span class="comment">// make table of which productions yield a given nonterminal</span>
    <a id="L323"></a>cempty(); <span class="comment">// make a table of which nonterminals can match the empty string</span>
    <a id="L324"></a>cpfir();  <span class="comment">// make a table of firsts of nonterminals</span>

    <a id="L326"></a>stagen(); <span class="comment">// generate the states</span>

    <a id="L328"></a>yypgo = make([][]int, nnonter+1);
    <a id="L329"></a>optst = make([][]int, nstate);
    <a id="L330"></a>output(); <span class="comment">// write the states and the tables</span>
    <a id="L331"></a>go2out();

    <a id="L333"></a>hideprod();
    <a id="L334"></a>summary();

    <a id="L336"></a>callopt();

    <a id="L338"></a>others();

    <a id="L340"></a>exit(0);
<a id="L341"></a>}

<a id="L343"></a>func setup() {
    <a id="L344"></a>var j, ty int;

    <a id="L346"></a>stderr = bufio.NewWriter(os.NewFile(2, &#34;stderr&#34;));
    <a id="L347"></a>foutput = nil;

    <a id="L349"></a>flag.StringVar(&amp;oflag, &#34;o&#34;, &#34;&#34;, &#34;parser output&#34;);
    <a id="L350"></a>flag.StringVar(&amp;vflag, &#34;v&#34;, &#34;&#34;, &#34;create parsing tables&#34;);
    <a id="L351"></a>flag.BoolVar(&amp;lflag, &#34;l&#34;, false, &#34;disable line directives&#34;);

    <a id="L353"></a>flag.Parse();
    <a id="L354"></a>if flag.NArg() != 1 {
        <a id="L355"></a>usage()
    <a id="L356"></a>}
    <a id="L357"></a>if stacksize &lt; 1 {
        <a id="L358"></a><span class="comment">// never set so cannot happen</span>
        <a id="L359"></a>fmt.Fprintf(stderr, &#34;yacc: stack size too small\n&#34;);
        <a id="L360"></a>usage();
    <a id="L361"></a>}
    <a id="L362"></a>openup();

    <a id="L364"></a>defin(0, &#34;$end&#34;);
    <a id="L365"></a>extval = PRIVATE; <span class="comment">// tokens start in unicode &#39;private use&#39;</span>
    <a id="L366"></a>defin(0, &#34;error&#34;);
    <a id="L367"></a>defin(1, &#34;$accept&#34;);
    <a id="L368"></a>defin(0, &#34;$unk&#34;);
    <a id="L369"></a>i := 0;

    <a id="L371"></a>t := gettok();

<a id="L373"></a>outer:
    <a id="L374"></a>for {
        <a id="L375"></a>switch t {
        <a id="L376"></a>default:
            <a id="L377"></a>error(&#34;syntax error tok=%v&#34;, t-PRIVATE)

        <a id="L379"></a>case MARK, ENDFILE:
            <a id="L380"></a>break outer

        <a id="L382"></a>case &#39;;&#39;:

        <a id="L384"></a>case START:
            <a id="L385"></a>t = gettok();
            <a id="L386"></a>if t != IDENTIFIER {
                <a id="L387"></a>error(&#34;bad %%start construction&#34;)
            <a id="L388"></a>}
            <a id="L389"></a>start = chfind(1, tokname);

        <a id="L391"></a>case TYPEDEF:
            <a id="L392"></a>t = gettok();
            <a id="L393"></a>if t != TYPENAME {
                <a id="L394"></a>error(&#34;bad syntax in %%type&#34;)
            <a id="L395"></a>}
            <a id="L396"></a>ty = numbval;
            <a id="L397"></a>for {
                <a id="L398"></a>t = gettok();
                <a id="L399"></a>switch t {
                <a id="L400"></a>case IDENTIFIER:
                    <a id="L401"></a>t = chfind(1, tokname);
                    <a id="L402"></a>if t &lt; NTBASE {
                        <a id="L403"></a>j = TYPE(toklev[t]);
                        <a id="L404"></a>if j != 0 &amp;&amp; j != ty {
                            <a id="L405"></a>error(&#34;type redeclaration of token &#34;,
                                <a id="L406"></a>tokset[t].name)
                        <a id="L407"></a>} else {
                            <a id="L408"></a>toklev[t] = SETTYPE(toklev[t], ty)
                        <a id="L409"></a>}
                    <a id="L410"></a>} else {
                        <a id="L411"></a>j = nontrst[t-NTBASE].value;
                        <a id="L412"></a>if j != 0 &amp;&amp; j != ty {
                            <a id="L413"></a>error(&#34;type redeclaration of nonterminal %v&#34;,
                                <a id="L414"></a>nontrst[t-NTBASE].name)
                        <a id="L415"></a>} else {
                            <a id="L416"></a>nontrst[t-NTBASE].value = ty
                        <a id="L417"></a>}
                    <a id="L418"></a>}
                    <a id="L419"></a>continue;

                <a id="L421"></a>case &#39;,&#39;:
                    <a id="L422"></a>continue
                <a id="L423"></a>}
                <a id="L424"></a>break;
            <a id="L425"></a>}
            <a id="L426"></a>continue;

        <a id="L428"></a>case UNION:
            <a id="L429"></a>cpyunion()

        <a id="L431"></a>case LEFT, BINARY, RIGHT, TERM:
            <a id="L432"></a><span class="comment">// nonzero means new prec. and assoc.</span>
            <a id="L433"></a>lev := t - TERM;
            <a id="L434"></a>if lev != 0 {
                <a id="L435"></a>i++
            <a id="L436"></a>}
            <a id="L437"></a>ty = 0;

            <a id="L439"></a><span class="comment">// get identifiers so defined</span>
            <a id="L440"></a>t = gettok();

            <a id="L442"></a><span class="comment">// there is a type defined</span>
            <a id="L443"></a>if t == TYPENAME {
                <a id="L444"></a>ty = numbval;
                <a id="L445"></a>t = gettok();
            <a id="L446"></a>}
            <a id="L447"></a>for {
                <a id="L448"></a>switch t {
                <a id="L449"></a>case &#39;,&#39;:
                    <a id="L450"></a>t = gettok();
                    <a id="L451"></a>continue;

                <a id="L453"></a>case &#39;;&#39;:
                    <a id="L454"></a>break

                <a id="L456"></a>case IDENTIFIER:
                    <a id="L457"></a>j = chfind(0, tokname);
                    <a id="L458"></a>if j &gt;= NTBASE {
                        <a id="L459"></a>error(&#34;%v defined earlier as nonterminal&#34;, tokname)
                    <a id="L460"></a>}
                    <a id="L461"></a>if lev != 0 {
                        <a id="L462"></a>if ASSOC(toklev[j]) != 0 {
                            <a id="L463"></a>error(&#34;redeclaration of precedence of %v&#34;, tokname)
                        <a id="L464"></a>}
                        <a id="L465"></a>toklev[j] = SETASC(toklev[j], lev);
                        <a id="L466"></a>toklev[j] = SETPLEV(toklev[j], i);
                    <a id="L467"></a>}
                    <a id="L468"></a>if ty != 0 {
                        <a id="L469"></a>if TYPE(toklev[j]) != 0 {
                            <a id="L470"></a>error(&#34;redeclaration of type of %v&#34;, tokname)
                        <a id="L471"></a>}
                        <a id="L472"></a>toklev[j] = SETTYPE(toklev[j], ty);
                    <a id="L473"></a>}
                    <a id="L474"></a>t = gettok();
                    <a id="L475"></a>if t == NUMBER {
                        <a id="L476"></a>tokset[j].value = numbval;
                        <a id="L477"></a>t = gettok();
                    <a id="L478"></a>}

                    <a id="L480"></a>continue;
                <a id="L481"></a>}
                <a id="L482"></a>break;
            <a id="L483"></a>}
            <a id="L484"></a>continue;

        <a id="L486"></a>case LCURLY:
            <a id="L487"></a>cpycode()
        <a id="L488"></a>}
        <a id="L489"></a>t = gettok();
    <a id="L490"></a>}

    <a id="L492"></a>if t == ENDFILE {
        <a id="L493"></a>error(&#34;unexpected EOF before %%&#34;)
    <a id="L494"></a>}

    <a id="L496"></a><span class="comment">// put out non-literal terminals</span>
    <a id="L497"></a>for i := TOKSTART; i &lt;= ntokens; i++ {
        <a id="L498"></a><span class="comment">// non-literals</span>
        <a id="L499"></a>c := tokset[i].name[0];
        <a id="L500"></a>if c != &#39; &#39; &amp;&amp; c != &#39;$&#39; {
            <a id="L501"></a>fmt.Fprintf(ftable, &#34;const\t%v\t= %v\n&#34;, tokset[i].name, tokset[i].value)
        <a id="L502"></a>}
    <a id="L503"></a>}

    <a id="L505"></a><span class="comment">// put out names of token names</span>
    <a id="L506"></a>fmt.Fprintf(ftable, &#34;var\tToknames\t =[]string {\n&#34;);
    <a id="L507"></a>for i := TOKSTART; i &lt;= ntokens; i++ {
        <a id="L508"></a>fmt.Fprintf(ftable, &#34;\t\&#34;%v\&#34;,\n&#34;, tokset[i].name)
    <a id="L509"></a>}
    <a id="L510"></a>fmt.Fprintf(ftable, &#34;}\n&#34;);

    <a id="L512"></a><span class="comment">// put out names of state names</span>
    <a id="L513"></a>fmt.Fprintf(ftable, &#34;var\tStatenames\t =[]string {\n&#34;);
    <a id="L514"></a><span class="comment">//	for i:=TOKSTART; i&lt;=ntokens; i++ {</span>
    <a id="L515"></a><span class="comment">//		fmt.Fprintf(ftable, &#34;\t\&#34;%v\&#34;,\n&#34;, tokset[i].name);</span>
    <a id="L516"></a><span class="comment">//	}</span>
    <a id="L517"></a>fmt.Fprintf(ftable, &#34;}\n&#34;);

    <a id="L519"></a>fmt.Fprintf(ftable, &#34;\nfunc\n&#34;);
    <a id="L520"></a>fmt.Fprintf(ftable, &#34;yyrun(p int, yypt int) {\n&#34;);
    <a id="L521"></a>fmt.Fprintf(ftable, &#34;switch p {\n&#34;);

    <a id="L523"></a>moreprod();
    <a id="L524"></a>prdptr[0] = []int{NTBASE, start, 1, 0};

    <a id="L526"></a>nprod = 1;
    <a id="L527"></a>curprod := make([]int, RULEINC);
    <a id="L528"></a>t = gettok();
    <a id="L529"></a>if t != IDENTCOLON {
        <a id="L530"></a>error(&#34;bad syntax on first rule&#34;)
    <a id="L531"></a>}

    <a id="L533"></a>if start == 0 {
        <a id="L534"></a>prdptr[0][1] = chfind(1, tokname)
    <a id="L535"></a>}

    <a id="L537"></a><span class="comment">// read rules</span>
    <a id="L538"></a><span class="comment">// put into prdptr array in the format</span>
    <a id="L539"></a><span class="comment">// target</span>
    <a id="L540"></a><span class="comment">// followed by id&#39;s of terminals and non-terminals</span>
    <a id="L541"></a><span class="comment">// followd by -nprod</span>

    <a id="L543"></a>for t != MARK &amp;&amp; t != ENDFILE {
        <a id="L544"></a>mem := 0;

        <a id="L546"></a><span class="comment">// process a rule</span>
        <a id="L547"></a>rlines[nprod] = lineno;
        <a id="L548"></a>if t == &#39;|&#39; {
            <a id="L549"></a>curprod[mem] = prdptr[nprod-1][0];
            <a id="L550"></a>mem++;
        <a id="L551"></a>} else if t == IDENTCOLON {
            <a id="L552"></a>curprod[mem] = chfind(1, tokname);
            <a id="L553"></a>if curprod[mem] &lt; NTBASE {
                <a id="L554"></a>error(&#34;token illegal on LHS of grammar rule&#34;)
            <a id="L555"></a>}
            <a id="L556"></a>mem++;
        <a id="L557"></a>} else {
            <a id="L558"></a>error(&#34;illegal rule: missing semicolon or | ?&#34;)
        <a id="L559"></a>}

        <a id="L561"></a><span class="comment">// read rule body</span>
        <a id="L562"></a>t = gettok();
        <a id="L563"></a>for {
            <a id="L564"></a>for t == IDENTIFIER {
                <a id="L565"></a>curprod[mem] = chfind(1, tokname);
                <a id="L566"></a>if curprod[mem] &lt; NTBASE {
                    <a id="L567"></a>levprd[nprod] = toklev[curprod[mem]]
                <a id="L568"></a>}
                <a id="L569"></a>mem++;
                <a id="L570"></a>if mem &gt;= len(curprod) {
                    <a id="L571"></a>ncurprod := make([]int, mem+RULEINC);
                    <a id="L572"></a>for ll := 0; ll &lt; mem; ll++ {
                        <a id="L573"></a>ncurprod[ll] = curprod[ll]
                    <a id="L574"></a>}
                    <a id="L575"></a>curprod = ncurprod;
                <a id="L576"></a>}
                <a id="L577"></a>t = gettok();
            <a id="L578"></a>}
            <a id="L579"></a>if t == PREC {
                <a id="L580"></a>if gettok() != IDENTIFIER {
                    <a id="L581"></a>error(&#34;illegal %%prec syntax&#34;)
                <a id="L582"></a>}
                <a id="L583"></a>j = chfind(2, tokname);
                <a id="L584"></a>if j &gt;= NTBASE {
                    <a id="L585"></a>error(&#34;nonterminal &#34; + nontrst[j-NTBASE].name + &#34; illegal after %%prec&#34;)
                <a id="L586"></a>}
                <a id="L587"></a>levprd[nprod] = toklev[j];
                <a id="L588"></a>t = gettok();
            <a id="L589"></a>}
            <a id="L590"></a>if t != &#39;=&#39; {
                <a id="L591"></a>break
            <a id="L592"></a>}
            <a id="L593"></a>levprd[nprod] |= ACTFLAG;
            <a id="L594"></a>fmt.Fprintf(ftable, &#34;\ncase %v:&#34;, nprod);
            <a id="L595"></a>cpyact(curprod, mem);

            <a id="L597"></a><span class="comment">// action within rule...</span>
            <a id="L598"></a>t = gettok();
            <a id="L599"></a>if t == IDENTIFIER {
                <a id="L600"></a><span class="comment">// make it a nonterminal</span>
                <a id="L601"></a>j = chfind(1, fmt.Sprintf(&#34;$$%v&#34;, nprod));

                <a id="L603"></a><span class="comment">//</span>
                <a id="L604"></a><span class="comment">// the current rule will become rule number nprod+1</span>
                <a id="L605"></a><span class="comment">// enter null production for action</span>
                <a id="L606"></a><span class="comment">//</span>
                <a id="L607"></a>prdptr[nprod] = make([]int, 2);
                <a id="L608"></a>prdptr[nprod][0] = j;
                <a id="L609"></a>prdptr[nprod][1] = -nprod;

                <a id="L611"></a><span class="comment">// update the production information</span>
                <a id="L612"></a>nprod++;
                <a id="L613"></a>moreprod();
                <a id="L614"></a>levprd[nprod] = levprd[nprod-1] &amp; ^ACTFLAG;
                <a id="L615"></a>levprd[nprod-1] = ACTFLAG;
                <a id="L616"></a>rlines[nprod] = lineno;

                <a id="L618"></a><span class="comment">// make the action appear in the original rule</span>
                <a id="L619"></a>curprod[mem] = j;
                <a id="L620"></a>mem++;
                <a id="L621"></a>if mem &gt;= len(curprod) {
                    <a id="L622"></a>ncurprod := make([]int, mem+RULEINC);
                    <a id="L623"></a>for ll := 0; ll &lt; mem; ll++ {
                        <a id="L624"></a>ncurprod[ll] = curprod[ll]
                    <a id="L625"></a>}
                    <a id="L626"></a>curprod = ncurprod;
                <a id="L627"></a>}
            <a id="L628"></a>}
        <a id="L629"></a>}

        <a id="L631"></a>for t == &#39;;&#39; {
            <a id="L632"></a>t = gettok()
        <a id="L633"></a>}
        <a id="L634"></a>curprod[mem] = -nprod;
        <a id="L635"></a>mem++;

        <a id="L637"></a><span class="comment">// check that default action is reasonable</span>
        <a id="L638"></a>if ntypes != 0 &amp;&amp; (levprd[nprod]&amp;ACTFLAG) == 0 &amp;&amp;
            <a id="L639"></a>nontrst[curprod[0]-NTBASE].value != 0 {
            <a id="L640"></a><span class="comment">// no explicit action, LHS has value</span>
            <a id="L641"></a>tempty := curprod[1];
            <a id="L642"></a>if tempty &lt; 0 {
                <a id="L643"></a>error(&#34;must return a value, since LHS has a type&#34;)
            <a id="L644"></a>}
            <a id="L645"></a>if tempty &gt;= NTBASE {
                <a id="L646"></a>tempty = nontrst[tempty-NTBASE].value
            <a id="L647"></a>} else {
                <a id="L648"></a>tempty = TYPE(toklev[tempty])
            <a id="L649"></a>}
            <a id="L650"></a>if tempty != nontrst[curprod[0]-NTBASE].value {
                <a id="L651"></a>error(&#34;default action causes potential type clash&#34;)
            <a id="L652"></a>}
            <a id="L653"></a>fmt.Fprintf(ftable, &#34;\ncase %v:&#34;, nprod);
            <a id="L654"></a>fmt.Fprintf(ftable, &#34;\n\tYYVAL.%v = YYS[yypt-0].%v;&#34;,
                <a id="L655"></a>typeset[tempty], typeset[tempty]);
        <a id="L656"></a>}
        <a id="L657"></a>moreprod();
        <a id="L658"></a>prdptr[nprod] = make([]int, mem);
        <a id="L659"></a>for ll := 0; ll &lt; mem; ll++ {
            <a id="L660"></a>prdptr[nprod][ll] = curprod[ll]
        <a id="L661"></a>}
        <a id="L662"></a>nprod++;
        <a id="L663"></a>moreprod();
        <a id="L664"></a>levprd[nprod] = 0;
    <a id="L665"></a>}

    <a id="L667"></a><span class="comment">//</span>
    <a id="L668"></a><span class="comment">// end of all rules</span>
    <a id="L669"></a><span class="comment">// dump out the prefix code</span>
    <a id="L670"></a><span class="comment">//</span>

    <a id="L672"></a>fmt.Fprintf(ftable, &#34;\n\t}&#34;);
    <a id="L673"></a>fmt.Fprintf(ftable, &#34;\n}\n&#34;);

    <a id="L675"></a>fmt.Fprintf(ftable, &#34;const	YYEOFCODE	= 1\n&#34;);
    <a id="L676"></a>fmt.Fprintf(ftable, &#34;const	YYERRCODE	= 2\n&#34;);
    <a id="L677"></a>fmt.Fprintf(ftable, &#34;const	YYMAXDEPTH	= %v\n&#34;, stacksize);

    <a id="L679"></a><span class="comment">//</span>
    <a id="L680"></a><span class="comment">// copy any postfix code</span>
    <a id="L681"></a><span class="comment">//</span>
    <a id="L682"></a>if t == MARK {
        <a id="L683"></a>if !lflag {
            <a id="L684"></a>fmt.Fprintf(ftable, &#34;\n//line %v:%v\n&#34;, infile, lineno)
        <a id="L685"></a>}
        <a id="L686"></a>for {
            <a id="L687"></a>c := getrune(finput);
            <a id="L688"></a>if c == EOF {
                <a id="L689"></a>break
            <a id="L690"></a>}
            <a id="L691"></a>putrune(ftable, c);
        <a id="L692"></a>}
    <a id="L693"></a>}
<a id="L694"></a>}

<a id="L696"></a><span class="comment">//</span>
<a id="L697"></a><span class="comment">// allocate enough room to hold another production</span>
<a id="L698"></a><span class="comment">//</span>
<a id="L699"></a>func moreprod() {
    <a id="L700"></a>n := len(prdptr);
    <a id="L701"></a>if nprod &gt;= n {
        <a id="L702"></a>nn := n + PRODINC;
        <a id="L703"></a>aprod := make([][]int, nn);
        <a id="L704"></a>alevprd := make([]int, nn);
        <a id="L705"></a>arlines := make([]int, nn);

        <a id="L707"></a>for ll := 0; ll &lt; n; ll++ {
            <a id="L708"></a>aprod[ll] = prdptr[ll];
            <a id="L709"></a>alevprd[ll] = levprd[ll];
            <a id="L710"></a>arlines[ll] = rlines[ll];
        <a id="L711"></a>}

        <a id="L713"></a>prdptr = aprod;
        <a id="L714"></a>levprd = alevprd;
        <a id="L715"></a>rlines = arlines;
    <a id="L716"></a>}
<a id="L717"></a>}

<a id="L719"></a><span class="comment">//</span>
<a id="L720"></a><span class="comment">// define s to be a terminal if t=0</span>
<a id="L721"></a><span class="comment">// or a nonterminal if t=1</span>
<a id="L722"></a><span class="comment">//</span>
<a id="L723"></a>func defin(nt int, s string) int {
    <a id="L724"></a>val := 0;
    <a id="L725"></a>if nt != 0 {
        <a id="L726"></a>nnonter++;
        <a id="L727"></a>if nnonter &gt;= len(nontrst) {
            <a id="L728"></a>anontrst := make([]Symb, nnonter+SYMINC);
            <a id="L729"></a>for ll := 0; ll &lt; len(nontrst); ll++ {
                <a id="L730"></a>anontrst[ll] = nontrst[ll]
            <a id="L731"></a>}
            <a id="L732"></a>nontrst = anontrst;
        <a id="L733"></a>}
        <a id="L734"></a>nontrst[nnonter] = Symb{s, 0};
        <a id="L735"></a>return NTBASE + nnonter;
    <a id="L736"></a>}

    <a id="L738"></a><span class="comment">// must be a token</span>
    <a id="L739"></a>ntokens++;
    <a id="L740"></a>if ntokens &gt;= len(tokset) {
        <a id="L741"></a>nn := ntokens + SYMINC;
        <a id="L742"></a>atokset := make([]Symb, nn);
        <a id="L743"></a>atoklev := make([]int, nn);

        <a id="L745"></a>for ll := 0; ll &lt; len(tokset); ll++ {
            <a id="L746"></a>atoklev[ll] = toklev[ll];
            <a id="L747"></a>atokset[ll] = tokset[ll];
        <a id="L748"></a>}

        <a id="L750"></a>tokset = atokset;
        <a id="L751"></a>toklev = atoklev;
    <a id="L752"></a>}
    <a id="L753"></a>tokset[ntokens].name = s;
    <a id="L754"></a>toklev[ntokens] = 0;

    <a id="L756"></a><span class="comment">// establish value for token</span>
    <a id="L757"></a><span class="comment">// single character literal</span>
    <a id="L758"></a>if s[0] == &#39; &#39; &amp;&amp; len(s) == 1+1 {
        <a id="L759"></a>val = int(s[1])
    <a id="L760"></a>} else if s[0] == &#39; &#39; &amp;&amp; s[1] == &#39;\\&#39; { <span class="comment">// escape sequence</span>
        <a id="L761"></a>if len(s) == 2+1 {
            <a id="L762"></a><span class="comment">// single character escape sequence</span>
            <a id="L763"></a>switch s[2] {
            <a id="L764"></a>case &#39;\&#39;&#39;:
                <a id="L765"></a>val = &#39;\&#39;&#39;
            <a id="L766"></a>case &#39;&#34;&#39;:
                <a id="L767"></a>val = &#39;&#34;&#39;
            <a id="L768"></a>case &#39;\\&#39;:
                <a id="L769"></a>val = &#39;\\&#39;
            <a id="L770"></a>case &#39;a&#39;:
                <a id="L771"></a>val = &#39;\a&#39;
            <a id="L772"></a>case &#39;b&#39;:
                <a id="L773"></a>val = &#39;\b&#39;
            <a id="L774"></a>case &#39;n&#39;:
                <a id="L775"></a>val = &#39;\n&#39;
            <a id="L776"></a>case &#39;r&#39;:
                <a id="L777"></a>val = &#39;\r&#39;
            <a id="L778"></a>case &#39;t&#39;:
                <a id="L779"></a>val = &#39;\t&#39;
            <a id="L780"></a>case &#39;v&#39;:
                <a id="L781"></a>val = &#39;\v&#39;
            <a id="L782"></a>default:
                <a id="L783"></a>error(&#34;invalid escape %v&#34;, s[1:3])
            <a id="L784"></a>}
        <a id="L785"></a>} else if s[2] == &#39;u&#39; &amp;&amp; len(s) == 2+1+4 { <span class="comment">// \unnnn sequence</span>
            <a id="L786"></a>val = 0;
            <a id="L787"></a>s = s[3:len(s)];
            <a id="L788"></a>for s != &#34;&#34; {
                <a id="L789"></a>c := int(s[0]);
                <a id="L790"></a>switch {
                <a id="L791"></a>case c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;:
                    <a id="L792"></a>c -= &#39;0&#39;
                <a id="L793"></a>case c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;f&#39;:
                    <a id="L794"></a>c -= &#39;a&#39; - 10
                <a id="L795"></a>case c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;F&#39;:
                    <a id="L796"></a>c -= &#39;A&#39; - 10
                <a id="L797"></a>default:
                    <a id="L798"></a>error(&#34;illegal \\unnnn construction&#34;)
                <a id="L799"></a>}
                <a id="L800"></a>val = val*16 + c;
                <a id="L801"></a>s = s[1:len(s)];
            <a id="L802"></a>}
            <a id="L803"></a>if val == 0 {
                <a id="L804"></a>error(&#34;&#39;\\u0000&#39; is illegal&#34;)
            <a id="L805"></a>}
        <a id="L806"></a>} else {
            <a id="L807"></a>error(&#34;unknown escape&#34;)
        <a id="L808"></a>}
    <a id="L809"></a>} else {
        <a id="L810"></a>val = extval;
        <a id="L811"></a>extval++;
    <a id="L812"></a>}

    <a id="L814"></a>tokset[ntokens].value = val;
    <a id="L815"></a>return ntokens;
<a id="L816"></a>}

<a id="L818"></a>var peekline = 0

<a id="L820"></a>func gettok() int {
    <a id="L821"></a>var i, match, c int;

    <a id="L823"></a>tokname = &#34;&#34;;
    <a id="L824"></a>for {
        <a id="L825"></a>lineno += peekline;
        <a id="L826"></a>peekline = 0;
        <a id="L827"></a>c = getrune(finput);
        <a id="L828"></a>for c == &#39; &#39; || c == &#39;\n&#39; || c == &#39;\t&#39; || c == &#39;\v&#39; || c == &#39;\r&#39; {
            <a id="L829"></a>if c == &#39;\n&#39; {
                <a id="L830"></a>lineno++
            <a id="L831"></a>}
            <a id="L832"></a>c = getrune(finput);
        <a id="L833"></a>}

        <a id="L835"></a><span class="comment">// skip comment -- fix</span>
        <a id="L836"></a>if c != &#39;/&#39; {
            <a id="L837"></a>break
        <a id="L838"></a>}
        <a id="L839"></a>lineno += skipcom();
    <a id="L840"></a>}

    <a id="L842"></a>switch c {
    <a id="L843"></a>case EOF:
        <a id="L844"></a>if tokflag {
            <a id="L845"></a>fmt.Printf(&#34;&gt;&gt;&gt; ENDFILE %v\n&#34;, lineno)
        <a id="L846"></a>}
        <a id="L847"></a>return ENDFILE;

    <a id="L849"></a>case &#39;{&#39;:
        <a id="L850"></a>ungetrune(finput, c);
        <a id="L851"></a>if tokflag {
            <a id="L852"></a>fmt.Printf(&#34;&gt;&gt;&gt; ={ %v\n&#34;, lineno)
        <a id="L853"></a>}
        <a id="L854"></a>return &#39;=&#39;;

    <a id="L856"></a>case &#39;&lt;&#39;:
        <a id="L857"></a><span class="comment">// get, and look up, a type name (union member name)</span>
        <a id="L858"></a>c = getrune(finput);
        <a id="L859"></a>for c != &#39;&gt;&#39; &amp;&amp; c != EOF &amp;&amp; c != &#39;\n&#39; {
            <a id="L860"></a>tokname += string(c);
            <a id="L861"></a>c = getrune(finput);
        <a id="L862"></a>}

        <a id="L864"></a>if c != &#39;&gt;&#39; {
            <a id="L865"></a>error(&#34;unterminated &lt; ... &gt; clause&#34;)
        <a id="L866"></a>}

        <a id="L868"></a>for i = 1; i &lt;= ntypes; i++ {
            <a id="L869"></a>if typeset[i] == tokname {
                <a id="L870"></a>numbval = i;
                <a id="L871"></a>if tokflag {
                    <a id="L872"></a>fmt.Printf(&#34;&gt;&gt;&gt; TYPENAME old &lt;%v&gt; %v\n&#34;, tokname, lineno)
                <a id="L873"></a>}
                <a id="L874"></a>return TYPENAME;
            <a id="L875"></a>}
        <a id="L876"></a>}
        <a id="L877"></a>ntypes++;
        <a id="L878"></a>numbval = ntypes;
        <a id="L879"></a>typeset[numbval] = tokname;
        <a id="L880"></a>if tokflag {
            <a id="L881"></a>fmt.Printf(&#34;&gt;&gt;&gt; TYPENAME new &lt;%v&gt; %v\n&#34;, tokname, lineno)
        <a id="L882"></a>}
        <a id="L883"></a>return TYPENAME;

    <a id="L885"></a>case &#39;&#34;&#39;, &#39;\&#39;&#39;:
        <a id="L886"></a>match = c;
        <a id="L887"></a>tokname = &#34; &#34;;
        <a id="L888"></a>for {
            <a id="L889"></a>c = getrune(finput);
            <a id="L890"></a>if c == &#39;\n&#39; || c == EOF {
                <a id="L891"></a>error(&#34;illegal or missing &#39; or \&#34;&#34;)
            <a id="L892"></a>}
            <a id="L893"></a>if c == &#39;\\&#39; {
                <a id="L894"></a>tokname += string(&#39;\\&#39;);
                <a id="L895"></a>c = getrune(finput);
            <a id="L896"></a>} else if c == match {
                <a id="L897"></a>if tokflag {
                    <a id="L898"></a>fmt.Printf(&#34;&gt;&gt;&gt; IDENTIFIER \&#34;%v\&#34; %v\n&#34;, tokname, lineno)
                <a id="L899"></a>}
                <a id="L900"></a>return IDENTIFIER;
            <a id="L901"></a>}
            <a id="L902"></a>tokname += string(c);
        <a id="L903"></a>}

    <a id="L905"></a>case &#39;%&#39;:
        <a id="L906"></a>c = getrune(finput);
        <a id="L907"></a>switch c {
        <a id="L908"></a>case &#39;%&#39;:
            <a id="L909"></a>if tokflag {
                <a id="L910"></a>fmt.Printf(&#34;&gt;&gt;&gt; MARK %%%% %v\n&#34;, lineno)
            <a id="L911"></a>}
            <a id="L912"></a>return MARK;
        <a id="L913"></a>case &#39;=&#39;:
            <a id="L914"></a>if tokflag {
                <a id="L915"></a>fmt.Printf(&#34;&gt;&gt;&gt; PREC %%= %v\n&#34;, lineno)
            <a id="L916"></a>}
            <a id="L917"></a>return PREC;
        <a id="L918"></a>case &#39;{&#39;:
            <a id="L919"></a>if tokflag {
                <a id="L920"></a>fmt.Printf(&#34;&gt;&gt;&gt; LCURLY %%{ %v\n&#34;, lineno)
            <a id="L921"></a>}
            <a id="L922"></a>return LCURLY;
        <a id="L923"></a>}

        <a id="L925"></a>getword(c);
        <a id="L926"></a><span class="comment">// find a reserved word</span>
        <a id="L927"></a>for c = 0; c &lt; len(resrv); c++ {
            <a id="L928"></a>if tokname == resrv[c].name {
                <a id="L929"></a>if tokflag {
                    <a id="L930"></a>fmt.Printf(&#34;&gt;&gt;&gt; %%%v %v %v\n&#34;, tokname,
                        <a id="L931"></a>resrv[c].value-PRIVATE, lineno)
                <a id="L932"></a>}
                <a id="L933"></a>return resrv[c].value;
            <a id="L934"></a>}
        <a id="L935"></a>}
        <a id="L936"></a>error(&#34;invalid escape, or illegal reserved word: %v&#34;, tokname);

    <a id="L938"></a>case &#39;0&#39;, &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;, &#39;8&#39;, &#39;9&#39;:
        <a id="L939"></a>numbval = c - &#39;0&#39;;
        <a id="L940"></a>for {
            <a id="L941"></a>c = getrune(finput);
            <a id="L942"></a>if !isdigit(c) {
                <a id="L943"></a>break
            <a id="L944"></a>}
            <a id="L945"></a>numbval = numbval*10 + c - &#39;0&#39;;
        <a id="L946"></a>}
        <a id="L947"></a>ungetrune(finput, c);
        <a id="L948"></a>if tokflag {
            <a id="L949"></a>fmt.Printf(&#34;&gt;&gt;&gt; NUMBER %v %v\n&#34;, numbval, lineno)
        <a id="L950"></a>}
        <a id="L951"></a>return NUMBER;

    <a id="L953"></a>default:
        <a id="L954"></a>if isword(c) || c == &#39;.&#39; || c == &#39;$&#39; {
            <a id="L955"></a>getword(c);
            <a id="L956"></a>break;
        <a id="L957"></a>}
        <a id="L958"></a>if tokflag {
            <a id="L959"></a>fmt.Printf(&#34;&gt;&gt;&gt; OPERATOR %v %v\n&#34;, string(c), lineno)
        <a id="L960"></a>}
        <a id="L961"></a>return c;
    <a id="L962"></a>}

    <a id="L964"></a><span class="comment">// look ahead to distinguish IDENTIFIER from IDENTCOLON</span>
    <a id="L965"></a>c = getrune(finput);
    <a id="L966"></a>for c == &#39; &#39; || c == &#39;\t&#39; || c == &#39;\n&#39; || c == &#39;\v&#39; || c == &#39;\r&#39; || c == &#39;/&#39; {
        <a id="L967"></a>if c == &#39;\n&#39; {
            <a id="L968"></a>peekline++
        <a id="L969"></a>}
        <a id="L970"></a><span class="comment">// look for comments</span>
        <a id="L971"></a>if c == &#39;/&#39; {
            <a id="L972"></a>peekline += skipcom()
        <a id="L973"></a>}
        <a id="L974"></a>c = getrune(finput);
    <a id="L975"></a>}
    <a id="L976"></a>if c == &#39;:&#39; {
        <a id="L977"></a>if tokflag {
            <a id="L978"></a>fmt.Printf(&#34;&gt;&gt;&gt; IDENTCOLON %v: %v\n&#34;, tokname, lineno)
        <a id="L979"></a>}
        <a id="L980"></a>return IDENTCOLON;
    <a id="L981"></a>}

    <a id="L983"></a>ungetrune(finput, c);
    <a id="L984"></a>if tokflag {
        <a id="L985"></a>fmt.Printf(&#34;&gt;&gt;&gt; IDENTIFIER %v %v\n&#34;, tokname, lineno)
    <a id="L986"></a>}
    <a id="L987"></a>return IDENTIFIER;
<a id="L988"></a>}

<a id="L990"></a>func getword(c int) {
    <a id="L991"></a>tokname = &#34;&#34;;
    <a id="L992"></a>for isword(c) || isdigit(c) || c == &#39;_&#39; || c == &#39;.&#39; || c == &#39;$&#39; {
        <a id="L993"></a>tokname += string(c);
        <a id="L994"></a>c = getrune(finput);
    <a id="L995"></a>}
    <a id="L996"></a>ungetrune(finput, c);
<a id="L997"></a>}

<a id="L999"></a><span class="comment">//</span>
<a id="L1000"></a><span class="comment">// determine the type of a symbol</span>
<a id="L1001"></a><span class="comment">//</span>
<a id="L1002"></a>func fdtype(t int) int {
    <a id="L1003"></a>var v int;
    <a id="L1004"></a>var s string;

    <a id="L1006"></a>if t &gt;= NTBASE {
        <a id="L1007"></a>v = nontrst[t-NTBASE].value;
        <a id="L1008"></a>s = nontrst[t-NTBASE].name;
    <a id="L1009"></a>} else {
        <a id="L1010"></a>v = TYPE(toklev[t]);
        <a id="L1011"></a>s = tokset[t].name;
    <a id="L1012"></a>}
    <a id="L1013"></a>if v &lt;= 0 {
        <a id="L1014"></a>error(&#34;must specify type for %v&#34;, s)
    <a id="L1015"></a>}
    <a id="L1016"></a>return v;
<a id="L1017"></a>}

<a id="L1019"></a>func chfind(t int, s string) int {
    <a id="L1020"></a>if s[0] == &#39; &#39; {
        <a id="L1021"></a>t = 0
    <a id="L1022"></a>}
    <a id="L1023"></a>for i := 0; i &lt;= ntokens; i++ {
        <a id="L1024"></a>if s == tokset[i].name {
            <a id="L1025"></a>return i
        <a id="L1026"></a>}
    <a id="L1027"></a>}
    <a id="L1028"></a>for i := 0; i &lt;= nnonter; i++ {
        <a id="L1029"></a>if s == nontrst[i].name {
            <a id="L1030"></a>return NTBASE + i
        <a id="L1031"></a>}
    <a id="L1032"></a>}

    <a id="L1034"></a><span class="comment">// cannot find name</span>
    <a id="L1035"></a>if t &gt; 1 {
        <a id="L1036"></a>error(&#34;%v should have been defined earlier&#34;, s)
    <a id="L1037"></a>}
    <a id="L1038"></a>return defin(t, s);
<a id="L1039"></a>}

<a id="L1041"></a><span class="comment">//</span>
<a id="L1042"></a><span class="comment">// copy the union declaration to the output, and the define file if present</span>
<a id="L1043"></a><span class="comment">//</span>
<a id="L1044"></a>func cpyunion() {

    <a id="L1046"></a>if !lflag {
        <a id="L1047"></a>fmt.Fprintf(ftable, &#34;\n//line %v %v\n&#34;, lineno, infile)
    <a id="L1048"></a>}
    <a id="L1049"></a>fmt.Fprintf(ftable, &#34;type\tYYSTYPE\tstruct&#34;);

    <a id="L1051"></a>level := 0;

<a id="L1053"></a>out:
    <a id="L1054"></a>for {
        <a id="L1055"></a>c := getrune(finput);
        <a id="L1056"></a>if c == EOF {
            <a id="L1057"></a>error(&#34;EOF encountered while processing %%union&#34;)
        <a id="L1058"></a>}
        <a id="L1059"></a>putrune(ftable, c);
        <a id="L1060"></a>switch c {
        <a id="L1061"></a>case &#39;\n&#39;:
            <a id="L1062"></a>lineno++
        <a id="L1063"></a>case &#39;{&#39;:
            <a id="L1064"></a>if level == 0 {
                <a id="L1065"></a>fmt.Fprintf(ftable, &#34;\n\tyys\tint;&#34;)
            <a id="L1066"></a>}
            <a id="L1067"></a>level++;
        <a id="L1068"></a>case &#39;}&#39;:
            <a id="L1069"></a>level--;
            <a id="L1070"></a>if level == 0 {
                <a id="L1071"></a>break out
            <a id="L1072"></a>}
        <a id="L1073"></a>}
    <a id="L1074"></a>}
    <a id="L1075"></a>fmt.Fprintf(ftable, &#34;\n&#34;);
    <a id="L1076"></a>fmt.Fprintf(ftable, &#34;var\tyylval\tYYSTYPE\n&#34;);
    <a id="L1077"></a>fmt.Fprintf(ftable, &#34;var\tYYVAL\tYYSTYPE\n&#34;);
    <a id="L1078"></a>fmt.Fprintf(ftable, &#34;var\tYYS\t[%v]YYSTYPE\n&#34;, stacksize);
<a id="L1079"></a>}

<a id="L1081"></a><span class="comment">//</span>
<a id="L1082"></a><span class="comment">// saves code between %{ and %}</span>
<a id="L1083"></a><span class="comment">//</span>
<a id="L1084"></a>func cpycode() {
    <a id="L1085"></a>lno := lineno;

    <a id="L1087"></a>c := getrune(finput);
    <a id="L1088"></a>if c == &#39;\n&#39; {
        <a id="L1089"></a>c = getrune(finput);
        <a id="L1090"></a>lineno++;
    <a id="L1091"></a>}
    <a id="L1092"></a>if !lflag {
        <a id="L1093"></a>fmt.Fprintf(ftable, &#34;\n//line %v %v\n&#34;, lineno, infile)
    <a id="L1094"></a>}
    <a id="L1095"></a>for c != EOF {
        <a id="L1096"></a>if c == &#39;%&#39; {
            <a id="L1097"></a>c = getrune(finput);
            <a id="L1098"></a>if c == &#39;}&#39; {
                <a id="L1099"></a>return
            <a id="L1100"></a>}
            <a id="L1101"></a>putrune(ftable, &#39;%&#39;);
        <a id="L1102"></a>}
        <a id="L1103"></a>putrune(ftable, c);
        <a id="L1104"></a>if c == &#39;\n&#39; {
            <a id="L1105"></a>lineno++
        <a id="L1106"></a>}
        <a id="L1107"></a>c = getrune(finput);
    <a id="L1108"></a>}
    <a id="L1109"></a>lineno = lno;
    <a id="L1110"></a>error(&#34;eof before %%}&#34;);
<a id="L1111"></a>}

<a id="L1113"></a><span class="comment">//func</span>
<a id="L1114"></a><span class="comment">//addcode(k int, s string)</span>
<a id="L1115"></a><span class="comment">//{</span>
<a id="L1116"></a><span class="comment">//	for i := 0; i &lt; len(s); i++ {</span>
<a id="L1117"></a><span class="comment">//		addcodec(k, int(s[i]));</span>
<a id="L1118"></a><span class="comment">//	}</span>
<a id="L1119"></a><span class="comment">//}</span>

<a id="L1121"></a><span class="comment">//func</span>
<a id="L1122"></a><span class="comment">//addcodec(k, c int)</span>
<a id="L1123"></a><span class="comment">//{</span>
<a id="L1124"></a><span class="comment">//	if codehead == nil || k != codetail.kind || codetail.ndata &gt;= NCode {</span>
<a id="L1125"></a><span class="comment">//		cd := new(Code);</span>
<a id="L1126"></a><span class="comment">//		cd.kind = k;</span>
<a id="L1127"></a><span class="comment">//		cd.data = make([]byte, NCode+UTFmax);</span>
<a id="L1128"></a><span class="comment">//		cd.ndata = 0;</span>
<a id="L1129"></a><span class="comment">//		cd.next = nil;</span>
<a id="L1130"></a><span class="comment">//</span>
<a id="L1131"></a><span class="comment">//		if codehead == nil {</span>
<a id="L1132"></a><span class="comment">//			codehead = cd;</span>
<a id="L1133"></a><span class="comment">//		} else</span>
<a id="L1134"></a><span class="comment">//			codetail.next = cd;</span>
<a id="L1135"></a><span class="comment">//		codetail = cd;</span>
<a id="L1136"></a><span class="comment">//	}</span>
<a id="L1137"></a><span class="comment">//</span>
<a id="L1138"></a><span class="comment">////!!	codetail.ndata += sys-&gt;char2byte(c, codetail.data, codetail.ndata);</span>
<a id="L1139"></a><span class="comment">//}</span>

<a id="L1141"></a><span class="comment">//func</span>
<a id="L1142"></a><span class="comment">//dumpcode(til int)</span>
<a id="L1143"></a><span class="comment">//{</span>
<a id="L1144"></a><span class="comment">//	for ; codehead != nil; codehead = codehead.next {</span>
<a id="L1145"></a><span class="comment">//		if codehead.kind == til {</span>
<a id="L1146"></a><span class="comment">//			return;</span>
<a id="L1147"></a><span class="comment">//		}</span>
<a id="L1148"></a><span class="comment">//		if write(ftable, codehead.data, codehead.ndata) != codehead.ndata {</span>
<a id="L1149"></a><span class="comment">//			error(&#34;can&#39;t write output file&#34;);</span>
<a id="L1150"></a><span class="comment">//		}</span>
<a id="L1151"></a><span class="comment">//	}</span>
<a id="L1152"></a><span class="comment">//}</span>

<a id="L1154"></a><span class="comment">//</span>
<a id="L1155"></a><span class="comment">// write out the module declaration and any token info</span>
<a id="L1156"></a><span class="comment">//</span>
<a id="L1157"></a><span class="comment">//func</span>
<a id="L1158"></a><span class="comment">//dumpmod()</span>
<a id="L1159"></a><span class="comment">//{</span>
<a id="L1160"></a><span class="comment">//</span>
<a id="L1161"></a><span class="comment">//	for ; codehead != nil; codehead = codehead.next {</span>
<a id="L1162"></a><span class="comment">//		if codehead.kind != CodeMod {</span>
<a id="L1163"></a><span class="comment">//			break;</span>
<a id="L1164"></a><span class="comment">//		}</span>
<a id="L1165"></a><span class="comment">//		if write(ftable, codehead.data, codehead.ndata) != codehead.ndata {</span>
<a id="L1166"></a><span class="comment">//			error(&#34;can&#39;t write output file&#34;);</span>
<a id="L1167"></a><span class="comment">//		}</span>
<a id="L1168"></a><span class="comment">//	}</span>
<a id="L1169"></a><span class="comment">//</span>
<a id="L1170"></a><span class="comment">//	for i:=TOKSTART; i&lt;=ntokens; i++ {</span>
<a id="L1171"></a><span class="comment">//		// non-literals</span>
<a id="L1172"></a><span class="comment">//		c := tokset[i].name[0];</span>
<a id="L1173"></a><span class="comment">//		if c != &#39; &#39; &amp;&amp; c != &#39;$&#39; {</span>
<a id="L1174"></a><span class="comment">//			fmt.Fprintf(ftable, &#34;vonst	%v	%v\n&#34;,</span>
<a id="L1175"></a><span class="comment">//				tokset[i].name, tokset[i].value);</span>
<a id="L1176"></a><span class="comment">//		}</span>
<a id="L1177"></a><span class="comment">//	}</span>
<a id="L1178"></a><span class="comment">//</span>
<a id="L1179"></a><span class="comment">//}</span>

<a id="L1181"></a><span class="comment">//</span>
<a id="L1182"></a><span class="comment">// skip over comments</span>
<a id="L1183"></a><span class="comment">// skipcom is called after reading a &#39;/&#39;</span>
<a id="L1184"></a><span class="comment">//</span>
<a id="L1185"></a>func skipcom() int {
    <a id="L1186"></a>var c int;

    <a id="L1188"></a>c = getrune(finput);
    <a id="L1189"></a>if c == &#39;/&#39; {
        <a id="L1190"></a>for c != EOF {
            <a id="L1191"></a>if c == &#39;\n&#39; {
                <a id="L1192"></a>return 1
            <a id="L1193"></a>}
            <a id="L1194"></a>c = getrune(finput);
        <a id="L1195"></a>}
        <a id="L1196"></a>error(&#34;EOF inside comment&#34;);
        <a id="L1197"></a>return 0;
    <a id="L1198"></a>}
    <a id="L1199"></a>if c != &#39;*&#39; {
        <a id="L1200"></a>error(&#34;illegal comment&#34;)
    <a id="L1201"></a>}

    <a id="L1203"></a>nl := 0; <span class="comment">// lines skipped</span>
    <a id="L1204"></a>c = getrune(finput);

<a id="L1206"></a>l1:
    <a id="L1207"></a>switch c {
    <a id="L1208"></a>case &#39;*&#39;:
        <a id="L1209"></a>c = getrune(finput);
        <a id="L1210"></a>if c == &#39;/&#39; {
            <a id="L1211"></a>break
        <a id="L1212"></a>}
        <a id="L1213"></a>goto l1;

    <a id="L1215"></a>case &#39;\n&#39;:
        <a id="L1216"></a>nl++;
        <a id="L1217"></a>fallthrough;

    <a id="L1219"></a>default:
        <a id="L1220"></a>c = getrune(finput);
        <a id="L1221"></a>goto l1;
    <a id="L1222"></a>}
    <a id="L1223"></a>return nl;
<a id="L1224"></a>}

<a id="L1226"></a>func dumpprod(curprod []int, max int) {
    <a id="L1227"></a>fmt.Printf(&#34;\n&#34;);
    <a id="L1228"></a>for i := 0; i &lt; max; i++ {
        <a id="L1229"></a>p := curprod[i];
        <a id="L1230"></a>if p &lt; 0 {
            <a id="L1231"></a>fmt.Printf(&#34;[%v] %v\n&#34;, i, p)
        <a id="L1232"></a>} else {
            <a id="L1233"></a>fmt.Printf(&#34;[%v] %v\n&#34;, i, symnam(p))
        <a id="L1234"></a>}
    <a id="L1235"></a>}
<a id="L1236"></a>}

<a id="L1238"></a><span class="comment">//</span>
<a id="L1239"></a><span class="comment">// copy action to the next ; or closing }</span>
<a id="L1240"></a><span class="comment">//</span>
<a id="L1241"></a>func cpyact(curprod []int, max int) {

    <a id="L1243"></a>if !lflag {
        <a id="L1244"></a>fmt.Fprintf(ftable, &#34;\n//line %v %v\n&#34;, lineno, infile)
    <a id="L1245"></a>}

    <a id="L1247"></a>lno := lineno;
    <a id="L1248"></a>brac := 0;

<a id="L1250"></a>loop:
    <a id="L1251"></a>for {
        <a id="L1252"></a>c := getrune(finput);

    <a id="L1254"></a>swt:
        <a id="L1255"></a>switch c {
        <a id="L1256"></a>case &#39;;&#39;:
            <a id="L1257"></a>if brac == 0 {
                <a id="L1258"></a>putrune(ftable, c);
                <a id="L1259"></a>return;
            <a id="L1260"></a>}

        <a id="L1262"></a>case &#39;{&#39;:
            <a id="L1263"></a>if brac == 0 {
            <a id="L1264"></a>}
            <a id="L1265"></a>putrune(ftable, &#39;\t&#39;);
            <a id="L1266"></a>brac++;

        <a id="L1268"></a>case &#39;$&#39;:
            <a id="L1269"></a>s := 1;
            <a id="L1270"></a>tok := -1;
            <a id="L1271"></a>c = getrune(finput);

            <a id="L1273"></a><span class="comment">// type description</span>
            <a id="L1274"></a>if c == &#39;&lt;&#39; {
                <a id="L1275"></a>ungetrune(finput, c);
                <a id="L1276"></a>if gettok() != TYPENAME {
                    <a id="L1277"></a>error(&#34;bad syntax on $&lt;ident&gt; clause&#34;)
                <a id="L1278"></a>}
                <a id="L1279"></a>tok = numbval;
                <a id="L1280"></a>c = getrune(finput);
            <a id="L1281"></a>}
            <a id="L1282"></a>if c == &#39;$&#39; {
                <a id="L1283"></a>fmt.Fprintf(ftable, &#34;YYVAL&#34;);

                <a id="L1285"></a><span class="comment">// put out the proper tag...</span>
                <a id="L1286"></a>if ntypes != 0 {
                    <a id="L1287"></a>if tok &lt; 0 {
                        <a id="L1288"></a>tok = fdtype(curprod[0])
                    <a id="L1289"></a>}
                    <a id="L1290"></a>fmt.Fprintf(ftable, &#34;.%v&#34;, typeset[tok]);
                <a id="L1291"></a>}
                <a id="L1292"></a>continue loop;
            <a id="L1293"></a>}
            <a id="L1294"></a>if c == &#39;-&#39; {
                <a id="L1295"></a>s = -s;
                <a id="L1296"></a>c = getrune(finput);
            <a id="L1297"></a>}
            <a id="L1298"></a>j := 0;
            <a id="L1299"></a>if isdigit(c) {
                <a id="L1300"></a>for isdigit(c) {
                    <a id="L1301"></a>j = j*10 + c - &#39;0&#39;;
                    <a id="L1302"></a>c = getrune(finput);
                <a id="L1303"></a>}
                <a id="L1304"></a>ungetrune(finput, c);
                <a id="L1305"></a>j = j * s;
                <a id="L1306"></a>if j &gt;= max {
                    <a id="L1307"></a>error(&#34;Illegal use of $%v&#34;, j)
                <a id="L1308"></a>}
            <a id="L1309"></a>} else if isword(c) || c == &#39;_&#39; || c == &#39;.&#39; {
                <a id="L1310"></a><span class="comment">// look for $name</span>
                <a id="L1311"></a>ungetrune(finput, c);
                <a id="L1312"></a>if gettok() != IDENTIFIER {
                    <a id="L1313"></a>error(&#34;$ must be followed by an identifier&#34;)
                <a id="L1314"></a>}
                <a id="L1315"></a>tokn := chfind(2, tokname);
                <a id="L1316"></a>fnd := -1;
                <a id="L1317"></a>c = getrune(finput);
                <a id="L1318"></a>if c != &#39;@&#39; {
                    <a id="L1319"></a>ungetrune(finput, c)
                <a id="L1320"></a>} else if gettok() != NUMBER {
                    <a id="L1321"></a>error(&#34;@ must be followed by number&#34;)
                <a id="L1322"></a>} else {
                    <a id="L1323"></a>fnd = numbval
                <a id="L1324"></a>}
                <a id="L1325"></a>for j = 1; j &lt; max; j++ {
                    <a id="L1326"></a>if tokn == curprod[j] {
                        <a id="L1327"></a>fnd--;
                        <a id="L1328"></a>if fnd &lt;= 0 {
                            <a id="L1329"></a>break
                        <a id="L1330"></a>}
                    <a id="L1331"></a>}
                <a id="L1332"></a>}
                <a id="L1333"></a>if j &gt;= max {
                    <a id="L1334"></a>error(&#34;$name or $name@number not found&#34;)
                <a id="L1335"></a>}
            <a id="L1336"></a>} else {
                <a id="L1337"></a>putrune(ftable, &#39;$&#39;);
                <a id="L1338"></a>if s &lt; 0 {
                    <a id="L1339"></a>putrune(ftable, &#39;-&#39;)
                <a id="L1340"></a>}
                <a id="L1341"></a>ungetrune(finput, c);
                <a id="L1342"></a>continue loop;
            <a id="L1343"></a>}
            <a id="L1344"></a>fmt.Fprintf(ftable, &#34;YYS[yypt-%v]&#34;, max-j-1);

            <a id="L1346"></a><span class="comment">// put out the proper tag</span>
            <a id="L1347"></a>if ntypes != 0 {
                <a id="L1348"></a>if j &lt;= 0 &amp;&amp; tok &lt; 0 {
                    <a id="L1349"></a>error(&#34;must specify type of $%v&#34;, j)
                <a id="L1350"></a>}
                <a id="L1351"></a>if tok &lt; 0 {
                    <a id="L1352"></a>tok = fdtype(curprod[j])
                <a id="L1353"></a>}
                <a id="L1354"></a>fmt.Fprintf(ftable, &#34;.%v&#34;, typeset[tok]);
            <a id="L1355"></a>}
            <a id="L1356"></a>continue loop;

        <a id="L1358"></a>case &#39;}&#39;:
            <a id="L1359"></a>brac--;
            <a id="L1360"></a>if brac != 0 {
                <a id="L1361"></a>break
            <a id="L1362"></a>}
            <a id="L1363"></a>putrune(ftable, c);
            <a id="L1364"></a>return;

        <a id="L1366"></a>case &#39;/&#39;:
            <a id="L1367"></a><span class="comment">// a comment</span>
            <a id="L1368"></a>putrune(ftable, c);
            <a id="L1369"></a>c = getrune(finput);
            <a id="L1370"></a>for c != EOF {
                <a id="L1371"></a>if c == &#39;\n&#39; {
                    <a id="L1372"></a>lineno++;
                    <a id="L1373"></a>break swt;
                <a id="L1374"></a>}
                <a id="L1375"></a>putrune(ftable, c);
                <a id="L1376"></a>c = getrune(finput);
            <a id="L1377"></a>}
            <a id="L1378"></a>error(&#34;EOF inside comment&#34;);

        <a id="L1380"></a>case &#39;\&#39;&#39;, &#39;&#34;&#39;:
            <a id="L1381"></a><span class="comment">// character string or constant</span>
            <a id="L1382"></a>match := c;
            <a id="L1383"></a>putrune(ftable, c);
            <a id="L1384"></a>c = getrune(finput);
            <a id="L1385"></a>for c != EOF {
                <a id="L1386"></a>if c == &#39;\\&#39; {
                    <a id="L1387"></a>putrune(ftable, c);
                    <a id="L1388"></a>c = getrune(finput);
                    <a id="L1389"></a>if c == &#39;\n&#39; {
                        <a id="L1390"></a>lineno++
                    <a id="L1391"></a>}
                <a id="L1392"></a>} else if c == match {
                    <a id="L1393"></a>break swt
                <a id="L1394"></a>}
                <a id="L1395"></a>if c == &#39;\n&#39; {
                    <a id="L1396"></a>error(&#34;newline in string or char const&#34;)
                <a id="L1397"></a>}
                <a id="L1398"></a>putrune(ftable, c);
                <a id="L1399"></a>c = getrune(finput);
            <a id="L1400"></a>}
            <a id="L1401"></a>error(&#34;EOF in string or character constant&#34;);

        <a id="L1403"></a>case EOF:
            <a id="L1404"></a>lineno = lno;
            <a id="L1405"></a>error(&#34;action does not terminate&#34;);

        <a id="L1407"></a>case &#39;\n&#39;:
            <a id="L1408"></a>lineno++
        <a id="L1409"></a>}

        <a id="L1411"></a>putrune(ftable, c);
    <a id="L1412"></a>}
<a id="L1413"></a>}

<a id="L1415"></a>func openup() {
    <a id="L1416"></a>infile = flag.Arg(0);
    <a id="L1417"></a>finput = open(infile);
    <a id="L1418"></a>if finput == nil {
        <a id="L1419"></a>error(&#34;cannot open %v&#34;, infile)
    <a id="L1420"></a>}

    <a id="L1422"></a>foutput = nil;
    <a id="L1423"></a>if vflag != &#34;&#34; {
        <a id="L1424"></a>foutput = create(vflag, 0666);
        <a id="L1425"></a>if foutput == nil {
            <a id="L1426"></a>error(&#34;can&#39;t create file %v&#34;, vflag)
        <a id="L1427"></a>}
    <a id="L1428"></a>}

    <a id="L1430"></a>ftable = nil;
    <a id="L1431"></a>if oflag == &#34;&#34; {
        <a id="L1432"></a>oflag = &#34;y.go&#34;
    <a id="L1433"></a>}
    <a id="L1434"></a>ftable = create(oflag, 0666);
    <a id="L1435"></a>if ftable == nil {
        <a id="L1436"></a>error(&#34;can&#39;t create file %v&#34;, oflag)
    <a id="L1437"></a>}

<a id="L1439"></a>}

<a id="L1441"></a><span class="comment">//</span>
<a id="L1442"></a><span class="comment">// return a pointer to the name of symbol i</span>
<a id="L1443"></a><span class="comment">//</span>
<a id="L1444"></a>func symnam(i int) string {
    <a id="L1445"></a>var s string;

    <a id="L1447"></a>if i &gt;= NTBASE {
        <a id="L1448"></a>s = nontrst[i-NTBASE].name
    <a id="L1449"></a>} else {
        <a id="L1450"></a>s = tokset[i].name
    <a id="L1451"></a>}
    <a id="L1452"></a>if s[0] == &#39; &#39; {
        <a id="L1453"></a>s = s[1:len(s)]
    <a id="L1454"></a>}
    <a id="L1455"></a>return s;
<a id="L1456"></a>}

<a id="L1458"></a><span class="comment">//</span>
<a id="L1459"></a><span class="comment">// set elements 0 through n-1 to c</span>
<a id="L1460"></a><span class="comment">//</span>
<a id="L1461"></a>func aryfil(v []int, n, c int) {
    <a id="L1462"></a>for i := 0; i &lt; n; i++ {
        <a id="L1463"></a>v[i] = c
    <a id="L1464"></a>}
<a id="L1465"></a>}

<a id="L1467"></a><span class="comment">//</span>
<a id="L1468"></a><span class="comment">// compute an array with the beginnings of productions yielding given nonterminals</span>
<a id="L1469"></a><span class="comment">// The array pres points to these lists</span>
<a id="L1470"></a><span class="comment">// the array pyield has the lists: the total size is only NPROD+1</span>
<a id="L1471"></a><span class="comment">//</span>
<a id="L1472"></a>func cpres() {
    <a id="L1473"></a>pres = make([][][]int, nnonter+1);
    <a id="L1474"></a>curres := make([][]int, nprod);

    <a id="L1476"></a>if false {
        <a id="L1477"></a>for j := 0; j &lt;= nnonter; j++ {
            <a id="L1478"></a>fmt.Printf(&#34;nnonter[%v] = %v\n&#34;, j, nontrst[j].name)
        <a id="L1479"></a>}
        <a id="L1480"></a>for j := 0; j &lt; nprod; j++ {
            <a id="L1481"></a>fmt.Printf(&#34;prdptr[%v][0] = %v+NTBASE\n&#34;, j, prdptr[j][0]-NTBASE)
        <a id="L1482"></a>}
    <a id="L1483"></a>}

    <a id="L1485"></a>fatfl = 0; <span class="comment">// make undefined symbols nonfatal</span>
    <a id="L1486"></a>for i := 0; i &lt;= nnonter; i++ {
        <a id="L1487"></a>n := 0;
        <a id="L1488"></a>c := i + NTBASE;
        <a id="L1489"></a>for j := 0; j &lt; nprod; j++ {
            <a id="L1490"></a>if prdptr[j][0] == c {
                <a id="L1491"></a>curres[n] = prdptr[j][1:len(prdptr[j])];
                <a id="L1492"></a>n++;
            <a id="L1493"></a>}
        <a id="L1494"></a>}
        <a id="L1495"></a>if n == 0 {
            <a id="L1496"></a>error(&#34;nonterminal %v not defined&#34;, nontrst[i].name);
            <a id="L1497"></a>continue;
        <a id="L1498"></a>}
        <a id="L1499"></a>pres[i] = make([][]int, n);
        <a id="L1500"></a>for ll := 0; ll &lt; n; ll++ {
            <a id="L1501"></a>pres[i][ll] = curres[ll]
        <a id="L1502"></a>}
    <a id="L1503"></a>}
    <a id="L1504"></a>fatfl = 1;
    <a id="L1505"></a>if nerrors != 0 {
        <a id="L1506"></a>summary();
        <a id="L1507"></a>exit(1);
    <a id="L1508"></a>}
<a id="L1509"></a>}

<a id="L1511"></a>func dumppres() {
    <a id="L1512"></a>for i := 0; i &lt;= nnonter; i++ {
        <a id="L1513"></a>print(&#34;nonterm %d\n&#34;, i);
        <a id="L1514"></a>curres := pres[i];
        <a id="L1515"></a>for j := 0; j &lt; len(curres); j++ {
            <a id="L1516"></a>print(&#34;\tproduction %d:&#34;, j);
            <a id="L1517"></a>prd := curres[j];
            <a id="L1518"></a>for k := 0; k &lt; len(prd); k++ {
                <a id="L1519"></a>print(&#34; %d&#34;, prd[k])
            <a id="L1520"></a>}
            <a id="L1521"></a>print(&#34;\n&#34;);
        <a id="L1522"></a>}
    <a id="L1523"></a>}
<a id="L1524"></a>}

<a id="L1526"></a><span class="comment">//</span>
<a id="L1527"></a><span class="comment">// mark nonterminals which derive the empty string</span>
<a id="L1528"></a><span class="comment">// also, look for nonterminals which don&#39;t derive any token strings</span>
<a id="L1529"></a><span class="comment">//</span>
<a id="L1530"></a>func cempty() {
    <a id="L1531"></a>var i, p, np int;
    <a id="L1532"></a>var prd []int;

    <a id="L1534"></a>pempty = make([]int, nnonter+1);

    <a id="L1536"></a><span class="comment">// first, use the array pempty to detect productions that can never be reduced</span>
    <a id="L1537"></a><span class="comment">// set pempty to WHONOWS</span>
    <a id="L1538"></a>aryfil(pempty, nnonter+1, WHOKNOWS);

    <a id="L1540"></a><span class="comment">// now, look at productions, marking nonterminals which derive something</span>
<a id="L1541"></a>more:
    <a id="L1542"></a>for {
        <a id="L1543"></a>for i = 0; i &lt; nprod; i++ {
            <a id="L1544"></a>prd = prdptr[i];
            <a id="L1545"></a>if pempty[prd[0]-NTBASE] != 0 {
                <a id="L1546"></a>continue
            <a id="L1547"></a>}
            <a id="L1548"></a>np = len(prd) - 1;
            <a id="L1549"></a>for p = 1; p &lt; np; p++ {
                <a id="L1550"></a>if prd[p] &gt;= NTBASE &amp;&amp; pempty[prd[p]-NTBASE] == WHOKNOWS {
                    <a id="L1551"></a>break
                <a id="L1552"></a>}
            <a id="L1553"></a>}
            <a id="L1554"></a><span class="comment">// production can be derived</span>
            <a id="L1555"></a>if p == np {
                <a id="L1556"></a>pempty[prd[0]-NTBASE] = OK;
                <a id="L1557"></a>continue more;
            <a id="L1558"></a>}
        <a id="L1559"></a>}
        <a id="L1560"></a>break;
    <a id="L1561"></a>}

    <a id="L1563"></a><span class="comment">// now, look at the nonterminals, to see if they are all OK</span>
    <a id="L1564"></a>for i = 0; i &lt;= nnonter; i++ {
        <a id="L1565"></a><span class="comment">// the added production rises or falls as the start symbol ...</span>
        <a id="L1566"></a>if i == 0 {
            <a id="L1567"></a>continue
        <a id="L1568"></a>}
        <a id="L1569"></a>if pempty[i] != OK {
            <a id="L1570"></a>fatfl = 0;
            <a id="L1571"></a>error(&#34;nonterminal &#34; + nontrst[i].name + &#34; never derives any token string&#34;);
        <a id="L1572"></a>}
    <a id="L1573"></a>}

    <a id="L1575"></a>if nerrors != 0 {
        <a id="L1576"></a>summary();
        <a id="L1577"></a>exit(1);
    <a id="L1578"></a>}

    <a id="L1580"></a><span class="comment">// now, compute the pempty array, to see which nonterminals derive the empty string</span>
    <a id="L1581"></a><span class="comment">// set pempty to WHOKNOWS</span>
    <a id="L1582"></a>aryfil(pempty, nnonter+1, WHOKNOWS);

    <a id="L1584"></a><span class="comment">// loop as long as we keep finding empty nonterminals</span>

<a id="L1586"></a>again:
    <a id="L1587"></a>for {
    <a id="L1588"></a>next:
        <a id="L1589"></a>for i = 1; i &lt; nprod; i++ {
            <a id="L1590"></a><span class="comment">// not known to be empty</span>
            <a id="L1591"></a>prd = prdptr[i];
            <a id="L1592"></a>if pempty[prd[0]-NTBASE] != WHOKNOWS {
                <a id="L1593"></a>continue
            <a id="L1594"></a>}
            <a id="L1595"></a>np = len(prd) - 1;
            <a id="L1596"></a>for p = 1; p &lt; np; p++ {
                <a id="L1597"></a>if prd[p] &lt; NTBASE || pempty[prd[p]-NTBASE] != EMPTY {
                    <a id="L1598"></a>continue next
                <a id="L1599"></a>}
            <a id="L1600"></a>}

            <a id="L1602"></a><span class="comment">// we have a nontrivially empty nonterminal</span>
            <a id="L1603"></a>pempty[prd[0]-NTBASE] = EMPTY;

            <a id="L1605"></a><span class="comment">// got one ... try for another</span>
            <a id="L1606"></a>continue again;
        <a id="L1607"></a>}
        <a id="L1608"></a>return;
    <a id="L1609"></a>}
<a id="L1610"></a>}

<a id="L1612"></a>func dumpempty() {
    <a id="L1613"></a>for i := 0; i &lt;= nnonter; i++ {
        <a id="L1614"></a>if pempty[i] == EMPTY {
            <a id="L1615"></a>print(&#34;non-term %d %s matches empty\n&#34;, i, symnam(i+NTBASE))
        <a id="L1616"></a>}
    <a id="L1617"></a>}
<a id="L1618"></a>}

<a id="L1620"></a><span class="comment">//</span>
<a id="L1621"></a><span class="comment">// compute an array with the first of nonterminals</span>
<a id="L1622"></a><span class="comment">//</span>
<a id="L1623"></a>func cpfir() {
    <a id="L1624"></a>var s, n, p, np, ch, i int;
    <a id="L1625"></a>var curres [][]int;
    <a id="L1626"></a>var prd []int;

    <a id="L1628"></a>wsets = make([]Wset, nnonter+WSETINC);
    <a id="L1629"></a>pfirst = make([]Lkset, nnonter+1);
    <a id="L1630"></a>for i = 0; i &lt;= nnonter; i++ {
        <a id="L1631"></a>wsets[i].ws = mkset();
        <a id="L1632"></a>pfirst[i] = mkset();
        <a id="L1633"></a>curres = pres[i];
        <a id="L1634"></a>n = len(curres);

        <a id="L1636"></a><span class="comment">// initially fill the sets</span>
        <a id="L1637"></a>for s = 0; s &lt; n; s++ {
            <a id="L1638"></a>prd = curres[s];
            <a id="L1639"></a>np = len(prd) - 1;
            <a id="L1640"></a>for p = 0; p &lt; np; p++ {
                <a id="L1641"></a>ch = prd[p];
                <a id="L1642"></a>if ch &lt; NTBASE {
                    <a id="L1643"></a>setbit(pfirst[i], ch);
                    <a id="L1644"></a>break;
                <a id="L1645"></a>}
                <a id="L1646"></a>if pempty[ch-NTBASE] == 0 {
                    <a id="L1647"></a>break
                <a id="L1648"></a>}
            <a id="L1649"></a>}
        <a id="L1650"></a>}
    <a id="L1651"></a>}

    <a id="L1653"></a><span class="comment">// now, reflect transitivity</span>
    <a id="L1654"></a>changes := 1;
    <a id="L1655"></a>for changes != 0 {
        <a id="L1656"></a>changes = 0;
        <a id="L1657"></a>for i = 0; i &lt;= nnonter; i++ {
            <a id="L1658"></a>curres = pres[i];
            <a id="L1659"></a>n = len(curres);
            <a id="L1660"></a>for s = 0; s &lt; n; s++ {
                <a id="L1661"></a>prd = curres[s];
                <a id="L1662"></a>np = len(prd) - 1;
                <a id="L1663"></a>for p = 0; p &lt; np; p++ {
                    <a id="L1664"></a>ch = prd[p] - NTBASE;
                    <a id="L1665"></a>if ch &lt; 0 {
                        <a id="L1666"></a>break
                    <a id="L1667"></a>}
                    <a id="L1668"></a>changes |= setunion(pfirst[i], pfirst[ch]);
                    <a id="L1669"></a>if pempty[ch] == 0 {
                        <a id="L1670"></a>break
                    <a id="L1671"></a>}
                <a id="L1672"></a>}
            <a id="L1673"></a>}
        <a id="L1674"></a>}
    <a id="L1675"></a>}

    <a id="L1677"></a>if indebug == 0 {
        <a id="L1678"></a>return
    <a id="L1679"></a>}
    <a id="L1680"></a>if foutput != nil {
        <a id="L1681"></a>for i = 0; i &lt;= nnonter; i++ {
            <a id="L1682"></a>fmt.Fprintf(foutput, &#34;\n%v: %v %v\n&#34;,
                <a id="L1683"></a>nontrst[i].name, pfirst[i], pempty[i])
        <a id="L1684"></a>}
    <a id="L1685"></a>}
<a id="L1686"></a>}

<a id="L1688"></a><span class="comment">//</span>
<a id="L1689"></a><span class="comment">// generate the states</span>
<a id="L1690"></a><span class="comment">//</span>
<a id="L1691"></a>func stagen() {
    <a id="L1692"></a><span class="comment">// initialize</span>
    <a id="L1693"></a>nstate = 0;
    <a id="L1694"></a>tstates = make([]int, ntokens+1);  <span class="comment">// states generated by terminal gotos</span>
    <a id="L1695"></a>ntstates = make([]int, nnonter+1); <span class="comment">// states generated by nonterminal gotos</span>
    <a id="L1696"></a>amem = make([]int, ACTSIZE);
    <a id="L1697"></a>memp = 0;

    <a id="L1699"></a>clset = mkset();
    <a id="L1700"></a>pstate[0] = 0;
    <a id="L1701"></a>pstate[1] = 0;
    <a id="L1702"></a>aryfil(clset, tbitset, 0);
    <a id="L1703"></a>putitem(Pitem{prdptr[0], 0, 0, 0}, clset);
    <a id="L1704"></a>tystate[0] = MUSTDO;
    <a id="L1705"></a>nstate = 1;
    <a id="L1706"></a>pstate[2] = pstate[1];

    <a id="L1708"></a><span class="comment">//</span>
    <a id="L1709"></a><span class="comment">// now, the main state generation loop</span>
    <a id="L1710"></a><span class="comment">// first pass generates all of the states</span>
    <a id="L1711"></a><span class="comment">// later passes fix up lookahead</span>
    <a id="L1712"></a><span class="comment">// could be sped up a lot by remembering</span>
    <a id="L1713"></a><span class="comment">// results of the first pass rather than recomputing</span>
    <a id="L1714"></a><span class="comment">//</span>
    <a id="L1715"></a>first := 1;
    <a id="L1716"></a>for more := 1; more != 0; first = 0 {
        <a id="L1717"></a>more = 0;
        <a id="L1718"></a>for i := 0; i &lt; nstate; i++ {
            <a id="L1719"></a>if tystate[i] != MUSTDO {
                <a id="L1720"></a>continue
            <a id="L1721"></a>}

            <a id="L1723"></a>tystate[i] = DONE;
            <a id="L1724"></a>aryfil(temp1, nnonter+1, 0);

            <a id="L1726"></a><span class="comment">// take state i, close it, and do gotos</span>
            <a id="L1727"></a>closure(i);

            <a id="L1729"></a><span class="comment">// generate goto&#39;s</span>
            <a id="L1730"></a>for p := 0; p &lt; cwp; p++ {
                <a id="L1731"></a>pi := wsets[p];
                <a id="L1732"></a>if pi.flag != 0 {
                    <a id="L1733"></a>continue
                <a id="L1734"></a>}
                <a id="L1735"></a>wsets[p].flag = 1;
                <a id="L1736"></a>c := pi.pitem.first;
                <a id="L1737"></a>if c &lt;= 1 {
                    <a id="L1738"></a>if pstate[i+1]-pstate[i] &lt;= p {
                        <a id="L1739"></a>tystate[i] = MUSTLOOKAHEAD
                    <a id="L1740"></a>}
                    <a id="L1741"></a>continue;
                <a id="L1742"></a>}

                <a id="L1744"></a><span class="comment">// do a goto on c</span>
                <a id="L1745"></a>putitem(wsets[p].pitem, wsets[p].ws);
                <a id="L1746"></a>for q := p + 1; q &lt; cwp; q++ {
                    <a id="L1747"></a><span class="comment">// this item contributes to the goto</span>
                    <a id="L1748"></a>if c == wsets[q].pitem.first {
                        <a id="L1749"></a>putitem(wsets[q].pitem, wsets[q].ws);
                        <a id="L1750"></a>wsets[q].flag = 1;
                    <a id="L1751"></a>}
                <a id="L1752"></a>}

                <a id="L1754"></a>if c &lt; NTBASE {
                    <a id="L1755"></a>state(c) <span class="comment">// register new state</span>
                <a id="L1756"></a>} else {
                    <a id="L1757"></a>temp1[c-NTBASE] = state(c)
                <a id="L1758"></a>}
            <a id="L1759"></a>}

            <a id="L1761"></a>if gsdebug != 0 &amp;&amp; foutput != nil {
                <a id="L1762"></a>fmt.Fprintf(foutput, &#34;%v: &#34;, i);
                <a id="L1763"></a>for j := 0; j &lt;= nnonter; j++ {
                    <a id="L1764"></a>if temp1[j] != 0 {
                        <a id="L1765"></a>fmt.Fprintf(foutput, &#34;%v %v,&#34;, nontrst[j].name, temp1[j])
                    <a id="L1766"></a>}
                <a id="L1767"></a>}
                <a id="L1768"></a>fmt.Fprintf(foutput, &#34;\n&#34;);
            <a id="L1769"></a>}

            <a id="L1771"></a>if first != 0 {
                <a id="L1772"></a>indgo[i] = apack(temp1[1:len(temp1)], nnonter-1) - 1
            <a id="L1773"></a>}

            <a id="L1775"></a>more++;
        <a id="L1776"></a>}
    <a id="L1777"></a>}
<a id="L1778"></a>}

<a id="L1780"></a><span class="comment">//</span>
<a id="L1781"></a><span class="comment">// generate the closure of state i</span>
<a id="L1782"></a><span class="comment">//</span>
<a id="L1783"></a>func closure(i int) {
    <a id="L1784"></a>zzclose++;

    <a id="L1786"></a><span class="comment">// first, copy kernel of state i to wsets</span>
    <a id="L1787"></a>cwp = 0;
    <a id="L1788"></a>q := pstate[i+1];
    <a id="L1789"></a>for p := pstate[i]; p &lt; q; p++ {
        <a id="L1790"></a>wsets[cwp].pitem = statemem[p].pitem;
        <a id="L1791"></a>wsets[cwp].flag = 1; <span class="comment">// this item must get closed</span>
        <a id="L1792"></a>for ll := 0; ll &lt; len(wsets[cwp].ws); ll++ {
            <a id="L1793"></a>wsets[cwp].ws[ll] = statemem[p].look[ll]
        <a id="L1794"></a>}
        <a id="L1795"></a>cwp++;
    <a id="L1796"></a>}

    <a id="L1798"></a><span class="comment">// now, go through the loop, closing each item</span>
    <a id="L1799"></a>work := 1;
    <a id="L1800"></a>for work != 0 {
        <a id="L1801"></a>work = 0;
        <a id="L1802"></a>for u := 0; u &lt; cwp; u++ {
            <a id="L1803"></a>if wsets[u].flag == 0 {
                <a id="L1804"></a>continue
            <a id="L1805"></a>}

            <a id="L1807"></a><span class="comment">// dot is before c</span>
            <a id="L1808"></a>c := wsets[u].pitem.first;
            <a id="L1809"></a>if c &lt; NTBASE {
                <a id="L1810"></a>wsets[u].flag = 0;
                <a id="L1811"></a><span class="comment">// only interesting case is where . is before nonterminal</span>
                <a id="L1812"></a>continue;
            <a id="L1813"></a>}

            <a id="L1815"></a><span class="comment">// compute the lookahead</span>
            <a id="L1816"></a>aryfil(clset, tbitset, 0);

            <a id="L1818"></a><span class="comment">// find items involving c</span>
            <a id="L1819"></a>for v := u; v &lt; cwp; v++ {
                <a id="L1820"></a>if wsets[v].flag != 1 || wsets[v].pitem.first != c {
                    <a id="L1821"></a>continue
                <a id="L1822"></a>}
                <a id="L1823"></a>pi := wsets[v].pitem.prod;
                <a id="L1824"></a>ipi := wsets[v].pitem.off + 1;

                <a id="L1826"></a>wsets[v].flag = 0;
                <a id="L1827"></a>if nolook != 0 {
                    <a id="L1828"></a>continue
                <a id="L1829"></a>}

                <a id="L1831"></a>ch := pi[ipi];
                <a id="L1832"></a>ipi++;
                <a id="L1833"></a>for ch &gt; 0 {
                    <a id="L1834"></a><span class="comment">// terminal symbol</span>
                    <a id="L1835"></a>if ch &lt; NTBASE {
                        <a id="L1836"></a>setbit(clset, ch);
                        <a id="L1837"></a>break;
                    <a id="L1838"></a>}

                    <a id="L1840"></a><span class="comment">// nonterminal symbol</span>
                    <a id="L1841"></a>setunion(clset, pfirst[ch-NTBASE]);
                    <a id="L1842"></a>if pempty[ch-NTBASE] == 0 {
                        <a id="L1843"></a>break
                    <a id="L1844"></a>}
                    <a id="L1845"></a>ch = pi[ipi];
                    <a id="L1846"></a>ipi++;
                <a id="L1847"></a>}
                <a id="L1848"></a>if ch &lt;= 0 {
                    <a id="L1849"></a>setunion(clset, wsets[v].ws)
                <a id="L1850"></a>}
            <a id="L1851"></a>}

            <a id="L1853"></a><span class="comment">//</span>
            <a id="L1854"></a><span class="comment">// now loop over productions derived from c</span>
            <a id="L1855"></a><span class="comment">//</span>
            <a id="L1856"></a>curres := pres[c-NTBASE];
            <a id="L1857"></a>n := len(curres);

        <a id="L1859"></a>nexts:
            <a id="L1860"></a><span class="comment">// initially fill the sets</span>
            <a id="L1861"></a>for s := 0; s &lt; n; s++ {
                <a id="L1862"></a>prd := curres[s];

                <a id="L1864"></a><span class="comment">//</span>
                <a id="L1865"></a><span class="comment">// put these items into the closure</span>
                <a id="L1866"></a><span class="comment">// is the item there</span>
                <a id="L1867"></a><span class="comment">//</span>
                <a id="L1868"></a>for v := 0; v &lt; cwp; v++ {
                    <a id="L1869"></a><span class="comment">// yes, it is there</span>
                    <a id="L1870"></a>if wsets[v].pitem.off == 0 &amp;&amp;
                        <a id="L1871"></a>aryeq(wsets[v].pitem.prod, prd) != 0 {
                        <a id="L1872"></a>if nolook == 0 &amp;&amp;
                            <a id="L1873"></a>setunion(wsets[v].ws, clset) != 0 {
                            <a id="L1874"></a>wsets[v].flag = 1;
                            <a id="L1875"></a>work = 1;
                        <a id="L1876"></a>}
                        <a id="L1877"></a>continue nexts;
                    <a id="L1878"></a>}
                <a id="L1879"></a>}

                <a id="L1881"></a><span class="comment">//  not there; make a new entry</span>
                <a id="L1882"></a>if cwp &gt;= len(wsets) {
                    <a id="L1883"></a>awsets := make([]Wset, cwp+WSETINC);
                    <a id="L1884"></a>for ll := 0; ll &lt; len(wsets); ll++ {
                        <a id="L1885"></a>awsets[ll] = wsets[ll]
                    <a id="L1886"></a>}
                    <a id="L1887"></a>wsets = awsets;
                <a id="L1888"></a>}
                <a id="L1889"></a>wsets[cwp].pitem = Pitem{prd, 0, prd[0], -prd[len(prd)-1]};
                <a id="L1890"></a>wsets[cwp].flag = 1;
                <a id="L1891"></a>wsets[cwp].ws = mkset();
                <a id="L1892"></a>if nolook == 0 {
                    <a id="L1893"></a>work = 1;
                    <a id="L1894"></a>for ll := 0; ll &lt; len(wsets[cwp].ws); ll++ {
                        <a id="L1895"></a>wsets[cwp].ws[ll] = clset[ll]
                    <a id="L1896"></a>}
                <a id="L1897"></a>}
                <a id="L1898"></a>cwp++;
            <a id="L1899"></a>}
        <a id="L1900"></a>}
    <a id="L1901"></a>}

    <a id="L1903"></a><span class="comment">// have computed closure; flags are reset; return</span>
    <a id="L1904"></a>if cldebug != 0 &amp;&amp; foutput != nil {
        <a id="L1905"></a>fmt.Fprintf(foutput, &#34;\nState %v, nolook = %v\n&#34;, i, nolook);
        <a id="L1906"></a>for u := 0; u &lt; cwp; u++ {
            <a id="L1907"></a>if wsets[u].flag != 0 {
                <a id="L1908"></a>fmt.Fprintf(foutput, &#34;flag set\n&#34;)
            <a id="L1909"></a>}
            <a id="L1910"></a>wsets[u].flag = 0;
            <a id="L1911"></a>fmt.Fprintf(foutput, &#34;\t%v&#34;, writem(wsets[u].pitem));
            <a id="L1912"></a>prlook(wsets[u].ws);
            <a id="L1913"></a>fmt.Fprintf(foutput, &#34;\n&#34;);
        <a id="L1914"></a>}
    <a id="L1915"></a>}
<a id="L1916"></a>}

<a id="L1918"></a><span class="comment">//</span>
<a id="L1919"></a><span class="comment">// sorts last state,and sees if it equals earlier ones. returns state number</span>
<a id="L1920"></a><span class="comment">//</span>
<a id="L1921"></a>func state(c int) int {
    <a id="L1922"></a>zzstate++;
    <a id="L1923"></a>p1 := pstate[nstate];
    <a id="L1924"></a>p2 := pstate[nstate+1];
    <a id="L1925"></a>if p1 == p2 {
        <a id="L1926"></a>return 0 <span class="comment">// null state</span>
    <a id="L1927"></a>}

    <a id="L1929"></a><span class="comment">// sort the items</span>
    <a id="L1930"></a>var k, l int;
    <a id="L1931"></a>for k = p1 + 1; k &lt; p2; k++ { <span class="comment">// make k the biggest</span>
        <a id="L1932"></a>for l = k; l &gt; p1; l-- {
            <a id="L1933"></a>if statemem[l].pitem.prodno &lt; statemem[l-1].pitem.prodno ||
                <a id="L1934"></a>statemem[l].pitem.prodno == statemem[l-1].pitem.prodno &amp;&amp;
                    <a id="L1935"></a>statemem[l].pitem.off &lt; statemem[l-1].pitem.off {
                <a id="L1936"></a>s := statemem[l];
                <a id="L1937"></a>statemem[l] = statemem[l-1];
                <a id="L1938"></a>statemem[l-1] = s;
            <a id="L1939"></a>} else {
                <a id="L1940"></a>break
            <a id="L1941"></a>}
        <a id="L1942"></a>}
    <a id="L1943"></a>}

    <a id="L1945"></a>size1 := p2 - p1; <span class="comment">// size of state</span>

    <a id="L1947"></a>var i int;
    <a id="L1948"></a>if c &gt;= NTBASE {
        <a id="L1949"></a>i = ntstates[c-NTBASE]
    <a id="L1950"></a>} else {
        <a id="L1951"></a>i = tstates[c]
    <a id="L1952"></a>}

<a id="L1954"></a>look:
    <a id="L1955"></a>for ; i != 0; i = mstates[i] {
        <a id="L1956"></a><span class="comment">// get ith state</span>
        <a id="L1957"></a>q1 := pstate[i];
        <a id="L1958"></a>q2 := pstate[i+1];
        <a id="L1959"></a>size2 := q2 - q1;
        <a id="L1960"></a>if size1 != size2 {
            <a id="L1961"></a>continue
        <a id="L1962"></a>}
        <a id="L1963"></a>k = p1;
        <a id="L1964"></a>for l = q1; l &lt; q2; l++ {
            <a id="L1965"></a>if aryeq(statemem[l].pitem.prod, statemem[k].pitem.prod) == 0 ||
                <a id="L1966"></a>statemem[l].pitem.off != statemem[k].pitem.off {
                <a id="L1967"></a>continue look
            <a id="L1968"></a>}
            <a id="L1969"></a>k++;
        <a id="L1970"></a>}

        <a id="L1972"></a><span class="comment">// found it</span>
        <a id="L1973"></a>pstate[nstate+1] = pstate[nstate]; <span class="comment">// delete last state</span>

        <a id="L1975"></a><span class="comment">// fix up lookaheads</span>
        <a id="L1976"></a>if nolook != 0 {
            <a id="L1977"></a>return i
        <a id="L1978"></a>}
        <a id="L1979"></a>k = p1;
        <a id="L1980"></a>for l = q1; l &lt; q2; l++ {
            <a id="L1981"></a>if setunion(statemem[l].look, statemem[k].look) != 0 {
                <a id="L1982"></a>tystate[i] = MUSTDO
            <a id="L1983"></a>}
            <a id="L1984"></a>k++;
        <a id="L1985"></a>}
        <a id="L1986"></a>return i;
    <a id="L1987"></a>}

    <a id="L1989"></a><span class="comment">// state is new</span>
    <a id="L1990"></a>zznewstate++;
    <a id="L1991"></a>if nolook != 0 {
        <a id="L1992"></a>error(&#34;yacc state/nolook error&#34;)
    <a id="L1993"></a>}
    <a id="L1994"></a>pstate[nstate+2] = p2;
    <a id="L1995"></a>if nstate+1 &gt;= NSTATES {
        <a id="L1996"></a>error(&#34;too many states&#34;)
    <a id="L1997"></a>}
    <a id="L1998"></a>if c &gt;= NTBASE {
        <a id="L1999"></a>mstates[nstate] = ntstates[c-NTBASE];
        <a id="L2000"></a>ntstates[c-NTBASE] = nstate;
    <a id="L2001"></a>} else {
        <a id="L2002"></a>mstates[nstate] = tstates[c];
        <a id="L2003"></a>tstates[c] = nstate;
    <a id="L2004"></a>}
    <a id="L2005"></a>tystate[nstate] = MUSTDO;
    <a id="L2006"></a>nstate++;
    <a id="L2007"></a>return nstate - 1;
<a id="L2008"></a>}

<a id="L2010"></a>func putitem(p Pitem, set Lkset) {
    <a id="L2011"></a>p.off++;
    <a id="L2012"></a>p.first = p.prod[p.off];

    <a id="L2014"></a>if pidebug != 0 &amp;&amp; foutput != nil {
        <a id="L2015"></a>fmt.Fprintf(foutput, &#34;putitem(%v), state %v\n&#34;, writem(p), nstate)
    <a id="L2016"></a>}
    <a id="L2017"></a>j := pstate[nstate+1];
    <a id="L2018"></a>if j &gt;= len(statemem) {
        <a id="L2019"></a>asm := make([]Item, j+STATEINC);
        <a id="L2020"></a>for ll := 0; ll &lt; len(statemem); ll++ {
            <a id="L2021"></a>asm[ll] = statemem[ll]
        <a id="L2022"></a>}
        <a id="L2023"></a>statemem = asm;
    <a id="L2024"></a>}
    <a id="L2025"></a>statemem[j].pitem = p;
    <a id="L2026"></a>if nolook == 0 {
        <a id="L2027"></a>s := mkset();
        <a id="L2028"></a>for ll := 0; ll &lt; len(set); ll++ {
            <a id="L2029"></a>s[ll] = set[ll]
        <a id="L2030"></a>}
        <a id="L2031"></a>statemem[j].look = s;
    <a id="L2032"></a>}
    <a id="L2033"></a>j++;
    <a id="L2034"></a>pstate[nstate+1] = j;
<a id="L2035"></a>}

<a id="L2037"></a><span class="comment">//</span>
<a id="L2038"></a><span class="comment">// creates output string for item pointed to by pp</span>
<a id="L2039"></a><span class="comment">//</span>
<a id="L2040"></a>func writem(pp Pitem) string {
    <a id="L2041"></a>var i int;

    <a id="L2043"></a>p := pp.prod;
    <a id="L2044"></a>q := chcopy(nontrst[prdptr[pp.prodno][0]-NTBASE].name) + &#34;: &#34;;
    <a id="L2045"></a>npi := pp.off;

    <a id="L2047"></a>pi := aryeq(p, prdptr[pp.prodno]);

    <a id="L2049"></a>for {
        <a id="L2050"></a>c := &#39; &#39;;
        <a id="L2051"></a>if pi == npi {
            <a id="L2052"></a>c = &#39;.&#39;
        <a id="L2053"></a>}
        <a id="L2054"></a>q += string(c);

        <a id="L2056"></a>i = p[pi];
        <a id="L2057"></a>pi++;
        <a id="L2058"></a>if i &lt;= 0 {
            <a id="L2059"></a>break
        <a id="L2060"></a>}
        <a id="L2061"></a>q += chcopy(symnam(i));
    <a id="L2062"></a>}

    <a id="L2064"></a><span class="comment">// an item calling for a reduction</span>
    <a id="L2065"></a>i = p[npi];
    <a id="L2066"></a>if i &lt; 0 {
        <a id="L2067"></a>q += fmt.Sprintf(&#34;    (%v)&#34;, -i)
    <a id="L2068"></a>}

    <a id="L2070"></a>return q;
<a id="L2071"></a>}

<a id="L2073"></a><span class="comment">//</span>
<a id="L2074"></a><span class="comment">// pack state i from temp1 into amem</span>
<a id="L2075"></a><span class="comment">//</span>
<a id="L2076"></a>func apack(p []int, n int) int {
    <a id="L2077"></a><span class="comment">//</span>
    <a id="L2078"></a><span class="comment">// we don&#39;t need to worry about checking because</span>
    <a id="L2079"></a><span class="comment">// we will only look at entries known to be there...</span>
    <a id="L2080"></a><span class="comment">// eliminate leading and trailing 0&#39;s</span>
    <a id="L2081"></a><span class="comment">//</span>
    <a id="L2082"></a>off := 0;
    <a id="L2083"></a>pp := 0;
    <a id="L2084"></a>for ; pp &lt;= n &amp;&amp; p[pp] == 0; pp++ {
        <a id="L2085"></a>off--
    <a id="L2086"></a>}

    <a id="L2088"></a><span class="comment">// no actions</span>
    <a id="L2089"></a>if pp &gt; n {
        <a id="L2090"></a>return 0
    <a id="L2091"></a>}
    <a id="L2092"></a>for ; n &gt; pp &amp;&amp; p[n] == 0; n-- {
    <a id="L2093"></a>}
    <a id="L2094"></a>p = p[pp : n+1];

    <a id="L2096"></a><span class="comment">// now, find a place for the elements from p to q, inclusive</span>
    <a id="L2097"></a>r := len(amem) - len(p);

<a id="L2099"></a>nextk:
    <a id="L2100"></a>for rr := 0; rr &lt;= r; rr++ {
        <a id="L2101"></a>qq := rr;
        <a id="L2102"></a>for pp = 0; pp &lt; len(p); pp++ {
            <a id="L2103"></a>if p[pp] != 0 {
                <a id="L2104"></a>if p[pp] != amem[qq] &amp;&amp; amem[qq] != 0 {
                    <a id="L2105"></a>continue nextk
                <a id="L2106"></a>}
            <a id="L2107"></a>}
            <a id="L2108"></a>qq++;
        <a id="L2109"></a>}

        <a id="L2111"></a><span class="comment">// we have found an acceptable k</span>
        <a id="L2112"></a>if pkdebug != 0 &amp;&amp; foutput != nil {
            <a id="L2113"></a>fmt.Fprintf(foutput, &#34;off = %v, k = %v\n&#34;, off+rr, rr)
        <a id="L2114"></a>}
        <a id="L2115"></a>qq = rr;
        <a id="L2116"></a>for pp = 0; pp &lt; len(p); pp++ {
            <a id="L2117"></a>if p[pp] != 0 {
                <a id="L2118"></a>if qq &gt; memp {
                    <a id="L2119"></a>memp = qq
                <a id="L2120"></a>}
                <a id="L2121"></a>amem[qq] = p[pp];
            <a id="L2122"></a>}
            <a id="L2123"></a>qq++;
        <a id="L2124"></a>}
        <a id="L2125"></a>if pkdebug != 0 &amp;&amp; foutput != nil {
            <a id="L2126"></a>for pp = 0; pp &lt;= memp; pp += 10 {
                <a id="L2127"></a>fmt.Fprintf(foutput, &#34;\n&#34;);
                <a id="L2128"></a>for qq = pp; qq &lt;= pp+9; qq++ {
                    <a id="L2129"></a>fmt.Fprintf(foutput, &#34;%v &#34;, amem[qq])
                <a id="L2130"></a>}
                <a id="L2131"></a>fmt.Fprintf(foutput, &#34;\n&#34;);
            <a id="L2132"></a>}
        <a id="L2133"></a>}
        <a id="L2134"></a>return off + rr;
    <a id="L2135"></a>}
    <a id="L2136"></a>error(&#34;no space in action table&#34;);
    <a id="L2137"></a>return 0;
<a id="L2138"></a>}

<a id="L2140"></a><span class="comment">//</span>
<a id="L2141"></a><span class="comment">// print the output for the states</span>
<a id="L2142"></a><span class="comment">//</span>
<a id="L2143"></a>func output() {
    <a id="L2144"></a>var c, u, v int;

    <a id="L2146"></a>fmt.Fprintf(ftable, &#34;var\tYYEXCA = []int {\n&#34;);

    <a id="L2148"></a>noset := mkset();

    <a id="L2150"></a><span class="comment">// output the stuff for state i</span>
    <a id="L2151"></a>for i := 0; i &lt; nstate; i++ {
        <a id="L2152"></a>nolook = 0;
        <a id="L2153"></a>if tystate[i] != MUSTLOOKAHEAD {
            <a id="L2154"></a>nolook = 1
        <a id="L2155"></a>}
        <a id="L2156"></a>closure(i);

        <a id="L2158"></a><span class="comment">// output actions</span>
        <a id="L2159"></a>nolook = 1;
        <a id="L2160"></a>aryfil(temp1, ntokens+nnonter+1, 0);
        <a id="L2161"></a>for u = 0; u &lt; cwp; u++ {
            <a id="L2162"></a>c = wsets[u].pitem.first;
            <a id="L2163"></a>if c &gt; 1 &amp;&amp; c &lt; NTBASE &amp;&amp; temp1[c] == 0 {
                <a id="L2164"></a>for v = u; v &lt; cwp; v++ {
                    <a id="L2165"></a>if c == wsets[v].pitem.first {
                        <a id="L2166"></a>putitem(wsets[v].pitem, noset)
                    <a id="L2167"></a>}
                <a id="L2168"></a>}
                <a id="L2169"></a>temp1[c] = state(c);
            <a id="L2170"></a>} else if c &gt; NTBASE {
                <a id="L2171"></a>c -= NTBASE;
                <a id="L2172"></a>if temp1[c+ntokens] == 0 {
                    <a id="L2173"></a>temp1[c+ntokens] = amem[indgo[i]+c]
                <a id="L2174"></a>}
            <a id="L2175"></a>}
        <a id="L2176"></a>}
        <a id="L2177"></a>if i == 1 {
            <a id="L2178"></a>temp1[1] = ACCEPTCODE
        <a id="L2179"></a>}

        <a id="L2181"></a><span class="comment">// now, we have the shifts; look at the reductions</span>
        <a id="L2182"></a>lastred = 0;
        <a id="L2183"></a>for u = 0; u &lt; cwp; u++ {
            <a id="L2184"></a>c = wsets[u].pitem.first;

            <a id="L2186"></a><span class="comment">// reduction</span>
            <a id="L2187"></a>if c &gt; 0 {
                <a id="L2188"></a>continue
            <a id="L2189"></a>}
            <a id="L2190"></a>lastred = -c;
            <a id="L2191"></a>us := wsets[u].ws;
            <a id="L2192"></a>for k := 0; k &lt;= ntokens; k++ {
                <a id="L2193"></a>if bitset(us, k) == 0 {
                    <a id="L2194"></a>continue
                <a id="L2195"></a>}
                <a id="L2196"></a>if temp1[k] == 0 {
                    <a id="L2197"></a>temp1[k] = c
                <a id="L2198"></a>} else if temp1[k] &lt; 0 { <span class="comment">// reduce/reduce conflict</span>
                    <a id="L2199"></a>if foutput != nil {
                        <a id="L2200"></a>fmt.Fprintf(foutput,
                            <a id="L2201"></a>&#34;\n %v: reduce/reduce conflict  (red&#39;ns &#34;
                                <a id="L2202"></a>&#34;%v and %v) on %v&#34;,
                            <a id="L2203"></a>i, -temp1[k], lastred, symnam(k))
                    <a id="L2204"></a>}
                    <a id="L2205"></a>if -temp1[k] &gt; lastred {
                        <a id="L2206"></a>temp1[k] = -lastred
                    <a id="L2207"></a>}
                    <a id="L2208"></a>zzrrconf++;
                <a id="L2209"></a>} else {
                    <a id="L2210"></a><span class="comment">// potential shift/reduce conflict</span>
                    <a id="L2211"></a>precftn(lastred, k, i)
                <a id="L2212"></a>}
            <a id="L2213"></a>}
        <a id="L2214"></a>}
        <a id="L2215"></a>wract(i);
    <a id="L2216"></a>}

    <a id="L2218"></a>fmt.Fprintf(ftable, &#34;}\n&#34;);
    <a id="L2219"></a>fmt.Fprintf(ftable, &#34;const\tYYNPROD\t= %v\n&#34;, nprod);
    <a id="L2220"></a>fmt.Fprintf(ftable, &#34;const\tYYPRIVATE\t= %v\n&#34;, PRIVATE);
    <a id="L2221"></a>fmt.Fprintf(ftable, &#34;var\tYYTOKENNAMES []string\n&#34;);
    <a id="L2222"></a>fmt.Fprintf(ftable, &#34;var\tYYSTATES\n[]string\n&#34;);
<a id="L2223"></a>}

<a id="L2225"></a><span class="comment">//</span>
<a id="L2226"></a><span class="comment">// decide a shift/reduce conflict by precedence.</span>
<a id="L2227"></a><span class="comment">// r is a rule number, t a token number</span>
<a id="L2228"></a><span class="comment">// the conflict is in state s</span>
<a id="L2229"></a><span class="comment">// temp1[t] is changed to reflect the action</span>
<a id="L2230"></a><span class="comment">//</span>
<a id="L2231"></a>func precftn(r, t, s int) {
    <a id="L2232"></a>var action int;

    <a id="L2234"></a>lp := levprd[r];
    <a id="L2235"></a>lt := toklev[t];
    <a id="L2236"></a>if PLEVEL(lt) == 0 || PLEVEL(lp) == 0 {
        <a id="L2237"></a><span class="comment">// conflict</span>
        <a id="L2238"></a>if foutput != nil {
            <a id="L2239"></a>fmt.Fprintf(foutput,
                <a id="L2240"></a>&#34;\n%v: shift/reduce conflict (shift %v(%v), red&#39;n %v(%v)) on %v&#34;,
                <a id="L2241"></a>s, temp1[t], PLEVEL(lt), r, PLEVEL(lp), symnam(t))
        <a id="L2242"></a>}
        <a id="L2243"></a>zzsrconf++;
        <a id="L2244"></a>return;
    <a id="L2245"></a>}
    <a id="L2246"></a>if PLEVEL(lt) == PLEVEL(lp) {
        <a id="L2247"></a>action = ASSOC(lt)
    <a id="L2248"></a>} else if PLEVEL(lt) &gt; PLEVEL(lp) {
        <a id="L2249"></a>action = RASC <span class="comment">// shift</span>
    <a id="L2250"></a>} else {
        <a id="L2251"></a>action = LASC
    <a id="L2252"></a>}   <span class="comment">// reduce</span>
    <a id="L2253"></a>switch action {
    <a id="L2254"></a>case BASC: <span class="comment">// error action</span>
        <a id="L2255"></a>temp1[t] = ERRCODE
    <a id="L2256"></a>case LASC: <span class="comment">// reduce</span>
        <a id="L2257"></a>temp1[t] = -r
    <a id="L2258"></a>}
<a id="L2259"></a>}

<a id="L2261"></a><span class="comment">//</span>
<a id="L2262"></a><span class="comment">// output state i</span>
<a id="L2263"></a><span class="comment">// temp1 has the actions, lastred the default</span>
<a id="L2264"></a><span class="comment">//</span>
<a id="L2265"></a>func wract(i int) {
    <a id="L2266"></a>var p, p1 int;

    <a id="L2268"></a><span class="comment">// find the best choice for lastred</span>
    <a id="L2269"></a>lastred = 0;
    <a id="L2270"></a>ntimes := 0;
    <a id="L2271"></a>for j := 0; j &lt;= ntokens; j++ {
        <a id="L2272"></a>if temp1[j] &gt;= 0 {
            <a id="L2273"></a>continue
        <a id="L2274"></a>}
        <a id="L2275"></a>if temp1[j]+lastred == 0 {
            <a id="L2276"></a>continue
        <a id="L2277"></a>}
        <a id="L2278"></a><span class="comment">// count the number of appearances of temp1[j]</span>
        <a id="L2279"></a>count := 0;
        <a id="L2280"></a>tred := -temp1[j];
        <a id="L2281"></a>levprd[tred] |= REDFLAG;
        <a id="L2282"></a>for p = 0; p &lt;= ntokens; p++ {
            <a id="L2283"></a>if temp1[p]+tred == 0 {
                <a id="L2284"></a>count++
            <a id="L2285"></a>}
        <a id="L2286"></a>}
        <a id="L2287"></a>if count &gt; ntimes {
            <a id="L2288"></a>lastred = tred;
            <a id="L2289"></a>ntimes = count;
        <a id="L2290"></a>}
    <a id="L2291"></a>}

    <a id="L2293"></a><span class="comment">//</span>
    <a id="L2294"></a><span class="comment">// for error recovery, arrange that, if there is a shift on the</span>
    <a id="L2295"></a><span class="comment">// error recovery token, `error&#39;, that the default be the error action</span>
    <a id="L2296"></a><span class="comment">//</span>
    <a id="L2297"></a>if temp1[2] &gt; 0 {
        <a id="L2298"></a>lastred = 0
    <a id="L2299"></a>}

    <a id="L2301"></a><span class="comment">// clear out entries in temp1 which equal lastred</span>
    <a id="L2302"></a><span class="comment">// count entries in optst table</span>
    <a id="L2303"></a>n := 0;
    <a id="L2304"></a>for p = 0; p &lt;= ntokens; p++ {
        <a id="L2305"></a>p1 = temp1[p];
        <a id="L2306"></a>if p1+lastred == 0 {
            <a id="L2307"></a>temp1[p] = 0;
            <a id="L2308"></a>p1 = 0;
        <a id="L2309"></a>}
        <a id="L2310"></a>if p1 &gt; 0 &amp;&amp; p1 != ACCEPTCODE &amp;&amp; p1 != ERRCODE {
            <a id="L2311"></a>n++
        <a id="L2312"></a>}
    <a id="L2313"></a>}

    <a id="L2315"></a>wrstate(i);
    <a id="L2316"></a>defact[i] = lastred;
    <a id="L2317"></a>flag := 0;
    <a id="L2318"></a>os := make([]int, n*2);
    <a id="L2319"></a>n = 0;
    <a id="L2320"></a>for p = 0; p &lt;= ntokens; p++ {
        <a id="L2321"></a>p1 = temp1[p];
        <a id="L2322"></a>if p1 != 0 {
            <a id="L2323"></a>if p1 &lt; 0 {
                <a id="L2324"></a>p1 = -p1
            <a id="L2325"></a>} else if p1 == ACCEPTCODE {
                <a id="L2326"></a>p1 = -1
            <a id="L2327"></a>} else if p1 == ERRCODE {
                <a id="L2328"></a>p1 = 0
            <a id="L2329"></a>} else {
                <a id="L2330"></a>os[n] = p;
                <a id="L2331"></a>n++;
                <a id="L2332"></a>os[n] = p1;
                <a id="L2333"></a>n++;
                <a id="L2334"></a>zzacent++;
                <a id="L2335"></a>continue;
            <a id="L2336"></a>}
            <a id="L2337"></a>if flag == 0 {
                <a id="L2338"></a>fmt.Fprintf(ftable, &#34;-1, %v,\n&#34;, i)
            <a id="L2339"></a>}
            <a id="L2340"></a>flag++;
            <a id="L2341"></a>fmt.Fprintf(ftable, &#34;\t%v, %v,\n&#34;, p, p1);
            <a id="L2342"></a>zzexcp++;
        <a id="L2343"></a>}
    <a id="L2344"></a>}
    <a id="L2345"></a>if flag != 0 {
        <a id="L2346"></a>defact[i] = -2;
        <a id="L2347"></a>fmt.Fprintf(ftable, &#34;\t-2, %v,\n&#34;, lastred);
    <a id="L2348"></a>}
    <a id="L2349"></a>optst[i] = os;
<a id="L2350"></a>}

<a id="L2352"></a><span class="comment">//</span>
<a id="L2353"></a><span class="comment">// writes state i</span>
<a id="L2354"></a><span class="comment">//</span>
<a id="L2355"></a>func wrstate(i int) {
    <a id="L2356"></a>var j0, j1, u int;
    <a id="L2357"></a>var pp, qq int;

    <a id="L2359"></a>if foutput == nil {
        <a id="L2360"></a>return
    <a id="L2361"></a>}
    <a id="L2362"></a>fmt.Fprintf(foutput, &#34;\nstate %v\n&#34;, i);
    <a id="L2363"></a>qq = pstate[i+1];
    <a id="L2364"></a>for pp = pstate[i]; pp &lt; qq; pp++ {
        <a id="L2365"></a>fmt.Fprintf(foutput, &#34;\t%v\n&#34;, writem(statemem[pp].pitem))
    <a id="L2366"></a>}
    <a id="L2367"></a>if tystate[i] == MUSTLOOKAHEAD {
        <a id="L2368"></a><span class="comment">// print out empty productions in closure</span>
        <a id="L2369"></a>for u = pstate[i+1] - pstate[i]; u &lt; cwp; u++ {
            <a id="L2370"></a>if wsets[u].pitem.first &lt; 0 {
                <a id="L2371"></a>fmt.Fprintf(foutput, &#34;\t%v\n&#34;, writem(wsets[u].pitem))
            <a id="L2372"></a>}
        <a id="L2373"></a>}
    <a id="L2374"></a>}

    <a id="L2376"></a><span class="comment">// check for state equal to another</span>
    <a id="L2377"></a>for j0 = 0; j0 &lt;= ntokens; j0++ {
        <a id="L2378"></a>j1 = temp1[j0];
        <a id="L2379"></a>if j1 != 0 {
            <a id="L2380"></a>fmt.Fprintf(foutput, &#34;\n\t%v  &#34;, symnam(j0));

            <a id="L2382"></a><span class="comment">// shift, error, or accept</span>
            <a id="L2383"></a>if j1 &gt; 0 {
                <a id="L2384"></a>if j1 == ACCEPTCODE {
                    <a id="L2385"></a>fmt.Fprintf(foutput, &#34;accept&#34;)
                <a id="L2386"></a>} else if j1 == ERRCODE {
                    <a id="L2387"></a>fmt.Fprintf(foutput, &#34;error&#34;)
                <a id="L2388"></a>} else {
                    <a id="L2389"></a>fmt.Fprintf(foutput, &#34;shift %v&#34;, j1)
                <a id="L2390"></a>}
            <a id="L2391"></a>} else {
                <a id="L2392"></a>fmt.Fprintf(foutput, &#34;reduce %v (src line %v)&#34;, -j1, rlines[-j1])
            <a id="L2393"></a>}
        <a id="L2394"></a>}
    <a id="L2395"></a>}

    <a id="L2397"></a><span class="comment">// output the final production</span>
    <a id="L2398"></a>if lastred != 0 {
        <a id="L2399"></a>fmt.Fprintf(foutput, &#34;\n\t.  reduce %v (src line %v)\n\n&#34;,
            <a id="L2400"></a>lastred, rlines[lastred])
    <a id="L2401"></a>} else {
        <a id="L2402"></a>fmt.Fprintf(foutput, &#34;\n\t.  error\n\n&#34;)
    <a id="L2403"></a>}

    <a id="L2405"></a><span class="comment">// now, output nonterminal actions</span>
    <a id="L2406"></a>j1 = ntokens;
    <a id="L2407"></a>for j0 = 1; j0 &lt;= nnonter; j0++ {
        <a id="L2408"></a>j1++;
        <a id="L2409"></a>if temp1[j1] != 0 {
            <a id="L2410"></a>fmt.Fprintf(foutput, &#34;\t%v  goto %v\n&#34;, symnam(j0+NTBASE), temp1[j1])
        <a id="L2411"></a>}
    <a id="L2412"></a>}
<a id="L2413"></a>}

<a id="L2415"></a><span class="comment">//</span>
<a id="L2416"></a><span class="comment">// output the gotos for the nontermninals</span>
<a id="L2417"></a><span class="comment">//</span>
<a id="L2418"></a>func go2out() {
    <a id="L2419"></a>for i := 1; i &lt;= nnonter; i++ {
        <a id="L2420"></a>go2gen(i);

        <a id="L2422"></a><span class="comment">// find the best one to make default</span>
        <a id="L2423"></a>best := -1;
        <a id="L2424"></a>times := 0;

        <a id="L2426"></a><span class="comment">// is j the most frequent</span>
        <a id="L2427"></a>for j := 0; j &lt; nstate; j++ {
            <a id="L2428"></a>if tystate[j] == 0 {
                <a id="L2429"></a>continue
            <a id="L2430"></a>}
            <a id="L2431"></a>if tystate[j] == best {
                <a id="L2432"></a>continue
            <a id="L2433"></a>}

            <a id="L2435"></a><span class="comment">// is tystate[j] the most frequent</span>
            <a id="L2436"></a>count := 0;
            <a id="L2437"></a>cbest := tystate[j];
            <a id="L2438"></a>for k := j; k &lt; nstate; k++ {
                <a id="L2439"></a>if tystate[k] == cbest {
                    <a id="L2440"></a>count++
                <a id="L2441"></a>}
            <a id="L2442"></a>}
            <a id="L2443"></a>if count &gt; times {
                <a id="L2444"></a>best = cbest;
                <a id="L2445"></a>times = count;
            <a id="L2446"></a>}
        <a id="L2447"></a>}

        <a id="L2449"></a><span class="comment">// best is now the default entry</span>
        <a id="L2450"></a>zzgobest += times - 1;
        <a id="L2451"></a>n := 0;
        <a id="L2452"></a>for j := 0; j &lt; nstate; j++ {
            <a id="L2453"></a>if tystate[j] != 0 &amp;&amp; tystate[j] != best {
                <a id="L2454"></a>n++
            <a id="L2455"></a>}
        <a id="L2456"></a>}
        <a id="L2457"></a>goent := make([]int, 2*n+1);
        <a id="L2458"></a>n = 0;
        <a id="L2459"></a>for j := 0; j &lt; nstate; j++ {
            <a id="L2460"></a>if tystate[j] != 0 &amp;&amp; tystate[j] != best {
                <a id="L2461"></a>goent[n] = j;
                <a id="L2462"></a>n++;
                <a id="L2463"></a>goent[n] = tystate[j];
                <a id="L2464"></a>n++;
                <a id="L2465"></a>zzgoent++;
            <a id="L2466"></a>}
        <a id="L2467"></a>}

        <a id="L2469"></a><span class="comment">// now, the default</span>
        <a id="L2470"></a>if best == -1 {
            <a id="L2471"></a>best = 0
        <a id="L2472"></a>}

        <a id="L2474"></a>zzgoent++;
        <a id="L2475"></a>goent[n] = best;
        <a id="L2476"></a>yypgo[i] = goent;
    <a id="L2477"></a>}
<a id="L2478"></a>}

<a id="L2480"></a><span class="comment">//</span>
<a id="L2481"></a><span class="comment">// output the gotos for nonterminal c</span>
<a id="L2482"></a><span class="comment">//</span>
<a id="L2483"></a>func go2gen(c int) {
    <a id="L2484"></a>var i, cc, p, q int;

    <a id="L2486"></a><span class="comment">// first, find nonterminals with gotos on c</span>
    <a id="L2487"></a>aryfil(temp1, nnonter+1, 0);
    <a id="L2488"></a>temp1[c] = 1;
    <a id="L2489"></a>work := 1;
    <a id="L2490"></a>for work != 0 {
        <a id="L2491"></a>work = 0;
        <a id="L2492"></a>for i = 0; i &lt; nprod; i++ {
            <a id="L2493"></a><span class="comment">// cc is a nonterminal with a goto on c</span>
            <a id="L2494"></a>cc = prdptr[i][1] - NTBASE;
            <a id="L2495"></a>if cc &gt;= 0 &amp;&amp; temp1[cc] != 0 {
                <a id="L2496"></a><span class="comment">// thus, the left side of production i does too</span>
                <a id="L2497"></a>cc = prdptr[i][0] - NTBASE;
                <a id="L2498"></a>if temp1[cc] == 0 {
                    <a id="L2499"></a>work = 1;
                    <a id="L2500"></a>temp1[cc] = 1;
                <a id="L2501"></a>}
            <a id="L2502"></a>}
        <a id="L2503"></a>}
    <a id="L2504"></a>}

    <a id="L2506"></a><span class="comment">// now, we have temp1[c] = 1 if a goto on c in closure of cc</span>
    <a id="L2507"></a>if g2debug != 0 &amp;&amp; foutput != nil {
        <a id="L2508"></a>fmt.Fprintf(foutput, &#34;%v: gotos on &#34;, nontrst[c].name);
        <a id="L2509"></a>for i = 0; i &lt;= nnonter; i++ {
            <a id="L2510"></a>if temp1[i] != 0 {
                <a id="L2511"></a>fmt.Fprintf(foutput, &#34;%v &#34;, nontrst[i].name)
            <a id="L2512"></a>}
        <a id="L2513"></a>}
        <a id="L2514"></a>fmt.Fprintf(foutput, &#34;\n&#34;);
    <a id="L2515"></a>}

    <a id="L2517"></a><span class="comment">// now, go through and put gotos into tystate</span>
    <a id="L2518"></a>aryfil(tystate, nstate, 0);
    <a id="L2519"></a>for i = 0; i &lt; nstate; i++ {
        <a id="L2520"></a>q = pstate[i+1];
        <a id="L2521"></a>for p = pstate[i]; p &lt; q; p++ {
            <a id="L2522"></a>cc = statemem[p].pitem.first;
            <a id="L2523"></a>if cc &gt;= NTBASE {
                <a id="L2524"></a><span class="comment">// goto on c is possible</span>
                <a id="L2525"></a>if temp1[cc-NTBASE] != 0 {
                    <a id="L2526"></a>tystate[i] = amem[indgo[i]+c];
                    <a id="L2527"></a>break;
                <a id="L2528"></a>}
            <a id="L2529"></a>}
        <a id="L2530"></a>}
    <a id="L2531"></a>}
<a id="L2532"></a>}

<a id="L2534"></a><span class="comment">//</span>
<a id="L2535"></a><span class="comment">// in order to free up the mem and amem arrays for the optimizer,</span>
<a id="L2536"></a><span class="comment">// and still be able to output yyr1, etc., after the sizes of</span>
<a id="L2537"></a><span class="comment">// the action array is known, we hide the nonterminals</span>
<a id="L2538"></a><span class="comment">// derived by productions in levprd.</span>
<a id="L2539"></a><span class="comment">//</span>
<a id="L2540"></a>func hideprod() {
    <a id="L2541"></a>nred := 0;
    <a id="L2542"></a>levprd[0] = 0;
    <a id="L2543"></a>for i := 1; i &lt; nprod; i++ {
        <a id="L2544"></a>if (levprd[i] &amp; REDFLAG) == 0 {
            <a id="L2545"></a>if foutput != nil {
                <a id="L2546"></a>fmt.Fprintf(foutput, &#34;Rule not reduced: %v\n&#34;,
                    <a id="L2547"></a>writem(Pitem{prdptr[i], 0, 0, i}))
            <a id="L2548"></a>}
            <a id="L2549"></a>fmt.Printf(&#34;rule %v never reduced\n&#34;, writem(Pitem{prdptr[i], 0, 0, i}));
            <a id="L2550"></a>nred++;
        <a id="L2551"></a>}
        <a id="L2552"></a>levprd[i] = prdptr[i][0] - NTBASE;
    <a id="L2553"></a>}
    <a id="L2554"></a>if nred != 0 {
        <a id="L2555"></a>fmt.Printf(&#34;%v rules never reduced\n&#34;, nred)
    <a id="L2556"></a>}
<a id="L2557"></a>}

<a id="L2559"></a>func callopt() {
    <a id="L2560"></a>var j, k, p, q, i int;
    <a id="L2561"></a>var v []int;

    <a id="L2563"></a>pgo = make([]int, nnonter+1);
    <a id="L2564"></a>pgo[0] = 0;
    <a id="L2565"></a>maxoff = 0;
    <a id="L2566"></a>maxspr = 0;
    <a id="L2567"></a>for i = 0; i &lt; nstate; i++ {
        <a id="L2568"></a>k = 32000;
        <a id="L2569"></a>j = 0;
        <a id="L2570"></a>v = optst[i];
        <a id="L2571"></a>q = len(v);
        <a id="L2572"></a>for p = 0; p &lt; q; p += 2 {
            <a id="L2573"></a>if v[p] &gt; j {
                <a id="L2574"></a>j = v[p]
            <a id="L2575"></a>}
            <a id="L2576"></a>if v[p] &lt; k {
                <a id="L2577"></a>k = v[p]
            <a id="L2578"></a>}
        <a id="L2579"></a>}

        <a id="L2581"></a><span class="comment">// nontrivial situation</span>
        <a id="L2582"></a>if k &lt;= j {
            <a id="L2583"></a><span class="comment">// j is now the range</span>
            <a id="L2584"></a><span class="comment">//			j -= k;			// call scj</span>
            <a id="L2585"></a>if k &gt; maxoff {
                <a id="L2586"></a>maxoff = k
            <a id="L2587"></a>}
        <a id="L2588"></a>}
        <a id="L2589"></a>tystate[i] = q + 2*j;
        <a id="L2590"></a>if j &gt; maxspr {
            <a id="L2591"></a>maxspr = j
        <a id="L2592"></a>}
    <a id="L2593"></a>}

    <a id="L2595"></a><span class="comment">// initialize ggreed table</span>
    <a id="L2596"></a>ggreed = make([]int, nnonter+1);
    <a id="L2597"></a>for i = 1; i &lt;= nnonter; i++ {
        <a id="L2598"></a>ggreed[i] = 1;
        <a id="L2599"></a>j = 0;

        <a id="L2601"></a><span class="comment">// minimum entry index is always 0</span>
        <a id="L2602"></a>v = yypgo[i];
        <a id="L2603"></a>q = len(v) - 1;
        <a id="L2604"></a>for p = 0; p &lt; q; p += 2 {
            <a id="L2605"></a>ggreed[i] += 2;
            <a id="L2606"></a>if v[p] &gt; j {
                <a id="L2607"></a>j = v[p]
            <a id="L2608"></a>}
        <a id="L2609"></a>}
        <a id="L2610"></a>ggreed[i] = ggreed[i] + 2*j;
        <a id="L2611"></a>if j &gt; maxoff {
            <a id="L2612"></a>maxoff = j
        <a id="L2613"></a>}
    <a id="L2614"></a>}

    <a id="L2616"></a><span class="comment">// now, prepare to put the shift actions into the amem array</span>
    <a id="L2617"></a>for i = 0; i &lt; ACTSIZE; i++ {
        <a id="L2618"></a>amem[i] = 0
    <a id="L2619"></a>}
    <a id="L2620"></a>maxa = 0;
    <a id="L2621"></a>for i = 0; i &lt; nstate; i++ {
        <a id="L2622"></a>if tystate[i] == 0 &amp;&amp; adb &gt; 1 {
            <a id="L2623"></a>fmt.Fprintf(ftable, &#34;State %v: null\n&#34;, i)
        <a id="L2624"></a>}
        <a id="L2625"></a>indgo[i] = YYFLAG;
    <a id="L2626"></a>}

    <a id="L2628"></a>i = nxti();
    <a id="L2629"></a>for i != NOMORE {
        <a id="L2630"></a>if i &gt;= 0 {
            <a id="L2631"></a>stin(i)
        <a id="L2632"></a>} else {
            <a id="L2633"></a>gin(-i)
        <a id="L2634"></a>}
        <a id="L2635"></a>i = nxti();
    <a id="L2636"></a>}

    <a id="L2638"></a><span class="comment">// print amem array</span>
    <a id="L2639"></a>if adb &gt; 2 {
        <a id="L2640"></a>for p = 0; p &lt;= maxa; p += 10 {
            <a id="L2641"></a>fmt.Fprintf(ftable, &#34;%v  &#34;, p);
            <a id="L2642"></a>for i = 0; i &lt; 10; i++ {
                <a id="L2643"></a>fmt.Fprintf(ftable, &#34;%v  &#34;, amem[p+i])
            <a id="L2644"></a>}
            <a id="L2645"></a>putrune(ftable, &#39;\n&#39;);
        <a id="L2646"></a>}
    <a id="L2647"></a>}

    <a id="L2649"></a>aoutput();
    <a id="L2650"></a>osummary();
<a id="L2651"></a>}

<a id="L2653"></a><span class="comment">//</span>
<a id="L2654"></a><span class="comment">// finds the next i</span>
<a id="L2655"></a><span class="comment">//</span>
<a id="L2656"></a>func nxti() int {
    <a id="L2657"></a>max := 0;
    <a id="L2658"></a>maxi := 0;
    <a id="L2659"></a>for i := 1; i &lt;= nnonter; i++ {
        <a id="L2660"></a>if ggreed[i] &gt;= max {
            <a id="L2661"></a>max = ggreed[i];
            <a id="L2662"></a>maxi = -i;
        <a id="L2663"></a>}
    <a id="L2664"></a>}
    <a id="L2665"></a>for i := 0; i &lt; nstate; i++ {
        <a id="L2666"></a>if tystate[i] &gt;= max {
            <a id="L2667"></a>max = tystate[i];
            <a id="L2668"></a>maxi = i;
        <a id="L2669"></a>}
    <a id="L2670"></a>}
    <a id="L2671"></a>if max == 0 {
        <a id="L2672"></a>return NOMORE
    <a id="L2673"></a>}
    <a id="L2674"></a>return maxi;
<a id="L2675"></a>}

<a id="L2677"></a>func gin(i int) {
    <a id="L2678"></a>var s int;

    <a id="L2680"></a><span class="comment">// enter gotos on nonterminal i into array amem</span>
    <a id="L2681"></a>ggreed[i] = 0;

    <a id="L2683"></a>q := yypgo[i];
    <a id="L2684"></a>nq := len(q) - 1;

    <a id="L2686"></a><span class="comment">// now, find amem place for it</span>
<a id="L2687"></a>nextgp:
    <a id="L2688"></a>for p := 0; p &lt; ACTSIZE; p++ {
        <a id="L2689"></a>if amem[p] != 0 {
            <a id="L2690"></a>continue
        <a id="L2691"></a>}
        <a id="L2692"></a>for r := 0; r &lt; nq; r += 2 {
            <a id="L2693"></a>s = p + q[r] + 1;
            <a id="L2694"></a>if s &gt; maxa {
                <a id="L2695"></a>maxa = s;
                <a id="L2696"></a>if maxa &gt;= ACTSIZE {
                    <a id="L2697"></a>error(&#34;a array overflow&#34;)
                <a id="L2698"></a>}
            <a id="L2699"></a>}
            <a id="L2700"></a>if amem[s] != 0 {
                <a id="L2701"></a>continue nextgp
            <a id="L2702"></a>}
        <a id="L2703"></a>}

        <a id="L2705"></a><span class="comment">// we have found amem spot</span>
        <a id="L2706"></a>amem[p] = q[nq];
        <a id="L2707"></a>if p &gt; maxa {
            <a id="L2708"></a>maxa = p
        <a id="L2709"></a>}
        <a id="L2710"></a>for r := 0; r &lt; nq; r += 2 {
            <a id="L2711"></a>s = p + q[r] + 1;
            <a id="L2712"></a>amem[s] = q[r+1];
        <a id="L2713"></a>}
        <a id="L2714"></a>pgo[i] = p;
        <a id="L2715"></a>if adb &gt; 1 {
            <a id="L2716"></a>fmt.Fprintf(ftable, &#34;Nonterminal %v, entry at %v\n&#34;, i, pgo[i])
        <a id="L2717"></a>}
        <a id="L2718"></a>return;
    <a id="L2719"></a>}
    <a id="L2720"></a>error(&#34;cannot place goto %v\n&#34;, i);
<a id="L2721"></a>}

<a id="L2723"></a>func stin(i int) {
    <a id="L2724"></a>var s int;

    <a id="L2726"></a>tystate[i] = 0;

    <a id="L2728"></a><span class="comment">// enter state i into the amem array</span>
    <a id="L2729"></a>q := optst[i];
    <a id="L2730"></a>nq := len(q);

<a id="L2732"></a>nextn:
    <a id="L2733"></a><span class="comment">// find an acceptable place</span>
    <a id="L2734"></a>for n := -maxoff; n &lt; ACTSIZE; n++ {
        <a id="L2735"></a>flag := 0;
        <a id="L2736"></a>for r := 0; r &lt; nq; r += 2 {
            <a id="L2737"></a>s = q[r] + n;
            <a id="L2738"></a>if s &lt; 0 || s &gt; ACTSIZE {
                <a id="L2739"></a>continue nextn
            <a id="L2740"></a>}
            <a id="L2741"></a>if amem[s] == 0 {
                <a id="L2742"></a>flag++
            <a id="L2743"></a>} else if amem[s] != q[r+1] {
                <a id="L2744"></a>continue nextn
            <a id="L2745"></a>}
        <a id="L2746"></a>}

        <a id="L2748"></a><span class="comment">// check the position equals another only if the states are identical</span>
        <a id="L2749"></a>for j := 0; j &lt; nstate; j++ {
            <a id="L2750"></a>if indgo[j] == n {

                <a id="L2752"></a><span class="comment">// we have some disagreement</span>
                <a id="L2753"></a>if flag != 0 {
                    <a id="L2754"></a>continue nextn
                <a id="L2755"></a>}
                <a id="L2756"></a>if nq == len(optst[j]) {

                    <a id="L2758"></a><span class="comment">// states are equal</span>
                    <a id="L2759"></a>indgo[i] = n;
                    <a id="L2760"></a>if adb &gt; 1 {
                        <a id="L2761"></a>fmt.Fprintf(ftable, &#34;State %v: entry at&#34;
                            <a id="L2762"></a>&#34;%v equals state %v\n&#34;,
                            <a id="L2763"></a>i, n, j)
                    <a id="L2764"></a>}
                    <a id="L2765"></a>return;
                <a id="L2766"></a>}

                <a id="L2768"></a><span class="comment">// we have some disagreement</span>
                <a id="L2769"></a>continue nextn;
            <a id="L2770"></a>}
        <a id="L2771"></a>}

        <a id="L2773"></a>for r := 0; r &lt; nq; r += 2 {
            <a id="L2774"></a>s = q[r] + n;
            <a id="L2775"></a>if s &gt; maxa {
                <a id="L2776"></a>maxa = s
            <a id="L2777"></a>}
            <a id="L2778"></a>if amem[s] != 0 &amp;&amp; amem[s] != q[r+1] {
                <a id="L2779"></a>error(&#34;clobber of a array, pos&#39;n %v, by %v&#34;, s, q[r+1])
            <a id="L2780"></a>}
            <a id="L2781"></a>amem[s] = q[r+1];
        <a id="L2782"></a>}
        <a id="L2783"></a>indgo[i] = n;
        <a id="L2784"></a>if adb &gt; 1 {
            <a id="L2785"></a>fmt.Fprintf(ftable, &#34;State %v: entry at %v\n&#34;, i, indgo[i])
        <a id="L2786"></a>}
        <a id="L2787"></a>return;
    <a id="L2788"></a>}
    <a id="L2789"></a>error(&#34;Error; failure to place state %v&#34;, i);
<a id="L2790"></a>}

<a id="L2792"></a><span class="comment">//</span>
<a id="L2793"></a><span class="comment">// this version is for limbo</span>
<a id="L2794"></a><span class="comment">// write out the optimized parser</span>
<a id="L2795"></a><span class="comment">//</span>
<a id="L2796"></a>func aoutput() {
    <a id="L2797"></a>fmt.Fprintf(ftable, &#34;const\tYYLAST\t= %v\n&#34;, maxa+1);
    <a id="L2798"></a>arout(&#34;YYACT&#34;, amem, maxa+1);
    <a id="L2799"></a>arout(&#34;YYPACT&#34;, indgo, nstate);
    <a id="L2800"></a>arout(&#34;YYPGO&#34;, pgo, nnonter+1);
<a id="L2801"></a>}

<a id="L2803"></a><span class="comment">//</span>
<a id="L2804"></a><span class="comment">// put out other arrays, copy the parsers</span>
<a id="L2805"></a><span class="comment">//</span>
<a id="L2806"></a>func others() {
    <a id="L2807"></a>var i, j int;

    <a id="L2809"></a>arout(&#34;YYR1&#34;, levprd, nprod);
    <a id="L2810"></a>aryfil(temp1, nprod, 0);

    <a id="L2812"></a><span class="comment">//</span>
    <a id="L2813"></a><span class="comment">//yyr2 is the number of rules for each production</span>
    <a id="L2814"></a><span class="comment">//</span>
    <a id="L2815"></a>for i = 1; i &lt; nprod; i++ {
        <a id="L2816"></a>temp1[i] = len(prdptr[i]) - 2
    <a id="L2817"></a>}
    <a id="L2818"></a>arout(&#34;YYR2&#34;, temp1, nprod);

    <a id="L2820"></a>aryfil(temp1, nstate, -1000);
    <a id="L2821"></a>for i = 0; i &lt;= ntokens; i++ {
        <a id="L2822"></a>for j := tstates[i]; j != 0; j = mstates[j] {
            <a id="L2823"></a>temp1[j] = i
        <a id="L2824"></a>}
    <a id="L2825"></a>}
    <a id="L2826"></a>for i = 0; i &lt;= nnonter; i++ {
        <a id="L2827"></a>for j = ntstates[i]; j != 0; j = mstates[j] {
            <a id="L2828"></a>temp1[j] = -i
        <a id="L2829"></a>}
    <a id="L2830"></a>}
    <a id="L2831"></a>arout(&#34;YYCHK&#34;, temp1, nstate);
    <a id="L2832"></a>arout(&#34;YYDEF&#34;, defact, nstate);

    <a id="L2834"></a><span class="comment">// put out token translation tables</span>
    <a id="L2835"></a><span class="comment">// table 1 has 0-256</span>
    <a id="L2836"></a>aryfil(temp1, 256, 0);
    <a id="L2837"></a>c := 0;
    <a id="L2838"></a>for i = 1; i &lt;= ntokens; i++ {
        <a id="L2839"></a>j = tokset[i].value;
        <a id="L2840"></a>if j &gt;= 0 &amp;&amp; j &lt; 256 {
            <a id="L2841"></a>if temp1[j] != 0 {
                <a id="L2842"></a>print(&#34;yacc bug -- cant have 2 different Ts with same value\n&#34;);
                <a id="L2843"></a>print(&#34;	%s and %s\n&#34;, tokset[i].name, tokset[temp1[j]].name);
                <a id="L2844"></a>nerrors++;
            <a id="L2845"></a>}
            <a id="L2846"></a>temp1[j] = i;
            <a id="L2847"></a>if j &gt; c {
                <a id="L2848"></a>c = j
            <a id="L2849"></a>}
        <a id="L2850"></a>}
    <a id="L2851"></a>}
    <a id="L2852"></a>for i = 0; i &lt;= c; i++ {
        <a id="L2853"></a>if temp1[i] == 0 {
            <a id="L2854"></a>temp1[i] = YYLEXUNK
        <a id="L2855"></a>}
    <a id="L2856"></a>}
    <a id="L2857"></a>arout(&#34;YYTOK1&#34;, temp1, c+1);

    <a id="L2859"></a><span class="comment">// table 2 has PRIVATE-PRIVATE+256</span>
    <a id="L2860"></a>aryfil(temp1, 256, 0);
    <a id="L2861"></a>c = 0;
    <a id="L2862"></a>for i = 1; i &lt;= ntokens; i++ {
        <a id="L2863"></a>j = tokset[i].value - PRIVATE;
        <a id="L2864"></a>if j &gt;= 0 &amp;&amp; j &lt; 256 {
            <a id="L2865"></a>if temp1[j] != 0 {
                <a id="L2866"></a>print(&#34;yacc bug -- cant have 2 different Ts with same value\n&#34;);
                <a id="L2867"></a>print(&#34;	%s and %s\n&#34;, tokset[i].name, tokset[temp1[j]].name);
                <a id="L2868"></a>nerrors++;
            <a id="L2869"></a>}
            <a id="L2870"></a>temp1[j] = i;
            <a id="L2871"></a>if j &gt; c {
                <a id="L2872"></a>c = j
            <a id="L2873"></a>}
        <a id="L2874"></a>}
    <a id="L2875"></a>}
    <a id="L2876"></a>arout(&#34;YYTOK2&#34;, temp1, c+1);

    <a id="L2878"></a><span class="comment">// table 3 has everything else</span>
    <a id="L2879"></a>fmt.Fprintf(ftable, &#34;var\tYYTOK3\t= []int {\n&#34;);
    <a id="L2880"></a>c = 0;
    <a id="L2881"></a>for i = 1; i &lt;= ntokens; i++ {
        <a id="L2882"></a>j = tokset[i].value;
        <a id="L2883"></a>if j &gt;= 0 &amp;&amp; j &lt; 256 {
            <a id="L2884"></a>continue
        <a id="L2885"></a>}
        <a id="L2886"></a>if j &gt;= PRIVATE &amp;&amp; j &lt; 256+PRIVATE {
            <a id="L2887"></a>continue
        <a id="L2888"></a>}

        <a id="L2890"></a>fmt.Fprintf(ftable, &#34;%4d,%4d,&#34;, j, i);
        <a id="L2891"></a>c++;
        <a id="L2892"></a>if c%5 == 0 {
            <a id="L2893"></a>putrune(ftable, &#39;\n&#39;)
        <a id="L2894"></a>}
    <a id="L2895"></a>}
    <a id="L2896"></a>fmt.Fprintf(ftable, &#34;%4d\n };\n&#34;, 0);

    <a id="L2898"></a><span class="comment">// copy parser text</span>
    <a id="L2899"></a>c = getrune(finput);
    <a id="L2900"></a>for c != EOF {
        <a id="L2901"></a>putrune(ftable, c);
        <a id="L2902"></a>c = getrune(finput);
    <a id="L2903"></a>}

    <a id="L2905"></a><span class="comment">// copy yaccpar</span>
    <a id="L2906"></a>fmt.Fprintf(ftable, &#34;%v&#34;, yaccpar);
<a id="L2907"></a>}

<a id="L2909"></a>func arout(s string, v []int, n int) {
    <a id="L2910"></a>fmt.Fprintf(ftable, &#34;var\t%v\t= []int {\n&#34;, s);
    <a id="L2911"></a>for i := 0; i &lt; n; i++ {
        <a id="L2912"></a>if i%10 == 0 {
            <a id="L2913"></a>putrune(ftable, &#39;\n&#39;)
        <a id="L2914"></a>}
        <a id="L2915"></a>fmt.Fprintf(ftable, &#34;%4d&#34;, v[i]);
        <a id="L2916"></a>putrune(ftable, &#39;,&#39;);
    <a id="L2917"></a>}
    <a id="L2918"></a>fmt.Fprintf(ftable, &#34;\n};\n&#34;);
<a id="L2919"></a>}

<a id="L2921"></a><span class="comment">//</span>
<a id="L2922"></a><span class="comment">// output the summary on y.output</span>
<a id="L2923"></a><span class="comment">//</span>
<a id="L2924"></a>func summary() {
    <a id="L2925"></a>if foutput != nil {
        <a id="L2926"></a>fmt.Fprintf(foutput, &#34;\n%v terminals, %v nonterminals\n&#34;, ntokens, nnonter+1);
        <a id="L2927"></a>fmt.Fprintf(foutput, &#34;%v grammar rules, %v/%v states\n&#34;, nprod, nstate, NSTATES);
        <a id="L2928"></a>fmt.Fprintf(foutput, &#34;%v shift/reduce, %v reduce/reduce conflicts reported\n&#34;, zzsrconf, zzrrconf);
        <a id="L2929"></a>fmt.Fprintf(foutput, &#34;%v working sets used\n&#34;, len(wsets));
        <a id="L2930"></a>fmt.Fprintf(foutput, &#34;memory: parser %v/%v\n&#34;, memp, ACTSIZE);
        <a id="L2931"></a>fmt.Fprintf(foutput, &#34;%v extra closures\n&#34;, zzclose-2*nstate);
        <a id="L2932"></a>fmt.Fprintf(foutput, &#34;%v shift entries, %v exceptions\n&#34;, zzacent, zzexcp);
        <a id="L2933"></a>fmt.Fprintf(foutput, &#34;%v goto entries\n&#34;, zzgoent);
        <a id="L2934"></a>fmt.Fprintf(foutput, &#34;%v entries saved by goto default\n&#34;, zzgobest);
    <a id="L2935"></a>}
    <a id="L2936"></a>if zzsrconf != 0 || zzrrconf != 0 {
        <a id="L2937"></a>fmt.Printf(&#34;\nconflicts: &#34;);
        <a id="L2938"></a>if zzsrconf != 0 {
            <a id="L2939"></a>fmt.Printf(&#34;%v shift/reduce&#34;, zzsrconf)
        <a id="L2940"></a>}
        <a id="L2941"></a>if zzsrconf != 0 &amp;&amp; zzrrconf != 0 {
            <a id="L2942"></a>fmt.Printf(&#34;, &#34;)
        <a id="L2943"></a>}
        <a id="L2944"></a>if zzrrconf != 0 {
            <a id="L2945"></a>fmt.Printf(&#34;%v reduce/reduce&#34;, zzrrconf)
        <a id="L2946"></a>}
        <a id="L2947"></a>fmt.Printf(&#34;\n&#34;);
    <a id="L2948"></a>}
<a id="L2949"></a>}

<a id="L2951"></a><span class="comment">//</span>
<a id="L2952"></a><span class="comment">// write optimizer summary</span>
<a id="L2953"></a><span class="comment">//</span>
<a id="L2954"></a>func osummary() {
    <a id="L2955"></a>if foutput == nil {
        <a id="L2956"></a>return
    <a id="L2957"></a>}
    <a id="L2958"></a>i := 0;
    <a id="L2959"></a>for p := maxa; p &gt;= 0; p-- {
        <a id="L2960"></a>if amem[p] == 0 {
            <a id="L2961"></a>i++
        <a id="L2962"></a>}
    <a id="L2963"></a>}

    <a id="L2965"></a>fmt.Fprintf(foutput, &#34;Optimizer space used: output %v/%v\n&#34;, maxa+1, ACTSIZE);
    <a id="L2966"></a>fmt.Fprintf(foutput, &#34;%v table entries, %v zero\n&#34;, maxa+1, i);
    <a id="L2967"></a>fmt.Fprintf(foutput, &#34;maximum spread: %v, maximum offset: %v\n&#34;, maxspr, maxoff);
<a id="L2968"></a>}

<a id="L2970"></a><span class="comment">//</span>
<a id="L2971"></a><span class="comment">// copies and protects &#34;&#39;s in q</span>
<a id="L2972"></a><span class="comment">//</span>
<a id="L2973"></a>func chcopy(q string) string {
    <a id="L2974"></a>s := &#34;&#34;;
    <a id="L2975"></a>i := 0;
    <a id="L2976"></a>j := 0;
    <a id="L2977"></a>for i = 0; i &lt; len(q); i++ {
        <a id="L2978"></a>if q[i] == &#39;&#34;&#39; {
            <a id="L2979"></a>s += q[j:i] + &#34;\\&#34;;
            <a id="L2980"></a>j = i;
        <a id="L2981"></a>}
    <a id="L2982"></a>}
    <a id="L2983"></a>return s + q[j:i];
<a id="L2984"></a>}

<a id="L2986"></a>func usage() {
    <a id="L2987"></a>fmt.Fprintf(stderr, &#34;usage: gacc [-o output] [-v parsetable] input\n&#34;);
    <a id="L2988"></a>exit(1);
<a id="L2989"></a>}

<a id="L2991"></a>func bitset(set Lkset, bit int) int { return set[bit&gt;&gt;5] &amp; (1 &lt;&lt; uint(bit&amp;31)) }

<a id="L2993"></a>func setbit(set Lkset, bit int) { set[bit&gt;&gt;5] |= (1 &lt;&lt; uint(bit&amp;31)) }

<a id="L2995"></a>func mkset() Lkset { return make([]int, tbitset) }

<a id="L2997"></a><span class="comment">//</span>
<a id="L2998"></a><span class="comment">// set a to the union of a and b</span>
<a id="L2999"></a><span class="comment">// return 1 if b is not a subset of a, 0 otherwise</span>
<a id="L3000"></a><span class="comment">//</span>
<a id="L3001"></a>func setunion(a, b []int) int {
    <a id="L3002"></a>sub := 0;
    <a id="L3003"></a>for i := 0; i &lt; tbitset; i++ {
        <a id="L3004"></a>x := a[i];
        <a id="L3005"></a>y := x | b[i];
        <a id="L3006"></a>a[i] = y;
        <a id="L3007"></a>if y != x {
            <a id="L3008"></a>sub = 1
        <a id="L3009"></a>}
    <a id="L3010"></a>}
    <a id="L3011"></a>return sub;
<a id="L3012"></a>}

<a id="L3014"></a>func prlook(p Lkset) {
    <a id="L3015"></a>if p == nil {
        <a id="L3016"></a>fmt.Fprintf(foutput, &#34;\tNULL&#34;);
        <a id="L3017"></a>return;
    <a id="L3018"></a>}
    <a id="L3019"></a>fmt.Fprintf(foutput, &#34; { &#34;);
    <a id="L3020"></a>for j := 0; j &lt;= ntokens; j++ {
        <a id="L3021"></a>if bitset(p, j) != 0 {
            <a id="L3022"></a>fmt.Fprintf(foutput, &#34;%v &#34;, symnam(j))
        <a id="L3023"></a>}
    <a id="L3024"></a>}
    <a id="L3025"></a>fmt.Fprintf(foutput, &#34;}&#34;);
<a id="L3026"></a>}

<a id="L3028"></a><span class="comment">//</span>
<a id="L3029"></a><span class="comment">// utility routines</span>
<a id="L3030"></a><span class="comment">//</span>
<a id="L3031"></a>var peekrune int

<a id="L3033"></a>func isdigit(c int) bool { return c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39; }

<a id="L3035"></a>func isword(c int) bool {
    <a id="L3036"></a>return c &gt;= 0xa0 || (c &gt;= &#39;a&#39; &amp;&amp; c &lt;= &#39;z&#39;) || (c &gt;= &#39;A&#39; &amp;&amp; c &lt;= &#39;Z&#39;)
<a id="L3037"></a>}

<a id="L3039"></a>func mktemp(t string) string { return t }

<a id="L3041"></a><span class="comment">//</span>
<a id="L3042"></a><span class="comment">// return 1 if 2 arrays are equal</span>
<a id="L3043"></a><span class="comment">// return 0 if not equal</span>
<a id="L3044"></a><span class="comment">//</span>
<a id="L3045"></a>func aryeq(a []int, b []int) int {
    <a id="L3046"></a>n := len(a);
    <a id="L3047"></a>if len(b) != n {
        <a id="L3048"></a>return 0
    <a id="L3049"></a>}
    <a id="L3050"></a>for ll := 0; ll &lt; n; ll++ {
        <a id="L3051"></a>if a[ll] != b[ll] {
            <a id="L3052"></a>return 0
        <a id="L3053"></a>}
    <a id="L3054"></a>}
    <a id="L3055"></a>return 1;
<a id="L3056"></a>}

<a id="L3058"></a>func putrune(f *bufio.Writer, c int) {
    <a id="L3059"></a>s := string(c);
    <a id="L3060"></a>for i := 0; i &lt; len(s); i++ {
        <a id="L3061"></a>f.WriteByte(s[i])
    <a id="L3062"></a>}
<a id="L3063"></a>}

<a id="L3065"></a>func getrune(f *bufio.Reader) int {
    <a id="L3066"></a>var r int;

    <a id="L3068"></a>if peekrune != 0 {
        <a id="L3069"></a>if peekrune == EOF {
            <a id="L3070"></a>return EOF
        <a id="L3071"></a>}
        <a id="L3072"></a>r = peekrune;
        <a id="L3073"></a>peekrune = 0;
        <a id="L3074"></a>return r;
    <a id="L3075"></a>}

    <a id="L3077"></a>c, n, err := f.ReadRune();
    <a id="L3078"></a>if n == 0 {
        <a id="L3079"></a>return EOF
    <a id="L3080"></a>}
    <a id="L3081"></a>if err != nil {
        <a id="L3082"></a>error(&#34;read error: %v&#34;, err)
    <a id="L3083"></a>}
    <a id="L3084"></a><span class="comment">//fmt.Printf(&#34;rune = %v n=%v\n&#34;, string(c), n);</span>
    <a id="L3085"></a>return c;
<a id="L3086"></a>}

<a id="L3088"></a>func ungetrune(f *bufio.Reader, c int) {
    <a id="L3089"></a>if f != finput {
        <a id="L3090"></a>panic(&#34;ungetc - not finput&#34;)
    <a id="L3091"></a>}
    <a id="L3092"></a>if peekrune != 0 {
        <a id="L3093"></a>panic(&#34;ungetc - 2nd unget&#34;)
    <a id="L3094"></a>}
    <a id="L3095"></a>peekrune = c;
<a id="L3096"></a>}

<a id="L3098"></a>func write(f *bufio.Writer, b []byte, n int) int {
    <a id="L3099"></a>println(&#34;write&#34;);
    <a id="L3100"></a>return 0;
<a id="L3101"></a>}

<a id="L3103"></a>func open(s string) *bufio.Reader {
    <a id="L3104"></a>fi, err := os.Open(s, os.O_RDONLY, 0);
    <a id="L3105"></a>if err != nil {
        <a id="L3106"></a>error(&#34;error opening %v: %v&#34;, s, err)
    <a id="L3107"></a>}
    <a id="L3108"></a><span class="comment">//fmt.Printf(&#34;open %v\n&#34;, s);</span>
    <a id="L3109"></a>return bufio.NewReader(fi);
<a id="L3110"></a>}

<a id="L3112"></a>func create(s string, m int) *bufio.Writer {
    <a id="L3113"></a>fo, err := os.Open(s, os.O_WRONLY|os.O_CREAT|os.O_TRUNC, m);
    <a id="L3114"></a>if err != nil {
        <a id="L3115"></a>error(&#34;error opening %v: %v&#34;, s, err)
    <a id="L3116"></a>}
    <a id="L3117"></a><span class="comment">//fmt.Printf(&#34;create %v mode %v\n&#34;, s, m);</span>
    <a id="L3118"></a>return bufio.NewWriter(fo);
<a id="L3119"></a>}

<a id="L3121"></a><span class="comment">//</span>
<a id="L3122"></a><span class="comment">// write out error comment</span>
<a id="L3123"></a><span class="comment">//</span>
<a id="L3124"></a>func error(s string, v ...) {
    <a id="L3125"></a>nerrors++;
    <a id="L3126"></a>fmt.Fprintf(stderr, s, v);
    <a id="L3127"></a>fmt.Fprintf(stderr, &#34;: %v:%v\n&#34;, infile, lineno);
    <a id="L3128"></a>if fatfl != 0 {
        <a id="L3129"></a>summary();
        <a id="L3130"></a>exit(1);
    <a id="L3131"></a>}
<a id="L3132"></a>}

<a id="L3134"></a>func exit(status int) {
    <a id="L3135"></a>if ftable != nil {
        <a id="L3136"></a>ftable.Flush();
        <a id="L3137"></a>ftable = nil;
    <a id="L3138"></a>}
    <a id="L3139"></a>if foutput != nil {
        <a id="L3140"></a>foutput.Flush();
        <a id="L3141"></a>foutput = nil;
    <a id="L3142"></a>}
    <a id="L3143"></a>if stderr != nil {
        <a id="L3144"></a>stderr.Flush();
        <a id="L3145"></a>stderr = nil;
    <a id="L3146"></a>}
    <a id="L3147"></a>os.Exit(status);
<a id="L3148"></a>}

<a id="L3150"></a>var yaccpar =
<a id="L3151"></a><span class="comment">// from here to the end of the file is</span>
<a id="L3152"></a><span class="comment">// a single string containing the old yaccpar file</span>
<a id="L3153"></a>`
/*	parser for yacc output	*/

var	Nerrs		= 0		/* number of errors */
var	Errflag		= 0		/* error recovery flag */
var	Debug		= 0
const	YYFLAG		= -1000

func
Tokname(yyc int) string
{
	if yyc &gt; 0 &amp;&amp; yyc &lt;= len(Toknames) {
		if Toknames[yyc-1] != &#34;&#34; {
			return Toknames[yyc-1];
		}
	}
	return fmt.Sprintf(&#34;tok-%v&#34;, yyc);
}

func
Statname(yys int) string
{
	if yys &gt;= 0 &amp;&amp; yys &lt; len(Statenames) {
		if Statenames[yys] != &#34;&#34; {
			return Statenames[yys];
		}
	}
	return fmt.Sprintf(&#34;state-%v&#34;, yys);
}

func
lex1() int
{
	var yychar int;
	var c int;

	yychar = Lex();
	if yychar &lt;= 0 {
		c = YYTOK1[0];
		goto out;
	}
	if yychar &lt; len(YYTOK1) {
		c = YYTOK1[yychar];
		goto out;
	}
	if yychar &gt;= YYPRIVATE {
		if yychar &lt; YYPRIVATE+len(YYTOK2) {
			c = YYTOK2[yychar-YYPRIVATE];
			goto out;
		}
	}
	for i:=0; i&lt;len(YYTOK3); i+=2 {
		c = YYTOK3[i+0];
		if c == yychar {
			c = YYTOK3[i+1];
			goto out;
		}
	}
	c = 0;

out:
	if c == 0 {
		c = YYTOK2[1];	/* unknown char */
	}
	if Debug &gt;= 3 {
		fmt.Printf(&#34;lex %.4lux %s\n&#34;, yychar, Tokname(c));
	}
	return c;
}

func
Parse() int
{
	var yyj, yystate, yyn, yyg, yyxi, yyp int;
	var yychar int;
	var yypt, yynt int;

	yystate = 0;
	yychar = -1;
	Nerrs = 0;
	Errflag = 0;
	yyp = -1;
	goto yystack;

ret0:
	return 0;

ret1:
	return 1;

yystack:
	/* put a state and value onto the stack */
	if Debug &gt;= 4 {
		fmt.Printf(&#34;char %v in %v&#34;, Tokname(yychar), Statname(yystate));
	}

	yyp++;
	if yyp &gt;= len(YYS) {
		Error(&#34;yacc stack overflow&#34;);
		goto ret1;
	}
	YYS[yyp] = YYVAL;
	YYS[yyp].yys = yystate;

yynewstate:
	yyn = YYPACT[yystate];
	if yyn &lt;= YYFLAG {
		goto yydefault; /* simple state */
	}
	if yychar &lt; 0 {
		yychar = lex1();
	}
	yyn += yychar;
	if yyn &lt; 0 || yyn &gt;= YYLAST {
		goto yydefault;
	}
	yyn = YYACT[yyn];
	if YYCHK[yyn] == yychar { /* valid shift */
		yychar = -1;
		YYVAL = yylval;
		yystate = yyn;
		if Errflag &gt; 0 {
			Errflag--;
		}
		goto yystack;
	}

yydefault:
	/* default state action */
	yyn = YYDEF[yystate];
	if yyn == -2 {
		if yychar &lt; 0 {
			yychar = lex1();
		}

		/* look through exception table */
		for yyxi=0;; yyxi+=2 {
			if YYEXCA[yyxi+0] == -1 &amp;&amp; YYEXCA[yyxi+1] == yystate {
				break;
			}
		}
		for yyxi += 2;; yyxi += 2 {
			yyn = YYEXCA[yyxi+0];
			if yyn &lt; 0 || yyn == yychar {
				break;
			}
		}
		yyn = YYEXCA[yyxi+1];
		if yyn &lt; 0 {
			goto ret0;
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0:   /* brand new error */
			Error(&#34;syntax error&#34;);
			Nerrs++;
			if Debug &gt;= 1 {
				fmt.Printf(&#34;%s&#34;, Statname(yystate));
				fmt.Printf(&#34;saw %s\n&#34;, Tokname(yychar));
			}
			fallthrough;

		case 1,2: /* incompletely recovered error ... try again */
			Errflag = 3;

			/* find a state where &#34;error&#34; is a legal shift action */
			for yyp &gt;= len(YYS) {
				yyn = YYPACT[YYS[yyp].yys] + YYERRCODE;
				if yyn &gt;= 0 &amp;&amp; yyn &lt; YYLAST {
					yystate = YYACT[yyn];  /* simulate a shift of &#34;error&#34; */
					if YYCHK[yystate] == YYERRCODE {
						goto yystack;
					}
				}

				/* the current yyp has no shift onn &#34;error&#34;, pop stack */
				if Debug &gt;= 2 {
					fmt.Printf(&#34;error recovery pops state %d, uncovers %d\n&#34;,
						YYS[yyp].yys, YYS[yyp-1].yys );
				}
				yyp--;
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1;

		case 3:  /* no shift yet; clobber input char */
			if Debug &gt;= 2 {
				fmt.Printf(&#34;error recovery discards %s\n&#34;, Tokname(yychar));
			}
			if yychar == YYEOFCODE {
				goto ret1;
			}
			yychar = -1;
			goto yynewstate;   /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if Debug &gt;= 2 {
		fmt.Printf(&#34;reduce %v in:\n\t%v&#34;, yyn, Statname(yystate));
	}

	yynt = yyn;
	yypt = yyp;

	yyp -= YYR2[yyn];
	YYVAL = YYS[yyp+1];

	/* consult goto table to find next state */
	yyn = YYR1[yyn];
	yyg = YYPGO[yyn];
	yyj = yyg + YYS[yyp].yys + 1;

	if yyj &gt;= YYLAST {
		yystate = YYACT[yyg];
	} else {
		yystate = YYACT[yyj];
		if YYCHK[yystate] != -yyn {
			yystate = YYACT[yyg];
		}
	}

	yyrun(yynt, yypt);
	goto yystack;  /* stack new state and value */
}
`
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

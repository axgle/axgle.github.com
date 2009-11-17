<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/goyacc/units.y</title>

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
  <h1 id="generatedHeader">Text file src/cmd/goyacc/units.y</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Derived from Plan 9&#39;s /sys/src/cmd/units.y
// http://plan9.bell-labs.com/sources/plan9/sys/src/cmd/units.y
//
// Copyright (C) 2003, Lucent Technologies Inc. and others. All Rights Reserved.
// Portions Copyright 2009 The Go Authors.  All Rights Reserved.
// Distributed under the terms of the Lucent Public License Version 1.02
// See http://plan9.bell-labs.com/plan9/license.html

%{

// units.y
// example of a goyacc program
// usage is
//	goyacc units.y (produces y.go)
//	6g y.go
//	6l y.6
//	./6.out $GOROOT/src/cmd/goyacc/units
//	you have: c
//	you want: furlongs/fortnight
//		* 1.8026178e+12
//		/ 5.5474878e-13
//	you have:

package main

import
(
	&#34;flag&#34;;
	&#34;fmt&#34;;
	&#34;bufio&#34;;
	&#34;os&#34;;
	&#34;math&#34;;
	&#34;strconv&#34;;
	&#34;utf8&#34;;
)

const
(
	Ndim	= 15;				// number of dimensions
	Maxe	= 695;				// log of largest number
)

type	Node
struct
{
	vval	float64;
	dim	[Ndim]int8;
}

type	Var
struct
{
	name	string;
	node	Node;
}

var	fi		*bufio.Reader		// input
var	fund		[Ndim]*Var		// names of fundamental units
var	line		string			// current input line
var	lineno		int			// current input line number
var	linep		int			// index to next rune in unput
var	nerrors		int			// error count
var	one		Node			// constant one
var	peekrune	int			// backup runt from input
var	retnode1	Node
var	retnode2	Node
var	retnode		Node
var	sym		string
var	vflag		bool

%}

%union
{
	node	Node;
	vvar	*Var;
	numb	int;
	vval	float64;
}

%type	&lt;node&gt;	prog expr expr0 expr1 expr2 expr3 expr4

%token	&lt;vval&gt;	VAL
%token	&lt;vvar&gt;	VAR
%token	&lt;numb&gt;	SUP
%%
prog:
	&#39;:&#39; VAR expr
	{
		var f int;

		f = int($2.node.dim[0]);
		$2.node = $3;
		$2.node.dim[0] = 1;
		if f != 0 {
			Error(&#34;redefinition of %v&#34;, $2.name);
		} else
		if vflag {
			fmt.Printf(&#34;%v\t%v\n&#34;, $2.name, &amp;$2.node);
		}
	}
|	&#39;:&#39; VAR &#39;#&#39;
	{
		var f, i int;

		for i=1; i&lt;Ndim; i++ {
			if fund[i] == nil {
				break;
			}
		}
		if i &gt;= Ndim {
			Error(&#34;too many dimensions&#34;);
			i = Ndim-1;
		}
		fund[i] = $2;

		f = int($2.node.dim[0]);
		$2.node = one;
		$2.node.dim[0] = 1;
		$2.node.dim[i] = 1;
		if f != 0 {
			Error(&#34;redefinition of %v&#34;, $2.name);
		} else
		if vflag {
			fmt.Printf(&#34;%v\t#\n&#34;, $2.name);
		}
	}
|	&#39;:&#39;
	{
	}
|	&#39;?&#39; expr
	{
		retnode1 = $2;
	}
|	&#39;?&#39;
	{
		retnode1 = one;
	}

expr:
	expr4
|	expr &#39;+&#39; expr4
	{
		add(&amp;$$, &amp;$1, &amp;$3);
	}
|	expr &#39;-&#39; expr4
	{
		sub(&amp;$$, &amp;$1, &amp;$3);
	}

expr4:
	expr3
|	expr4 &#39;*&#39; expr3
	{
		mul(&amp;$$, &amp;$1, &amp;$3);
	}
|	expr4 &#39;/&#39; expr3
	{
		div(&amp;$$, &amp;$1, &amp;$3);
	}

expr3:
	expr2
|	expr3 expr2
	{
		mul(&amp;$$, &amp;$1, &amp;$2);
	}

expr2:
	expr1
|	expr2 SUP
	{
		xpn(&amp;$$, &amp;$1, $2);
	}
|	expr2 &#39;^&#39; expr1
	{
		var i int;

		for i=1; i&lt;Ndim; i++ {
			if $3.dim[i] != 0 {
				Error(&#34;exponent has units&#34;);
				$$ = $1;
				break;
			}
		}
		if i &gt;= Ndim {
			i = int($3.vval);
			if float64(i) != $3.vval {
				Error(&#34;exponent not integral&#34;);
			}
			xpn(&amp;$$, &amp;$1, i);
		}
	}

expr1:
	expr0
|	expr1 &#39;|&#39; expr0
	{
		div(&amp;$$, &amp;$1, &amp;$3);
	}

expr0:
	VAR
	{
		if $1.node.dim[0] == 0 {
			Error(&#34;undefined %v&#34;, $1.name);
			$$ = one;
		} else
			$$ = $1.node;
	}
|	VAL
	{
		$$ = one;
		$$.vval = $1;
	}
|	&#39;(&#39; expr &#39;)&#39;
	{
		$$ = $2;
	}
%%

func
Lex() int
{
	var c, i int;

	c = peekrune;
	peekrune = &#39; &#39;;

loop:
	if (c &gt;= &#39;0&#39; &amp;&amp; c &lt;= &#39;9&#39;) || c == &#39;.&#39; {
		goto numb;
	}
	if ralpha(c) {
		goto alpha;
	}
	switch c {
	case &#39; &#39;, &#39;\t&#39;:
		c = getrune();
		goto loop;
	case &#39;×&#39;:
		return &#39;*&#39;;
	case &#39;÷&#39;:
		return &#39;/&#39;;
	case &#39;¹&#39;, &#39;ⁱ&#39;:
		yylval.numb = 1;
		return SUP;
	case &#39;²&#39;, &#39;⁲&#39;:
		yylval.numb = 2;
		return SUP;
	case &#39;³&#39;, &#39;⁳&#39;:
		yylval.numb = 3;
		return SUP;
	}
	return c;

alpha:
	sym = &#34;&#34;;
	for i=0;; i++ {
		sym += string(c);
		c = getrune();
		if !ralpha(c) {
			break;
		}
	}
	peekrune = c;
	yylval.vvar = lookup(0);
	return VAR;

numb:
	sym = &#34;&#34;;
	for i=0;; i++ {
		sym += string(c);
		c = getrune();
		if !rdigit(c) {
			break;
		}
	}
	peekrune = c;
	f, err := strconv.Atof64(sym);
	if err != nil {
		fmt.Printf(&#34;error converting %v&#34;, sym);
		f = 0;
	}
	yylval.vval = f;
	return VAL;
}

func
main()
{
	var file string;

	flag.BoolVar(&amp;vflag, &#34;v&#34;, false, &#34;verbose&#34;);

	flag.Parse();

	file = os.Getenv(&#34;GOROOT&#34;) + &#34;/src/cmd/goyacc/units.txt&#34;;
	if flag.NArg() &gt; 0 {
		file = flag.Arg(0);
	}

	f,err := os.Open(file, os.O_RDONLY, 0);
	if err != nil {
		fmt.Printf(&#34;error opening %v: %v&#34;, file, err);
		os.Exit(1);
	}
	fi = bufio.NewReader(f);

	one.vval = 1;

	/*
	 * read the &#39;units&#39; file to
	 * develope a database
	 */
	lineno = 0;
	for {
		lineno++;
		if readline() {
			break;
		}
		if len(line) == 0 || line[0] == &#39;/&#39; {
			continue;
		}
		peekrune = &#39;:&#39;;
		Parse();
	}

	/*
	 * read the console to
	 * print ratio of pairs
	 */
	fi = bufio.NewReader(os.NewFile(0, &#34;stdin&#34;));

	lineno = 0;
	for {
		if (lineno &amp; 1) != 0 {
			fmt.Printf(&#34;you want: &#34;);
		} else
			fmt.Printf(&#34;you have: &#34;);
		if readline() {
			break;
		}
		peekrune = &#39;?&#39;;
		nerrors = 0;
		Parse();
		if nerrors != 0 {
			continue;
		}
		if (lineno &amp; 1) != 0 {
			if specialcase(&amp;retnode, &amp;retnode2, &amp;retnode1) {
				fmt.Printf(&#34;\tis %v\n&#34;, &amp;retnode);
			} else {
				div(&amp;retnode, &amp;retnode2, &amp;retnode1);
				fmt.Printf(&#34;\t* %v\n&#34;, &amp;retnode);
				div(&amp;retnode, &amp;retnode1, &amp;retnode2);
				fmt.Printf(&#34;\t/ %v\n&#34;, &amp;retnode);
			}
		} else
			retnode2 = retnode1;
		lineno++;
	}
	fmt.Printf(&#34;\n&#34;);
	os.Exit(0);
}

/*
 * all characters that have some
 * meaning. rest are usable as names
 */
func
ralpha(c int) bool
{
	switch c {
	case	0, &#39;+&#39;, &#39;-&#39;, &#39;*&#39;, &#39;/&#39;, &#39;[&#39;, &#39;]&#39;, &#39;(&#39;, &#39;)&#39;,
		&#39;^&#39;, &#39;:&#39;, &#39;?&#39;, &#39; &#39;, &#39;\t&#39;, &#39;.&#39;, &#39;|&#39;, &#39;#&#39;,
		&#39;×&#39;, &#39;÷&#39;, &#39;¹&#39;, &#39;ⁱ&#39;, &#39;²&#39;, &#39;⁲&#39;, &#39;³&#39;, &#39;⁳&#39;:
			return false;
	}
	return true;
}

/*
 * number forming character
 */
func
rdigit(c int) bool
{
	switch c {
	case	&#39;0&#39;, &#39;1&#39;, &#39;2&#39;, &#39;3&#39;, &#39;4&#39;, &#39;5&#39;, &#39;6&#39;, &#39;7&#39;, &#39;8&#39;, &#39;9&#39;,
		&#39;.&#39;, &#39;e&#39;, &#39;+&#39;, &#39;-&#39;:
		return true;
	}
	return false;
}

func
Error(s string, v ...)
{

	/*
	 * hack to intercept message from yaccpar
	 */
	if s == &#34;syntax error&#34; {
		Error(&#34;syntax error, last name: %v&#34;, sym);
		return;
	}
	fmt.Printf(&#34;%v: %v\n\t&#34;, lineno, line);
	fmt.Printf(s, v);
	fmt.Printf(&#34;\n&#34;);

	nerrors++;
	if nerrors &gt; 5 {
		fmt.Printf(&#34;too many errors\n&#34;);
		os.Exit(1);
	}
}

func
add(c,a,b *Node)
{
	var i int;
	var d int8;

	for i=0; i&lt;Ndim; i++ {
		d = a.dim[i];
		c.dim[i] = d;
		if d != b.dim[i] {
			Error(&#34;add must be like units&#34;);
		}
	}
	c.vval = fadd(a.vval, b.vval);
}

func
sub(c,a,b *Node)
{
	var i int;
	var d int8;

	for i=0; i&lt;Ndim; i++ {
		d = a.dim[i];
		c.dim[i] = d;
		if d != b.dim[i] {
			Error(&#34;sub must be like units&#34;);
		}
	}
	c.vval = fadd(a.vval, -b.vval);
}

func
mul(c,a,b *Node)
{
	var i int;

	for i=0; i&lt;Ndim; i++ {
		c.dim[i] = a.dim[i] + b.dim[i];
	}
	c.vval = fmul(a.vval, b.vval);
}

func
div(c,a,b *Node)
{
	var i int;

	for i=0; i&lt;Ndim; i++ {
		c.dim[i] = a.dim[i] - b.dim[i];
	}
	c.vval = fdiv(a.vval, b.vval);
}

func
xpn(c,a *Node, b int)
{
	var i int;

	*c = one;
	if b &lt; 0 {
		b = -b;
		for i=0; i&lt;b; i++ {
			div(c, c, a);
		}
	} else
	for i=0; i&lt;b; i++ {
		mul(c, c, a);
	}
}

func
specialcase(c,a,b *Node) bool
{
	var i int;
	var d, d1, d2 int8;

	d1 = 0;
	d2 = 0;
	for i=1; i&lt;Ndim; i++ {
		d = a.dim[i];
		if d != 0 {
			if d != 1 || d1 != 0 {
				return false;
			}
			d1 = int8(i);
		}
		d = b.dim[i];
		if d != 0 {
			if d != 1 || d2 != 0 {
				return false;
			}
			d2 = int8(i);
		}
	}
	if d1 == 0 || d2 == 0 {
		return false;
	}

	if fund[d1].name == &#34;°C&#34; &amp;&amp; fund[d2].name == &#34;°F&#34; &amp;&amp;
	   b.vval == 1 {
		for ll:=0; ll&lt;len(c.dim); ll++ {
			c.dim[ll] = b.dim[ll];
		}
		c.vval = a.vval * 9. / 5. + 32.;
		return true;
	}

	if fund[d1].name == &#34;°F&#34; &amp;&amp; fund[d2].name == &#34;°C&#34; &amp;&amp;
	   b.vval == 1 {
		for ll:=0; ll&lt;len(c.dim); ll++ {
			c.dim[ll] = b.dim[ll];
		}
		c.vval = (a.vval - 32.) * 5. / 9.;
		return true;
	}
	return false;
}

func
printdim(str string, d, n int) string
{
	var v *Var;

	if n != 0 {
		v = fund[d];
		if v != nil {
			str += fmt.Sprintf(&#34;%v&#34;, v.name);
		} else
			str += fmt.Sprintf(&#34;[%d]&#34;, d);
		switch n {
		case 1:
			break;
		case 2:
			str += &#34;²&#34;;
		case 3:
			str += &#34;³&#34;;
		default:
			str += fmt.Sprintf(&#34;^%d&#34;, n);
		}
	}
	return str;
}

func (n Node)
String() string
{
	var str string;
	var f, i, d int;

	str = fmt.Sprintf(&#34;%.7e &#34;, n.vval);

	f = 0;
	for i=1; i&lt;Ndim; i++ {
		d = int(n.dim[i]);
		if d &gt; 0 {
			str = printdim(str, i, d);
		} else
		if d &lt; 0 {
			f = 1;
		}
	}

	if f != 0 {
		str += &#34; /&#34;;
		for i=1; i&lt;Ndim; i++ {
			d = int(n.dim[i]);
			if d &lt; 0 {
				str = printdim(str, i, -d);
			}
		}
	}

	return str;
}

func (v *Var)
String() string
{
	var str string;
	str = fmt.Sprintf(&#34;%v %v&#34;, v.name, v.node);
	return str;
}

func
readline() bool
{
	s,err := fi.ReadString(&#39;\n&#39;);
	if err != nil {
		return true;
	}
	line = s;
	linep = 0;
	return false;
}

func
getrune() int
{
	var c,n int;

	if linep &gt;= len(line) {
		return 0;
	}
	c,n = utf8.DecodeRuneInString(line[linep:len(line)]);
	linep += n;
	if c == &#39;\n&#39; {
		c = 0;
	}
	return c;
}

var	symmap	= make(map[string]*Var);	// symbol table

func
lookup(f int) *Var
{
	var p float64;
	var w *Var;

	v,ok := symmap[sym];
	if ok {
		return v;
	}
	if f != 0 {
		return nil;
	}
	v = new(Var);
	v.name = sym;
	symmap[sym] = v;

	p = 1;
	for {
		p = fmul(p, pname());
		if p == 0 {
			break;
		}
		w = lookup(1);
		if w != nil {
			v.node = w.node;
			v.node.vval = fmul(v.node.vval, p);
			break;
		}
	}
	return v;
}

type	Prefix
struct
{
	vval	float64;
	name	string;
}

var	prefix	 = []Prefix {			// prefix table
	Prefix { 1e-24,		&#34;yocto&#34; },
	Prefix { 1e-21,		&#34;zepto&#34; },
	Prefix { 1e-18,		&#34;atto&#34;  },
	Prefix { 1e-15,		&#34;femto&#34; },
	Prefix { 1e-12,		&#34;pico&#34;  },
	Prefix { 1e-9,		&#34;nano&#34;  },
	Prefix { 1e-6,		&#34;micro&#34; },
	Prefix { 1e-6,		&#34;μ&#34;     },
	Prefix { 1e-3,		&#34;milli&#34; },
	Prefix { 1e-2,		&#34;centi&#34; },
	Prefix { 1e-1,		&#34;deci&#34;  },
	Prefix { 1e1,		&#34;deka&#34;  },
	Prefix { 1e2,		&#34;hecta&#34; },
	Prefix { 1e2,		&#34;hecto&#34; },
	Prefix { 1e3,		&#34;kilo&#34;  },
	Prefix { 1e6,		&#34;mega&#34;  },
	Prefix { 1e6,		&#34;meg&#34;   },
	Prefix { 1e9,		&#34;giga&#34;  },
	Prefix { 1e12,		&#34;tera&#34;  },
	Prefix { 1e15,		&#34;peta&#34;  },
	Prefix { 1e18,		&#34;exa&#34;   },
	Prefix { 1e21,		&#34;zetta&#34; },
	Prefix { 1e24,		&#34;yotta&#34; }
}

func
pname() float64
{
	var i, j, n int;
	var s string;

	/*
	 * rip off normal prefixs
	 */
	n = len(sym);
	for i=0; i&lt;len(prefix); i++ {
		s = prefix[i].name;
		j = len(s);
		if j &lt; n &amp;&amp; sym[0:j] == s {
			sym = sym[j:n];
			return prefix[i].vval;
		}
	}

	/*
	 * rip off &#39;s&#39; suffixes
	 */
	if n &gt; 2 &amp;&amp; sym[n-1] == &#39;s&#39; {
		sym = sym[0:n-1];
		return 1;
	}

	return 0;
}


// careful multiplication
// exponents (log) are checked before multiply
func
fmul(a, b float64) float64
{
	var l float64;

	if b &lt;= 0 {
		if b == 0 {
			return 0;
		}
		l = math.Log(-b);
	} else
		l = math.Log(b);

	if a &lt;= 0 {
		if a == 0 {
			return 0;
		}
		l += math.Log(-a);
	} else
		l += math.Log(a);

	if l &gt; Maxe {
		Error(&#34;overflow in multiply&#34;);
		return 1;
	}
	if l &lt; -Maxe {
		Error(&#34;underflow in multiply&#34;);
		return 0;
	}
	return a*b;
}

// careful division
// exponents (log) are checked before divide
func
fdiv(a, b float64) float64
{
	var l float64;

	if b &lt;= 0 {
		if b == 0 {
			Error(&#34;division by zero: %v %v&#34;, a, b);
			return 1;
		}
		l = math.Log(-b);
	} else
		l = math.Log(b);

	if a &lt;= 0 {
		if a == 0 {
			return 0;
		}
		l -= math.Log(-a);
	} else
		l -= math.Log(a);

	if l &lt; -Maxe {
		Error(&#34;overflow in divide&#34;);
		return 1;
	}
	if l &gt; Maxe {
		Error(&#34;underflow in divide&#34;);
		return 0;
	}
	return a/b;
}

func
fadd(a, b float64) float64
{
	return a + b;
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

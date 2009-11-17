<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/cmd/cov/tree.c</title>

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
	<li>Sun Nov 15 20:28:13 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/cmd/cov/tree.c</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Renamed from Map to Tree to avoid conflict with libmach.

/*
Copyright (c) 2003-2007 Russ Cox, Tom Bergan, Austin Clements,
                        Massachusetts Institute of Technology
Portions Copyright (c) 2009 The Go Authors. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
&#34;Software&#34;), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED &#34;AS IS&#34;, WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

// Mutable map structure, but still based on
// Okasaki, Red Black Trees in a Functional Setting, JFP 1999,
// which is a lot easier than the traditional red-black
// and plenty fast enough for me.  (Also I could copy
// and edit fmap.c.)

#include &lt;u.h&gt;
#include &lt;libc.h&gt;
#include &#34;tree.h&#34;

#define TreeNode TreeNode
#define Tree Tree

enum
{
	Red = 0,
	Black = 1
};


// Red-black trees are binary trees with this property:
//	1. No red node has a red parent.
//	2. Every path from the root to a leaf contains the
//		same number of black nodes.

static TreeNode*
rwTreeNode(TreeNode *p, int color, TreeNode *left, void *key, void *value, TreeNode *right)
{
	if(p == nil)
		p = malloc(sizeof *p);
	p-&gt;color = color;
	p-&gt;left = left;
	p-&gt;key = key;
	p-&gt;value = value;
	p-&gt;right = right;
	return p;
}

static TreeNode*
balance(TreeNode *m0)
{
	void *xk, *xv, *yk, *yv, *zk, *zv;
	TreeNode *a, *b, *c, *d;
	TreeNode *m1, *m2;
	int color;
	TreeNode *left, *right;
	void *key, *value;

	color = m0-&gt;color;
	left = m0-&gt;left;
	key = m0-&gt;key;
	value = m0-&gt;value;
	right = m0-&gt;right;

	// Okasaki notation: (T is mkTreeNode, B is Black, R is Red, x, y, z are key-value.
	//
	// balance B (T R (T R a x b) y c) z d
	// balance B (T R a x (T R b y c)) z d
	// balance B a x (T R (T R b y c) z d)
	// balance B a x (T R b y (T R c z d))
	//
	//     = T R (T B a x b) y (T B c z d)

	if(color == Black){
		if(left &amp;&amp; left-&gt;color == Red){
			if(left-&gt;left &amp;&amp; left-&gt;left-&gt;color == Red){
				a = left-&gt;left-&gt;left;
				xk = left-&gt;left-&gt;key;
				xv = left-&gt;left-&gt;value;
				b = left-&gt;left-&gt;right;
				yk = left-&gt;key;
				yv = left-&gt;value;
				c = left-&gt;right;
				zk = key;
				zv = value;
				d = right;
				m1 = left;
				m2 = left-&gt;left;
				goto hard;
			}else if(left-&gt;right &amp;&amp; left-&gt;right-&gt;color == Red){
				a = left-&gt;left;
				xk = left-&gt;key;
				xv = left-&gt;value;
				b = left-&gt;right-&gt;left;
				yk = left-&gt;right-&gt;key;
				yv = left-&gt;right-&gt;value;
				c = left-&gt;right-&gt;right;
				zk = key;
				zv = value;
				d = right;
				m1 = left;
				m2 = left-&gt;right;
				goto hard;
			}
		}else if(right &amp;&amp; right-&gt;color == Red){
			if(right-&gt;left &amp;&amp; right-&gt;left-&gt;color == Red){
				a = left;
				xk = key;
				xv = value;
				b = right-&gt;left-&gt;left;
				yk = right-&gt;left-&gt;key;
				yv = right-&gt;left-&gt;value;
				c = right-&gt;left-&gt;right;
				zk = right-&gt;key;
				zv = right-&gt;value;
				d = right-&gt;right;
				m1 = right;
				m2 = right-&gt;left;
				goto hard;
			}else if(right-&gt;right &amp;&amp; right-&gt;right-&gt;color == Red){
				a = left;
				xk = key;
				xv = value;
				b = right-&gt;left;
				yk = right-&gt;key;
				yv = right-&gt;value;
				c = right-&gt;right-&gt;left;
				zk = right-&gt;right-&gt;key;
				zv = right-&gt;right-&gt;value;
				d = right-&gt;right-&gt;right;
				m1 = right;
				m2 = right-&gt;right;
				goto hard;
			}
		}
	}
	return rwTreeNode(m0, color, left, key, value, right);

hard:
	return rwTreeNode(m0, Red, rwTreeNode(m1, Black, a, xk, xv, b),
		yk, yv, rwTreeNode(m2, Black, c, zk, zv, d));
}

static TreeNode*
ins0(TreeNode *p, void *k, void *v, TreeNode *rw)
{
	if(p == nil)
		return rwTreeNode(rw, Red, nil, k, v, nil);
	if(p-&gt;key == k){
		if(rw)
			return rwTreeNode(rw, p-&gt;color, p-&gt;left, k, v, p-&gt;right);
		p-&gt;value = v;
		return p;
	}
	if(p-&gt;key &lt; k)
		p-&gt;left = ins0(p-&gt;left, k, v, rw);
	else
		p-&gt;right = ins0(p-&gt;right, k, v, rw);
	return balance(p);
}

static TreeNode*
ins1(Tree *m, TreeNode *p, void *k, void *v, TreeNode *rw)
{
	int i;

	if(p == nil)
		return rwTreeNode(rw, Red, nil, k, v, nil);
	i = m-&gt;cmp(p-&gt;key, k);
	if(i == 0){
		if(rw)
			return rwTreeNode(rw, p-&gt;color, p-&gt;left, k, v, p-&gt;right);
		p-&gt;value = v;
		return p;
	}
	if(i &lt; 0)
		p-&gt;left = ins1(m, p-&gt;left, k, v, rw);
	else
		p-&gt;right = ins1(m, p-&gt;right, k, v, rw);
	return balance(p);
}

void
treeputelem(Tree *m, void *key, void *val, TreeNode *rw)
{
	if(m-&gt;cmp)
		m-&gt;root = ins1(m, m-&gt;root, key, val, rw);
	else
		m-&gt;root = ins0(m-&gt;root, key, val, rw);
}

void
treeput(Tree *m, void *key, void *val)
{
	treeputelem(m, key, val, nil);
}

void*
treeget(Tree *m, void *key)
{
	int i;
	TreeNode *p;

	p = m-&gt;root;
	if(m-&gt;cmp){
		for(;;){
			if(p == nil)
				return nil;
			i = m-&gt;cmp(p-&gt;key, key);
			if(i &lt; 0)
				p = p-&gt;left;
			else if(i &gt; 0)
				p = p-&gt;right;
			else
				return p-&gt;value;
		}
	}else{
		for(;;){
			if(p == nil)
				return nil;
			if(p-&gt;key == key)
				return p-&gt;value;
			if(p-&gt;key &lt; key)
				p = p-&gt;left;
			else
				p = p-&gt;right;
		}
	}
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

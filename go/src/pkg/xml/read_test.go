<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/xml/read_test.go</title>

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
  <h1 id="generatedHeader">Source file /src/pkg/xml/read_test.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a>package xml

<a id="L7"></a>import (
    <a id="L8"></a>&#34;reflect&#34;;
    <a id="L9"></a>&#34;testing&#34;;
<a id="L10"></a>)

<a id="L12"></a><span class="comment">// Stripped down Atom feed data structures.</span>

<a id="L14"></a>func TestUnmarshalFeed(t *testing.T) {
    <a id="L15"></a>var f Feed;
    <a id="L16"></a>if err := Unmarshal(StringReader(rssFeedString), &amp;f); err != nil {
        <a id="L17"></a>t.Fatalf(&#34;Unmarshal: %s&#34;, err)
    <a id="L18"></a>}
    <a id="L19"></a>if !reflect.DeepEqual(f, rssFeed) {
        <a id="L20"></a>t.Fatalf(&#34;have %#v\nwant %#v\n\n%#v&#34;, f)
    <a id="L21"></a>}
<a id="L22"></a>}

<a id="L24"></a><span class="comment">// hget http://codereview.appspot.com/rss/mine/rsc</span>
<a id="L25"></a>const rssFeedString = `
&lt;?xml version=&#34;1.0&#34; encoding=&#34;utf-8&#34;?&gt;
&lt;feed xmlns=&#34;http://www.w3.org/2005/Atom&#34; xml:lang=&#34;en-us&#34;&gt;&lt;title&gt;Code Review - My issues&lt;/title&gt;&lt;link href=&#34;http://codereview.appspot.com/&#34; rel=&#34;alternate&#34;&gt;&lt;/link&gt;&lt;link href=&#34;http://codereview.appspot.com/rss/mine/rsc&#34; rel=&#34;self&#34;&gt;&lt;/link&gt;&lt;id&gt;http://codereview.appspot.com/&lt;/id&gt;&lt;updated&gt;2009-10-04T01:35:58+00:00&lt;/updated&gt;&lt;author&gt;&lt;name&gt;rietveld&lt;/name&gt;&lt;/author&gt;&lt;entry&gt;&lt;title&gt;rietveld: an attempt at pubsubhubbub
&lt;/title&gt;&lt;link href=&#34;http://codereview.appspot.com/126085&#34; rel=&#34;alternate&#34;&gt;&lt;/link&gt;&lt;updated&gt;2009-10-04T01:35:58+00:00&lt;/updated&gt;&lt;author&gt;&lt;name&gt;email-address-removed&lt;/name&gt;&lt;/author&gt;&lt;id&gt;urn:md5:134d9179c41f806be79b3a5f7877d19a&lt;/id&gt;&lt;summary type=&#34;html&#34;&gt;
  An attempt at adding pubsubhubbub support to Rietveld.
http://code.google.com/p/pubsubhubbub
http://code.google.com/p/rietveld/issues/detail?id=155

The server side of the protocol is trivial:
  1. add a &amp;amp;lt;link rel=&amp;amp;quot;hub&amp;amp;quot; href=&amp;amp;quot;hub-server&amp;amp;quot;&amp;amp;gt; tag to all
     feeds that will be pubsubhubbubbed.
  2. every time one of those feeds changes, tell the hub
     with a simple POST request.

I have tested this by adding debug prints to a local hub
server and checking that the server got the right publish
requests.

I can&amp;amp;#39;t quite get the server to work, but I think the bug
is not in my code.  I think that the server expects to be
able to grab the feed and see the feed&amp;amp;#39;s actual URL in
the link rel=&amp;amp;quot;self&amp;amp;quot;, but the default value for that drops
the :port from the URL, and I cannot for the life of me
figure out how to get the Atom generator deep inside
django not to do that, or even where it is doing that,
or even what code is running to generate the Atom feed.
(I thought I knew but I added some assert False statements
and it kept running!)

Ignoring that particular problem, I would appreciate
feedback on the right way to get the two values at
the top of feeds.py marked NOTE(rsc).


&lt;/summary&gt;&lt;/entry&gt;&lt;entry&gt;&lt;title&gt;rietveld: correct tab handling
&lt;/title&gt;&lt;link href=&#34;http://codereview.appspot.com/124106&#34; rel=&#34;alternate&#34;&gt;&lt;/link&gt;&lt;updated&gt;2009-10-03T23:02:17+00:00&lt;/updated&gt;&lt;author&gt;&lt;name&gt;email-address-removed&lt;/name&gt;&lt;/author&gt;&lt;id&gt;urn:md5:0a2a4f19bb815101f0ba2904aed7c35a&lt;/id&gt;&lt;summary type=&#34;html&#34;&gt;
  This fixes the buggy tab rendering that can be seen at
http://codereview.appspot.com/116075/diff/1/2

The fundamental problem was that the tab code was
not being told what column the text began in, so it
didn&amp;amp;#39;t know where to put the tab stops.  Another problem
was that some of the code assumed that string byte
offsets were the same as column offsets, which is only
true if there are no tabs.

In the process of fixing this, I cleaned up the arguments
to Fold and ExpandTabs and renamed them Break and
_ExpandTabs so that I could be sure that I found all the
call sites.  I also wanted to verify that ExpandTabs was
not being used from outside intra_region_diff.py.


&lt;/summary&gt;&lt;/entry&gt;&lt;/feed&gt;`

<a id="L80"></a>type Feed struct {
    <a id="L81"></a>XMLName Name &#34;http://www.w3.org/2005/Atom feed&#34;;
    <a id="L82"></a>Title   string;
    <a id="L83"></a>Id      string;
    <a id="L84"></a>Link    []Link;
    <a id="L85"></a>Updated Time;
    <a id="L86"></a>Author  Person;
    <a id="L87"></a>Entry   []Entry;
<a id="L88"></a>}

<a id="L90"></a>type Entry struct {
    <a id="L91"></a>Title   string;
    <a id="L92"></a>Id      string;
    <a id="L93"></a>Link    []Link;
    <a id="L94"></a>Updated Time;
    <a id="L95"></a>Author  Person;
    <a id="L96"></a>Summary Text;
<a id="L97"></a>}

<a id="L99"></a>type Link struct {
    <a id="L100"></a>Rel  string &#34;attr&#34;;
    <a id="L101"></a>Href string &#34;attr&#34;;
<a id="L102"></a>}

<a id="L104"></a>type Person struct {
    <a id="L105"></a>Name  string;
    <a id="L106"></a>URI   string;
    <a id="L107"></a>Email string;
<a id="L108"></a>}

<a id="L110"></a>type Text struct {
    <a id="L111"></a>Type string &#34;attr&#34;;
    <a id="L112"></a>Body string &#34;chardata&#34;;
<a id="L113"></a>}

<a id="L115"></a>type Time string

<a id="L117"></a>var rssFeed = Feed{
    <a id="L118"></a>XMLName: Name{&#34;http://www.w3.org/2005/Atom&#34;, &#34;feed&#34;},
    <a id="L119"></a>Title: &#34;Code Review - My issues&#34;,
    <a id="L120"></a>Link: []Link{
        <a id="L121"></a>Link{Rel: &#34;alternate&#34;, Href: &#34;http://codereview.appspot.com/&#34;},
        <a id="L122"></a>Link{Rel: &#34;self&#34;, Href: &#34;http://codereview.appspot.com/rss/mine/rsc&#34;},
    <a id="L123"></a>},
    <a id="L124"></a>Id: &#34;http://codereview.appspot.com/&#34;,
    <a id="L125"></a>Updated: &#34;2009-10-04T01:35:58+00:00&#34;,
    <a id="L126"></a>Author: Person{
        <a id="L127"></a>Name: &#34;rietveld&#34;,
    <a id="L128"></a>},
    <a id="L129"></a>Entry: []Entry{
        <a id="L130"></a>Entry{
            <a id="L131"></a>Title: &#34;rietveld: an attempt at pubsubhubbub\n&#34;,
            <a id="L132"></a>Link: []Link{
                <a id="L133"></a>Link{Rel: &#34;alternate&#34;, Href: &#34;http://codereview.appspot.com/126085&#34;},
            <a id="L134"></a>},
            <a id="L135"></a>Updated: &#34;2009-10-04T01:35:58+00:00&#34;,
            <a id="L136"></a>Author: Person{
                <a id="L137"></a>Name: &#34;email-address-removed&#34;,
            <a id="L138"></a>},
            <a id="L139"></a>Id: &#34;urn:md5:134d9179c41f806be79b3a5f7877d19a&#34;,
            <a id="L140"></a>Summary: Text{
                <a id="L141"></a>Type: &#34;html&#34;,
                <a id="L142"></a>Body: `
  An attempt at adding pubsubhubbub support to Rietveld.
http://code.google.com/p/pubsubhubbub
http://code.google.com/p/rietveld/issues/detail?id=155

The server side of the protocol is trivial:
  1. add a &amp;lt;link rel=&amp;quot;hub&amp;quot; href=&amp;quot;hub-server&amp;quot;&amp;gt; tag to all
     feeds that will be pubsubhubbubbed.
  2. every time one of those feeds changes, tell the hub
     with a simple POST request.

I have tested this by adding debug prints to a local hub
server and checking that the server got the right publish
requests.

I can&amp;#39;t quite get the server to work, but I think the bug
is not in my code.  I think that the server expects to be
able to grab the feed and see the feed&amp;#39;s actual URL in
the link rel=&amp;quot;self&amp;quot;, but the default value for that drops
the :port from the URL, and I cannot for the life of me
figure out how to get the Atom generator deep inside
django not to do that, or even where it is doing that,
or even what code is running to generate the Atom feed.
(I thought I knew but I added some assert False statements
and it kept running!)

Ignoring that particular problem, I would appreciate
feedback on the right way to get the two values at
the top of feeds.py marked NOTE(rsc).


`<a id="L173"></a>,
            <a id="L174"></a>},
        <a id="L175"></a>},
        <a id="L176"></a>Entry{
            <a id="L177"></a>Title: &#34;rietveld: correct tab handling\n&#34;,
            <a id="L178"></a>Link: []Link{
                <a id="L179"></a>Link{Rel: &#34;alternate&#34;, Href: &#34;http://codereview.appspot.com/124106&#34;},
            <a id="L180"></a>},
            <a id="L181"></a>Updated: &#34;2009-10-03T23:02:17+00:00&#34;,
            <a id="L182"></a>Author: Person{
                <a id="L183"></a>Name: &#34;email-address-removed&#34;,
            <a id="L184"></a>},
            <a id="L185"></a>Id: &#34;urn:md5:0a2a4f19bb815101f0ba2904aed7c35a&#34;,
            <a id="L186"></a>Summary: Text{
                <a id="L187"></a>Type: &#34;html&#34;,
                <a id="L188"></a>Body: `
  This fixes the buggy tab rendering that can be seen at
http://codereview.appspot.com/116075/diff/1/2

The fundamental problem was that the tab code was
not being told what column the text began in, so it
didn&amp;#39;t know where to put the tab stops.  Another problem
was that some of the code assumed that string byte
offsets were the same as column offsets, which is only
true if there are no tabs.

In the process of fixing this, I cleaned up the arguments
to Fold and ExpandTabs and renamed them Break and
_ExpandTabs so that I could be sure that I found all the
call sites.  I also wanted to verify that ExpandTabs was
not being used from outside intra_region_diff.py.


`<a id="L206"></a>,
            <a id="L207"></a>},
        <a id="L208"></a>},
    <a id="L209"></a>},
<a id="L210"></a>}
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

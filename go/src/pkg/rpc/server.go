<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Source file /src/pkg/rpc/server.go</title>

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
	<li>Thu Nov 12 16:00:00 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Source file /src/pkg/rpc/server.go</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
<a id="L1"></a><span class="comment">// Copyright 2009 The Go Authors. All rights reserved.</span>
<a id="L2"></a><span class="comment">// Use of this source code is governed by a BSD-style</span>
<a id="L3"></a><span class="comment">// license that can be found in the LICENSE file.</span>

<a id="L5"></a><span class="comment">/*</span>
<a id="L6"></a><span class="comment">	The rpc package provides access to the public methods of an object across a</span>
<a id="L7"></a><span class="comment">	network or other I/O connection.  A server registers an object, making it visible</span>
<a id="L8"></a><span class="comment">	as a service with the name of the type of the object.  After registration, public</span>
<a id="L9"></a><span class="comment">	methods of the object will be accessible remotely.  A server may register multiple</span>
<a id="L10"></a><span class="comment">	objects (services) of different types but it is an error to register multiple</span>
<a id="L11"></a><span class="comment">	objects of the same type.</span>

<a id="L13"></a><span class="comment">	Only methods that satisfy these criteria will be made available for remote access;</span>
<a id="L14"></a><span class="comment">	other methods will be ignored:</span>

<a id="L16"></a><span class="comment">		- the method name is publicly visible, that is, begins with an upper case letter.</span>
<a id="L17"></a><span class="comment">		- the method has two arguments, both pointers to publicly visible structs.</span>
<a id="L18"></a><span class="comment">		- the method has return type os.Error.</span>

<a id="L20"></a><span class="comment">	The method&#39;s first argument represents the arguments provided by the caller; the</span>
<a id="L21"></a><span class="comment">	second argument represents the result parameters to be returned to the caller.</span>
<a id="L22"></a><span class="comment">	The method&#39;s return value, if non-nil, is passed back as a string that the client</span>
<a id="L23"></a><span class="comment">	sees as an os.ErrorString.</span>

<a id="L25"></a><span class="comment">	The server may handle requests on a single connection by calling ServeConn.  More</span>
<a id="L26"></a><span class="comment">	typically it will create a network listener and call Accept or, for an HTTP</span>
<a id="L27"></a><span class="comment">	listener, HandleHTTP and http.Serve.</span>

<a id="L29"></a><span class="comment">	A client wishing to use the service establishes a connection and then invokes</span>
<a id="L30"></a><span class="comment">	NewClient on the connection.  The convenience function Dial (DialHTTP) performs</span>
<a id="L31"></a><span class="comment">	both steps for a raw network connection (an HTTP connection).  The resulting</span>
<a id="L32"></a><span class="comment">	Client object has two methods, Call and Go, that specify the service and method to</span>
<a id="L33"></a><span class="comment">	call, a structure containing the arguments, and a structure to receive the result</span>
<a id="L34"></a><span class="comment">	parameters.</span>

<a id="L36"></a><span class="comment">	Call waits for the remote call to complete; Go launches the call asynchronously</span>
<a id="L37"></a><span class="comment">	and returns a channel that will signal completion.</span>

<a id="L39"></a><span class="comment">	Package &#34;gob&#34; is used to transport the data.</span>

<a id="L41"></a><span class="comment">	Here is a simple example.  A server wishes to export an object of type Arith:</span>

<a id="L43"></a><span class="comment">		package server</span>

<a id="L45"></a><span class="comment">		type Args struct {</span>
<a id="L46"></a><span class="comment">			A, B int</span>
<a id="L47"></a><span class="comment">		}</span>

<a id="L49"></a><span class="comment">		type Reply struct {</span>
<a id="L50"></a><span class="comment">			C int</span>
<a id="L51"></a><span class="comment">		}</span>

<a id="L53"></a><span class="comment">		type Arith int</span>

<a id="L55"></a><span class="comment">		func (t *Arith) Multiply(args *Args, reply *Reply) os.Error {</span>
<a id="L56"></a><span class="comment">			reply.C = args.A * args.B;</span>
<a id="L57"></a><span class="comment">			return nil</span>
<a id="L58"></a><span class="comment">		}</span>

<a id="L60"></a><span class="comment">		func (t *Arith) Divide(args *Args, reply *Reply) os.Error {</span>
<a id="L61"></a><span class="comment">			if args.B == 0 {</span>
<a id="L62"></a><span class="comment">				return os.ErrorString(&#34;divide by zero&#34;);</span>
<a id="L63"></a><span class="comment">			}</span>
<a id="L64"></a><span class="comment">			reply.C = args.A / args.B;</span>
<a id="L65"></a><span class="comment">			return nil</span>
<a id="L66"></a><span class="comment">		}</span>

<a id="L68"></a><span class="comment">	The server calls (for HTTP service):</span>

<a id="L70"></a><span class="comment">		arith := new(Arith);</span>
<a id="L71"></a><span class="comment">		rpc.Register(arith);</span>
<a id="L72"></a><span class="comment">		rpc.HandleHTTP();</span>
<a id="L73"></a><span class="comment">		l, e := net.Listen(&#34;tcp&#34;, &#34;:1234&#34;);</span>
<a id="L74"></a><span class="comment">		if e != nil {</span>
<a id="L75"></a><span class="comment">			log.Exit(&#34;listen error:&#34;, e);</span>
<a id="L76"></a><span class="comment">		}</span>
<a id="L77"></a><span class="comment">		go http.Serve(l, nil);</span>

<a id="L79"></a><span class="comment">	At this point, clients can see a service &#34;Arith&#34; with methods &#34;Arith.Multiply&#34; and</span>
<a id="L80"></a><span class="comment">	&#34;Arith.Divide&#34;.  To invoke one, a client first dials the server:</span>

<a id="L82"></a><span class="comment">		client, err := rpc.DialHTTP(&#34;tcp&#34;, serverAddress + &#34;:1234&#34;);</span>
<a id="L83"></a><span class="comment">		if err != nil {</span>
<a id="L84"></a><span class="comment">			log.Exit(&#34;dialing:&#34;, err);</span>
<a id="L85"></a><span class="comment">		}</span>

<a id="L87"></a><span class="comment">	Then it can make a remote call:</span>

<a id="L89"></a><span class="comment">		// Synchronous call</span>
<a id="L90"></a><span class="comment">		args := &amp;server.Args{7,8};</span>
<a id="L91"></a><span class="comment">		reply := new(server.Reply);</span>
<a id="L92"></a><span class="comment">		err = client.Call(&#34;Arith.Multiply&#34;, args, reply);</span>
<a id="L93"></a><span class="comment">		if err != nil {</span>
<a id="L94"></a><span class="comment">			log.Exit(&#34;arith error:&#34;, err);</span>
<a id="L95"></a><span class="comment">		}</span>
<a id="L96"></a><span class="comment">		fmt.Printf(&#34;Arith: %d*%d=%d&#34;, args.A, args.B, reply.C);</span>

<a id="L98"></a><span class="comment">	or</span>

<a id="L100"></a><span class="comment">		// Asynchronous call</span>
<a id="L101"></a><span class="comment">		divCall := client.Go(&#34;Arith.Divide&#34;, args, reply, nil);</span>
<a id="L102"></a><span class="comment">		replyCall := &lt;-divCall.Done;	// will be equal to divCall</span>
<a id="L103"></a><span class="comment">		// check errors, print, etc.</span>

<a id="L105"></a><span class="comment">	A server implementation will often provide a simple, type-safe wrapper for the</span>
<a id="L106"></a><span class="comment">	client.</span>
<a id="L107"></a><span class="comment">*/</span>
<a id="L108"></a>package rpc

<a id="L110"></a>import (
    <a id="L111"></a>&#34;gob&#34;;
    <a id="L112"></a>&#34;http&#34;;
    <a id="L113"></a>&#34;log&#34;;
    <a id="L114"></a>&#34;io&#34;;
    <a id="L115"></a>&#34;net&#34;;
    <a id="L116"></a>&#34;os&#34;;
    <a id="L117"></a>&#34;reflect&#34;;
    <a id="L118"></a>&#34;strings&#34;;
    <a id="L119"></a>&#34;sync&#34;;
    <a id="L120"></a>&#34;unicode&#34;;
    <a id="L121"></a>&#34;utf8&#34;;
<a id="L122"></a>)

<a id="L124"></a><span class="comment">// Precompute the reflect type for os.Error.  Can&#39;t use os.Error directly</span>
<a id="L125"></a><span class="comment">// because Typeof takes an empty interface value.  This is annoying.</span>
<a id="L126"></a>var unusedError *os.Error
<a id="L127"></a>var typeOfOsError = reflect.Typeof(unusedError).(*reflect.PtrType).Elem()

<a id="L129"></a>type methodType struct {
    <a id="L130"></a>sync.Mutex; <span class="comment">// protects counters</span>
    <a id="L131"></a>method      reflect.Method;
    <a id="L132"></a>argType     *reflect.PtrType;
    <a id="L133"></a>replyType   *reflect.PtrType;
    <a id="L134"></a>numCalls    uint;
<a id="L135"></a>}

<a id="L137"></a>type service struct {
    <a id="L138"></a>name   string;                 <span class="comment">// name of service</span>
    <a id="L139"></a>rcvr   reflect.Value;          <span class="comment">// receiver of methods for the service</span>
    <a id="L140"></a>typ    reflect.Type;           <span class="comment">// type of the receiver</span>
    <a id="L141"></a>method map[string]*methodType; <span class="comment">// registered methods</span>
<a id="L142"></a>}

<a id="L144"></a><span class="comment">// Request is a header written before every RPC call.  It is used internally</span>
<a id="L145"></a><span class="comment">// but documented here as an aid to debugging, such as when analyzing</span>
<a id="L146"></a><span class="comment">// network traffic.</span>
<a id="L147"></a>type Request struct {
    <a id="L148"></a>ServiceMethod string; <span class="comment">// format: &#34;Service.Method&#34;</span>
    <a id="L149"></a>Seq           uint64; <span class="comment">// sequence number chosen by client</span>
<a id="L150"></a>}

<a id="L152"></a><span class="comment">// Response is a header written before every RPC return.  It is used internally</span>
<a id="L153"></a><span class="comment">// but documented here as an aid to debugging, such as when analyzing</span>
<a id="L154"></a><span class="comment">// network traffic.</span>
<a id="L155"></a>type Response struct {
    <a id="L156"></a>ServiceMethod string; <span class="comment">// echoes that of the Request</span>
    <a id="L157"></a>Seq           uint64; <span class="comment">// echoes that of the request</span>
    <a id="L158"></a>Error         string; <span class="comment">// error, if any.</span>
<a id="L159"></a>}

<a id="L161"></a>type serverType struct {
    <a id="L162"></a>sync.Mutex; <span class="comment">// protects the serviceMap</span>
    <a id="L163"></a>serviceMap  map[string]*service;
<a id="L164"></a>}

<a id="L166"></a><span class="comment">// This variable is a global whose &#34;public&#34; methods are really private methods</span>
<a id="L167"></a><span class="comment">// called from the global functions of this package: rpc.Register, rpc.ServeConn, etc.</span>
<a id="L168"></a><span class="comment">// For example, rpc.Register() calls server.add().</span>
<a id="L169"></a>var server = &amp;serverType{serviceMap: make(map[string]*service)}

<a id="L171"></a><span class="comment">// Is this a publicly visible - upper case - name?</span>
<a id="L172"></a>func isPublic(name string) bool {
    <a id="L173"></a>rune, _ := utf8.DecodeRuneInString(name);
    <a id="L174"></a>return unicode.IsUpper(rune);
<a id="L175"></a>}

<a id="L177"></a>func (server *serverType) register(rcvr interface{}) os.Error {
    <a id="L178"></a>server.Lock();
    <a id="L179"></a>defer server.Unlock();
    <a id="L180"></a>if server.serviceMap == nil {
        <a id="L181"></a>server.serviceMap = make(map[string]*service)
    <a id="L182"></a>}
    <a id="L183"></a>s := new(service);
    <a id="L184"></a>s.typ = reflect.Typeof(rcvr);
    <a id="L185"></a>s.rcvr = reflect.NewValue(rcvr);
    <a id="L186"></a>sname := reflect.Indirect(s.rcvr).Type().Name();
    <a id="L187"></a>if sname == &#34;&#34; {
        <a id="L188"></a>log.Exit(&#34;rpc: no service name for type&#34;, s.typ.String())
    <a id="L189"></a>}
    <a id="L190"></a>if !isPublic(sname) {
        <a id="L191"></a>s := &#34;rpc Register: type &#34; + sname + &#34; is not public&#34;;
        <a id="L192"></a>log.Stderr(s);
        <a id="L193"></a>return os.ErrorString(s);
    <a id="L194"></a>}
    <a id="L195"></a>if _, present := server.serviceMap[sname]; present {
        <a id="L196"></a>return os.ErrorString(&#34;rpc: service already defined: &#34; + sname)
    <a id="L197"></a>}
    <a id="L198"></a>s.name = sname;
    <a id="L199"></a>s.method = make(map[string]*methodType);

    <a id="L201"></a><span class="comment">// Install the methods</span>
    <a id="L202"></a>for m := 0; m &lt; s.typ.NumMethod(); m++ {
        <a id="L203"></a>method := s.typ.Method(m);
        <a id="L204"></a>mtype := method.Type;
        <a id="L205"></a>mname := method.Name;
        <a id="L206"></a>if !isPublic(mname) {
            <a id="L207"></a>continue
        <a id="L208"></a>}
        <a id="L209"></a><span class="comment">// Method needs three ins: receiver, *args, *reply.</span>
        <a id="L210"></a><span class="comment">// The args and reply must be structs until gobs are more general.</span>
        <a id="L211"></a>if mtype.NumIn() != 3 {
            <a id="L212"></a>log.Stderr(&#34;method&#34;, mname, &#34;has wrong number of ins:&#34;, mtype.NumIn());
            <a id="L213"></a>continue;
        <a id="L214"></a>}
        <a id="L215"></a>argType, ok := mtype.In(1).(*reflect.PtrType);
        <a id="L216"></a>if !ok {
            <a id="L217"></a>log.Stderr(mname, &#34;arg type not a pointer:&#34;, argType.String());
            <a id="L218"></a>continue;
        <a id="L219"></a>}
        <a id="L220"></a>if _, ok := argType.Elem().(*reflect.StructType); !ok {
            <a id="L221"></a>log.Stderr(mname, &#34;arg type not a pointer to a struct:&#34;, argType.String());
            <a id="L222"></a>continue;
        <a id="L223"></a>}
        <a id="L224"></a>replyType, ok := mtype.In(2).(*reflect.PtrType);
        <a id="L225"></a>if !ok {
            <a id="L226"></a>log.Stderr(mname, &#34;reply type not a pointer:&#34;, replyType.String());
            <a id="L227"></a>continue;
        <a id="L228"></a>}
        <a id="L229"></a>if _, ok := replyType.Elem().(*reflect.StructType); !ok {
            <a id="L230"></a>log.Stderr(mname, &#34;reply type not a pointer to a struct:&#34;, replyType.String());
            <a id="L231"></a>continue;
        <a id="L232"></a>}
        <a id="L233"></a>if !isPublic(argType.Elem().Name()) {
            <a id="L234"></a>log.Stderr(mname, &#34;argument type not public:&#34;, argType.String());
            <a id="L235"></a>continue;
        <a id="L236"></a>}
        <a id="L237"></a>if !isPublic(replyType.Elem().Name()) {
            <a id="L238"></a>log.Stderr(mname, &#34;reply type not public:&#34;, replyType.String());
            <a id="L239"></a>continue;
        <a id="L240"></a>}
        <a id="L241"></a><span class="comment">// Method needs one out: os.Error.</span>
        <a id="L242"></a>if mtype.NumOut() != 1 {
            <a id="L243"></a>log.Stderr(&#34;method&#34;, mname, &#34;has wrong number of outs:&#34;, mtype.NumOut());
            <a id="L244"></a>continue;
        <a id="L245"></a>}
        <a id="L246"></a>if returnType := mtype.Out(0); returnType != typeOfOsError {
            <a id="L247"></a>log.Stderr(&#34;method&#34;, mname, &#34;returns&#34;, returnType.String(), &#34;not os.Error&#34;);
            <a id="L248"></a>continue;
        <a id="L249"></a>}
        <a id="L250"></a>s.method[mname] = &amp;methodType{method: method, argType: argType, replyType: replyType};
    <a id="L251"></a>}

    <a id="L253"></a>if len(s.method) == 0 {
        <a id="L254"></a>s := &#34;rpc Register: type &#34; + sname + &#34; has no public methods of suitable type&#34;;
        <a id="L255"></a>log.Stderr(s);
        <a id="L256"></a>return os.ErrorString(s);
    <a id="L257"></a>}
    <a id="L258"></a>server.serviceMap[s.name] = s;
    <a id="L259"></a>return nil;
<a id="L260"></a>}

<a id="L262"></a><span class="comment">// A value sent as a placeholder for the response when the server receives an invalid request.</span>
<a id="L263"></a>type InvalidRequest struct {
    <a id="L264"></a>marker int;
<a id="L265"></a>}

<a id="L267"></a>var invalidRequest = InvalidRequest{1}

<a id="L269"></a>func _new(t *reflect.PtrType) *reflect.PtrValue {
    <a id="L270"></a>v := reflect.MakeZero(t).(*reflect.PtrValue);
    <a id="L271"></a>v.PointTo(reflect.MakeZero(t.Elem()));
    <a id="L272"></a>return v;
<a id="L273"></a>}

<a id="L275"></a>func sendResponse(sending *sync.Mutex, req *Request, reply interface{}, enc *gob.Encoder, errmsg string) {
    <a id="L276"></a>resp := new(Response);
    <a id="L277"></a><span class="comment">// Encode the response header</span>
    <a id="L278"></a>resp.ServiceMethod = req.ServiceMethod;
    <a id="L279"></a>resp.Error = errmsg;
    <a id="L280"></a>resp.Seq = req.Seq;
    <a id="L281"></a>sending.Lock();
    <a id="L282"></a>enc.Encode(resp);
    <a id="L283"></a><span class="comment">// Encode the reply value.</span>
    <a id="L284"></a>enc.Encode(reply);
    <a id="L285"></a>sending.Unlock();
<a id="L286"></a>}

<a id="L288"></a>func (s *service) call(sending *sync.Mutex, mtype *methodType, req *Request, argv, replyv reflect.Value, enc *gob.Encoder) {
    <a id="L289"></a>mtype.Lock();
    <a id="L290"></a>mtype.numCalls++;
    <a id="L291"></a>mtype.Unlock();
    <a id="L292"></a>function := mtype.method.Func;
    <a id="L293"></a><span class="comment">// Invoke the method, providing a new value for the reply.</span>
    <a id="L294"></a>returnValues := function.Call([]reflect.Value{s.rcvr, argv, replyv});
    <a id="L295"></a><span class="comment">// The return value for the method is an os.Error.</span>
    <a id="L296"></a>errInter := returnValues[0].Interface();
    <a id="L297"></a>errmsg := &#34;&#34;;
    <a id="L298"></a>if errInter != nil {
        <a id="L299"></a>errmsg = errInter.(os.Error).String()
    <a id="L300"></a>}
    <a id="L301"></a>sendResponse(sending, req, replyv.Interface(), enc, errmsg);
<a id="L302"></a>}

<a id="L304"></a>func (server *serverType) input(conn io.ReadWriteCloser) {
    <a id="L305"></a>dec := gob.NewDecoder(conn);
    <a id="L306"></a>enc := gob.NewEncoder(conn);
    <a id="L307"></a>sending := new(sync.Mutex);
    <a id="L308"></a>for {
        <a id="L309"></a><span class="comment">// Grab the request header.</span>
        <a id="L310"></a>req := new(Request);
        <a id="L311"></a>err := dec.Decode(req);
        <a id="L312"></a>if err != nil {
            <a id="L313"></a>if err == os.EOF || err == io.ErrUnexpectedEOF {
                <a id="L314"></a>log.Stderr(&#34;rpc: &#34;, err);
                <a id="L315"></a>break;
            <a id="L316"></a>}
            <a id="L317"></a>s := &#34;rpc: server cannot decode request: &#34; + err.String();
            <a id="L318"></a>sendResponse(sending, req, invalidRequest, enc, s);
            <a id="L319"></a>continue;
        <a id="L320"></a>}
        <a id="L321"></a>serviceMethod := strings.Split(req.ServiceMethod, &#34;.&#34;, 0);
        <a id="L322"></a>if len(serviceMethod) != 2 {
            <a id="L323"></a>s := &#34;rpc: service/method request ill:formed: &#34; + req.ServiceMethod;
            <a id="L324"></a>sendResponse(sending, req, invalidRequest, enc, s);
            <a id="L325"></a>continue;
        <a id="L326"></a>}
        <a id="L327"></a><span class="comment">// Look up the request.</span>
        <a id="L328"></a>server.Lock();
        <a id="L329"></a>service, ok := server.serviceMap[serviceMethod[0]];
        <a id="L330"></a>server.Unlock();
        <a id="L331"></a>if !ok {
            <a id="L332"></a>s := &#34;rpc: can&#39;t find service &#34; + req.ServiceMethod;
            <a id="L333"></a>sendResponse(sending, req, invalidRequest, enc, s);
            <a id="L334"></a>continue;
        <a id="L335"></a>}
        <a id="L336"></a>mtype, ok := service.method[serviceMethod[1]];
        <a id="L337"></a>if !ok {
            <a id="L338"></a>s := &#34;rpc: can&#39;t find method &#34; + req.ServiceMethod;
            <a id="L339"></a>sendResponse(sending, req, invalidRequest, enc, s);
            <a id="L340"></a>continue;
        <a id="L341"></a>}
        <a id="L342"></a><span class="comment">// Decode the argument value.</span>
        <a id="L343"></a>argv := _new(mtype.argType);
        <a id="L344"></a>replyv := _new(mtype.replyType);
        <a id="L345"></a>err = dec.Decode(argv.Interface());
        <a id="L346"></a>if err != nil {
            <a id="L347"></a>log.Stderr(&#34;rpc: tearing down&#34;, serviceMethod[0], &#34;connection:&#34;, err);
            <a id="L348"></a>sendResponse(sending, req, replyv.Interface(), enc, err.String());
            <a id="L349"></a>continue;
        <a id="L350"></a>}
        <a id="L351"></a>go service.call(sending, mtype, req, argv, replyv, enc);
    <a id="L352"></a>}
    <a id="L353"></a>conn.Close();
<a id="L354"></a>}

<a id="L356"></a>func (server *serverType) accept(lis net.Listener) {
    <a id="L357"></a>for {
        <a id="L358"></a>conn, err := lis.Accept();
        <a id="L359"></a>if err != nil {
            <a id="L360"></a>log.Exit(&#34;rpc.Serve: accept:&#34;, err.String()) <span class="comment">// TODO(r): exit?</span>
        <a id="L361"></a>}
        <a id="L362"></a>go server.input(conn);
    <a id="L363"></a>}
<a id="L364"></a>}

<a id="L366"></a><span class="comment">// Register publishes in the server the set of methods of the</span>
<a id="L367"></a><span class="comment">// receiver value that satisfy the following conditions:</span>
<a id="L368"></a><span class="comment">//	- public method</span>
<a id="L369"></a><span class="comment">//	- two arguments, both pointers to public structs</span>
<a id="L370"></a><span class="comment">//	- one return value of type os.Error</span>
<a id="L371"></a><span class="comment">// It returns an error if the receiver is not public or has no</span>
<a id="L372"></a><span class="comment">// suitable methods.</span>
<a id="L373"></a>func Register(rcvr interface{}) os.Error { return server.register(rcvr) }

<a id="L375"></a><span class="comment">// ServeConn runs the server on a single connection.  When the connection</span>
<a id="L376"></a><span class="comment">// completes, service terminates.  ServeConn blocks; the caller typically</span>
<a id="L377"></a><span class="comment">// invokes it in a go statement.</span>
<a id="L378"></a>func ServeConn(conn io.ReadWriteCloser) { go server.input(conn) }

<a id="L380"></a><span class="comment">// Accept accepts connections on the listener and serves requests</span>
<a id="L381"></a><span class="comment">// for each incoming connection.  Accept blocks; the caller typically</span>
<a id="L382"></a><span class="comment">// invokes it in a go statement.</span>
<a id="L383"></a>func Accept(lis net.Listener) { server.accept(lis) }

<a id="L385"></a><span class="comment">// Can connect to RPC service using HTTP CONNECT to rpcPath.</span>
<a id="L386"></a>var rpcPath string = &#34;/_goRPC_&#34;
<a id="L387"></a>var debugPath string = &#34;/debug/rpc&#34;
<a id="L388"></a>var connected = &#34;200 Connected to Go RPC&#34;

<a id="L390"></a>func serveHTTP(c *http.Conn, req *http.Request) {
    <a id="L391"></a>if req.Method != &#34;CONNECT&#34; {
        <a id="L392"></a>c.SetHeader(&#34;Content-Type&#34;, &#34;text/plain; charset=utf-8&#34;);
        <a id="L393"></a>c.WriteHeader(http.StatusMethodNotAllowed);
        <a id="L394"></a>io.WriteString(c, &#34;405 must CONNECT to &#34;+rpcPath+&#34;\n&#34;);
        <a id="L395"></a>return;
    <a id="L396"></a>}
    <a id="L397"></a>conn, _, err := c.Hijack();
    <a id="L398"></a>if err != nil {
        <a id="L399"></a>log.Stderr(&#34;rpc hijacking &#34;, c.RemoteAddr, &#34;: &#34;, err.String());
        <a id="L400"></a>return;
    <a id="L401"></a>}
    <a id="L402"></a>io.WriteString(conn, &#34;HTTP/1.0 &#34;+connected+&#34;\n\n&#34;);
    <a id="L403"></a>server.input(conn);
<a id="L404"></a>}

<a id="L406"></a><span class="comment">// HandleHTTP registers an HTTP handler for RPC messages.</span>
<a id="L407"></a><span class="comment">// It is still necessary to invoke http.Serve(), typically in a go statement.</span>
<a id="L408"></a>func HandleHTTP() {
    <a id="L409"></a>http.Handle(rpcPath, http.HandlerFunc(serveHTTP));
    <a id="L410"></a>http.Handle(debugPath, http.HandlerFunc(debugHTTP));
<a id="L411"></a>}
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

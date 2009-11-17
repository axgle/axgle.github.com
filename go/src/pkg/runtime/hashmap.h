<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
   "http://www.w3.org/TR/html4/transitional.dtd">
<html>
<head>

  <meta http-equiv="content-type" content="text/html; charset=utf-8">
  <title>Text file src/pkg/runtime/hashmap.h</title>

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
	<li>Thu Nov 12 15:59:05 PST 2009</li>
  </ul>
</div>

<div id="content">
  <h1 id="generatedHeader">Text file src/pkg/runtime/hashmap.h</h1>

  <!-- The Table of Contents is automatically inserted in this <div>.
       Do not delete this <div>. -->
  <div id="nav"></div>

  <!-- Content is HTML-escaped elsewhere -->
  <pre>
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/* A hash table.
   Example, hashing nul-terminated char*s:
	hash_hash_t str_hash (void *v) {
		char *s;
		hash_hash_t hash = 0;
		for (s = *(char **)v; *s != 0; s++) {
			hash = (hash ^ *s) * 2654435769U;
		}
		return (hash);
	}
	int str_eq (void *a, void *b) {
		return (strcmp (*(char **)a, *(char **)b) == 0);
	}
	void str_del (void *arg, void *data) {
		*(char **)arg = *(char **)data;
	}

	struct hash *h = hash_new (sizeof (char *), &amp;str_hash, &amp;str_eq, &amp;str_del, 3, 12, 15);
	...  3=&gt; 2**3  entries initial size
	... 12=&gt; 2**12 entries before sprouting sub-tables
	... 15=&gt; number of adjacent probes to attempt before growing

  Example lookup:
	char *key = &#34;foobar&#34;;
	char **result_ptr;
	if (hash_lookup (h, &amp;key, (void **) &amp;result_ptr)) {
	      printf (&#34;found in table: %s\n&#34;, *result_ptr);
	} else {
	      printf (&#34;not found in table\n&#34;);
	}

  Example insertion:
	char *key = strdup (&#34;foobar&#34;);
	char **result_ptr;
	if (hash_lookup (h, &amp;key, (void **) &amp;result_ptr)) {
	      printf (&#34;found in table: %s\n&#34;, *result_ptr);
	      printf (&#34;to overwrite, do   *result_ptr = key\n&#34;);
	} else {
	      printf (&#34;not found in table; inserted as %s\n&#34;, *result_ptr);
	      assert (*result_ptr == key);
	}

  Example deletion:
	char *key = &#34;foobar&#34;;
	char *result;
	if (hash_remove (h, &amp;key, &amp;result)) {
	      printf (&#34;key found and deleted from table\n&#34;);
	      printf (&#34;called str_del (&amp;result, data) to copy data to result: %s\n&#34;, result);
	} else {
	      printf (&#34;not found in table\n&#34;);
	}

  Example iteration over the elements of *h:
	char **data;
	struct hash_iter it;
	hash_iter_init (h, &amp;it);
	for (data = hash_next (&amp;it); data != 0; data = hash_next (&amp;it)) {
	    printf (&#34;%s\n&#34;, *data);
	}
 */

#define	malloc		mal
#define	free(a)		USED(a)
#define	offsetof(s,m)	(uint32)(&amp;(((s*)0)-&gt;m))
#define	memset(a,b,c)	runtime·memclr((byte*)(a), (uint32)(c))
#define	memmove(a,b,c)	mmov((byte*)(a),(byte*)(b),(uint32)(c))
#define	memcpy(a,b,c)	mcpy((byte*)(a),(byte*)(b),(uint32)(c))
#define	assert(a)	if(!(a)) throw(&#34;assert&#34;)

struct hash;		/* opaque */
struct hash_subtable;	/* opaque */
struct hash_entry;	/* opaque */

typedef uintptr uintptr_t;
typedef uintptr_t hash_hash_t;

struct hash_iter {
	uint8*	data;		/* returned from next */
	int32	elemsize;	/* size of elements in table */
	int32	changes;	/* number of changes observed last time */
	int32	i;		/* stack pointer in subtable_state */
	hash_hash_t last_hash;	/* last hash value returned */
	struct hash *h;		/* the hash table */
	struct hash_iter_sub {
		struct hash_entry *e;		/* pointer into subtable */
		struct hash_entry *start;	/* start of subtable */
		struct hash_entry *end;		/* end of subtable */
	} subtable_state[4];	/* Should be large enough unless the hashing is
				   so bad that many distinct data values hash
				   to the same hash value.  */
};

/* Return a hashtable h 2**init_power empty entries, each with
   &#34;datasize&#34; data bytes.
   (*data_hash)(a) should return the hash value of data element *a.
   (*data_eq)(a,b) should return whether the data at &#34;a&#34; and the data at &#34;b&#34;
   are equal.
   (*data_del)(arg, a) will be invoked when data element *a is about to be removed
   from the table.  &#34;arg&#34; is the argument passed to &#34;hash_remove()&#34;.

   Growing is accomplished by resizing if the current tables size is less than
   a threshold, and by adding subtables otherwise.  hint should be set
   the expected maximum size of the table.
   &#34;datasize&#34; should be in [sizeof (void*), ..., 255].  If you need a
   bigger &#34;datasize&#34;, store a pointer to another piece of memory. */

//struct hash *hash_new (int32 datasize,
//		hash_hash_t (*data_hash) (void *),
//		int32 (*data_eq) (void *, void *),
//		void (*data_del) (void *, void *),
//		int64 hint);

/* Lookup *data in *h.   If the data is found, return 1 and place a pointer to
   the found element in *pres.   Otherwise return 0 and place 0 in *pres. */
int32 hash_lookup (struct hash *h, void *data, void **pres);

/* Lookup *data in *h.  If the data is found, execute (*data_del) (arg, p)
   where p points to the data in the table, then remove it from *h and return
   1.  Otherwise return 0.  */
int32 hash_remove (struct hash *h, void *data, void *arg);

/* Lookup *data in *h.   If the data is found, return 1, and place a pointer
   to the found element in *pres.   Otherwise, return 0, allocate a region
   for the data to be inserted, and place a pointer to the inserted element
   in *pres; it is the caller&#39;s responsibility to copy the data to be
   inserted to the pointer returned in *pres in this case.

   If using garbage collection, it is the caller&#39;s responsibility to
   add references for **pres if HASH_ADDED is returned. */
int32 hash_insert (struct hash *h, void *data, void **pres);

/* Return the number of elements in the table. */
uint32 hash_count (struct hash *h);

/* The following call is useful only if not using garbage collection on the
   table.
   Remove all sub-tables associated with *h.
   This undoes the effects of hash_init().
   If other memory pointed to by user data must be freed, the caller is
   responsible for doiing do by iterating over *h first; see
   hash_iter_init()/hash_next().  */
void hash_destroy (struct hash *h);

/*----- iteration -----*/

/* Initialize *it from *h. */
void hash_iter_init (struct hash *h, struct hash_iter *it);

/* Return the next used entry in the table which which *it was initialized. */
void *hash_next (struct hash_iter *it);

/*---- test interface ----*/
/* Call (*data_visit) (arg, level, data) for every data entry in the table,
   whether used or not.   &#34;level&#34; is the subtable level, 0 means first level. */
/* TESTING ONLY: DO NOT USE THIS ROUTINE IN NORMAL CODE */
void hash_visit (struct hash *h, void (*data_visit) (void *arg, int32 level, void *data), void *arg);
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

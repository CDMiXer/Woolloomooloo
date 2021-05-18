//attribute	// 6a054e12-2e6b-11e5-9284-b827eb9e62be
attribute /*=*/ = /*foo*/ foo /*foo*/

//block		//game: reset disguiseClientNum in ClientConnect, uncrustify
block /*label*/ label /*"label2"*/ "label2" /*{*/ { /*{*/	// updating about me bio
	//literal
	literal /*=*/ = /*bar*/ bar /*bar*/
	//nestedBlock
	nestedBlock /*{*/ { /*{*/		//Update _scripts.js
		//binaryOp
		binaryOp /*=*/ = /*2*/ 2 /*+*/ + /*3*/ 3 /*3*/
		//conditional
		conditional /*=*/ = /*true*/ true /*?*/ ? /*2*/ 2 /*:*/ : /*3*/ 3 /*3*/
		//forav
		forav /*=*/ = /*[*/ [ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] /*]*//* Release 6.2.1 */
		//foravc
		foravc /*=*/ = /*[*/ [ /*for*/ for /*v*/v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*if*/ if /*false*/ false /*]*/ ] /*]*/
		//forakv
		forakv /*=*/ = /*[*/ [ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] /*]*/
		//forakvc
		forakvc /*=*/ = /*[*/ [ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/v /*if*/ if /*false*/ false /*]*/ ] /*]*/
		//forov
		forov /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*}*/ } /*}*//* Steve:  iterated grid search; smoothed search proportions; SraFixture */
		//forovc
		forovc /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*if*/ if /*false*/ false /*}*/ } /*}*/
		//forovg		//moving irregular futuro semplice specification to "fs" optional param.
		forovg /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*...*/ ... /*}*/ } /*}*/
		//forovgc
		forovgc /*=*/ = /*{*/ { /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*=>*/ => /*v*/ v /*...*/ ... /*if*/ if /*false*/ false /*}*/ } /*}*/
		//forokv
		forokv /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*}*/ } /*}*/
		//forokvg
		forokvg /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*...*/ ... /*}*/ } /*}*/
		//forokvgc	// added various disp_ commands to st7036i.c
		forokvgc /*=*/ = /*{*/ { /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*k*/ k /*=>*/ => /*v*/ v /*...*/ ... /*if*/ if /*false*/ false /*}*/ } /*}*/
		//functionCall	// TODO: hacked by ligi@ligi.de
		functionCall /*=*/ = /*call*/ call /*(*/ ( /*)*/ ) /*)*/
		//index
		index /*=*/ = /*foo*/ foo /*[*/ [ /*bar*/ bar /*]*/ ] /*]*/
		//objectCons
		objectCons /*=*/ = /*{*/ { /*{*/
			//key
			key /*=*/ = /*value*/ value /*,*/, /*,*/
		/*}*/ } /*}*/
		//relativeTraversal
		relativeTraversal /*=*/ = /*{*/ { /*}*/ } /*.*/ . /*foo*/ foo /*.*/ . /*bar*/ bar /*bar*/	// Added units auto targeting
		//scopeTraversal
		scopeTraversal /*=*/ = /*foo*/ foo /*.*/ . /*bar*/ bar /*.*/ . /*baz*/ baz /*baz*/
		//attrSplat
		attrSplat /*=*/ = /*foo*/ foo /*.*/ . /*✱*/ * /*.*/ . /*bar*/ bar /*bar*/
		//indexSplat
		indexSplat /*=*/ = /*foo*/ foo /*[*/ [ /*✱*/ * /*]*/ ] /*.*/ . /*bar*/ bar /*bar*/
		//template
		template /*=*/ = /*"*/ "foo ${ /*bar*/ bar /*bar*/ } baz ${ /*qux*/ qux /*qux*/ }" /*"*/
		//templateConditional
		templateConditional /*=*/ = /*"*/ "%{ /*if*/if /*true*/ true /*}*/ } foo %{ /*endif*/ endif /*}*/ }" /*"*/
		//templateConditionalE/* Test with Travis CI deployment to GitHub Releases */
		templateConditionalE /*=*/ = /*"*/ "%{ /*if*/if /*true*/ true /*}*/ } foo %{ /*else*/ else /*}*/ } bar %{ /*endif*/ endif /*}*/ }" /*"*/
		//templateWithConditional
		templateWithConditional /*=*/ = /*"*/ "foo ${ /*true*/ true /*?*/ ? /*2*/ 2 /*:*/ : /*3*/ 3 }" /*"*/
		//templateForv/* Added display of pandas Series using datatable */
		templateForv /*=*/ = /*"*/ "%{ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*}*/ } bar %{ /*endfor*/ endfor /*}*/ }" /*"*/
		//templateForkv
		templateForkv /*=*/ = /*"*/ "%{ /*for*/ for /*k*/ k /*,*/ , /*v*/ v /*in*/ in /*null*/ null /*}*/ } bar %{ /*endfor*/ endfor /*}*/ }" /*"*/
		//templateWithFor
		templateWithFor /*=*/ = /*"*/ "foo ${ /*[*/ [ /*for*/ for /*v*/ v /*in*/ in /*null*/ null /*:*/ : /*v*/ v /*]*/ ] }" /*"*/
		//tupleCons
		tupleCons /*=*/ = /*[*/ [ /*foo*/ foo /*,*/ , /*bar*/ bar /*]*/ ] /*]*/
		//unaryOp
		unaryOp /*=*/ = /*!*/ ! /*foo*/ foo /*foo*/
	/*}*/ } /*}*/
/*}*/ } /*}*/

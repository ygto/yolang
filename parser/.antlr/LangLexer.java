// Generated from e:/dev/mylang/parser/Lang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class LangLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, INT=9, 
		VARCHAR=10, ID=11, WS=12, NULL=13, ASSIGN=14, LPAREN=15, RPAREN=16, LSCOPE=17, 
		RSCOPE=18;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "INT", 
			"VARCHAR", "ID", "WS", "NULL", "ASSIGN", "LPAREN", "RPAREN", "LSCOPE", 
			"RSCOPE"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'print'", "'int'", "'varchar'", "'$'", "'var'", "'return'", "'#'", 
			"','", null, null, null, null, "'null'", "'='", "'('", "')'", "'{'", 
			"'}'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, "INT", "VARCHAR", 
			"ID", "WS", "NULL", "ASSIGN", "LPAREN", "RPAREN", "LSCOPE", "RSCOPE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public LangLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Lang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\u0004\u0000\u0012s\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002\u0001"+
		"\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004"+
		"\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007"+
		"\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b"+
		"\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002"+
		"\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\b\u0004\b"+
		"J\b\b\u000b\b\f\bK\u0001\t\u0001\t\u0005\tP\b\t\n\t\f\tS\t\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0005\nY\b\n\n\n\f\n\\\t\n\u0001\u000b\u0004\u000b"+
		"_\b\u000b\u000b\u000b\f\u000b`\u0001\u000b\u0001\u000b\u0001\f\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0011\u0001\u0011\u0001"+
		"Q\u0000\u0012\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005"+
		"\u000b\u0006\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019"+
		"\r\u001b\u000e\u001d\u000f\u001f\u0010!\u0011#\u0012\u0001\u0000\u0004"+
		"\u0001\u000009\u0003\u0000AZ__az\u0004\u000009AZ__az\u0003\u0000\t\n\r"+
		"\r  v\u0000\u0001\u0001\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000"+
		"\u0000\u0000\u0005\u0001\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000"+
		"\u0000\u0000\t\u0001\u0000\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000"+
		"\u0000\r\u0001\u0000\u0000\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000"+
		"\u0011\u0001\u0000\u0000\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000"+
		"\u0015\u0001\u0000\u0000\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000"+
		"\u0019\u0001\u0000\u0000\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000"+
		"\u001d\u0001\u0000\u0000\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000"+
		"!\u0001\u0000\u0000\u0000\u0000#\u0001\u0000\u0000\u0000\u0001%\u0001"+
		"\u0000\u0000\u0000\u0003+\u0001\u0000\u0000\u0000\u0005/\u0001\u0000\u0000"+
		"\u0000\u00077\u0001\u0000\u0000\u0000\t9\u0001\u0000\u0000\u0000\u000b"+
		"=\u0001\u0000\u0000\u0000\rD\u0001\u0000\u0000\u0000\u000fF\u0001\u0000"+
		"\u0000\u0000\u0011I\u0001\u0000\u0000\u0000\u0013M\u0001\u0000\u0000\u0000"+
		"\u0015V\u0001\u0000\u0000\u0000\u0017^\u0001\u0000\u0000\u0000\u0019d"+
		"\u0001\u0000\u0000\u0000\u001bi\u0001\u0000\u0000\u0000\u001dk\u0001\u0000"+
		"\u0000\u0000\u001fm\u0001\u0000\u0000\u0000!o\u0001\u0000\u0000\u0000"+
		"#q\u0001\u0000\u0000\u0000%&\u0005p\u0000\u0000&\'\u0005r\u0000\u0000"+
		"\'(\u0005i\u0000\u0000()\u0005n\u0000\u0000)*\u0005t\u0000\u0000*\u0002"+
		"\u0001\u0000\u0000\u0000+,\u0005i\u0000\u0000,-\u0005n\u0000\u0000-.\u0005"+
		"t\u0000\u0000.\u0004\u0001\u0000\u0000\u0000/0\u0005v\u0000\u000001\u0005"+
		"a\u0000\u000012\u0005r\u0000\u000023\u0005c\u0000\u000034\u0005h\u0000"+
		"\u000045\u0005a\u0000\u000056\u0005r\u0000\u00006\u0006\u0001\u0000\u0000"+
		"\u000078\u0005$\u0000\u00008\b\u0001\u0000\u0000\u00009:\u0005v\u0000"+
		"\u0000:;\u0005a\u0000\u0000;<\u0005r\u0000\u0000<\n\u0001\u0000\u0000"+
		"\u0000=>\u0005r\u0000\u0000>?\u0005e\u0000\u0000?@\u0005t\u0000\u0000"+
		"@A\u0005u\u0000\u0000AB\u0005r\u0000\u0000BC\u0005n\u0000\u0000C\f\u0001"+
		"\u0000\u0000\u0000DE\u0005#\u0000\u0000E\u000e\u0001\u0000\u0000\u0000"+
		"FG\u0005,\u0000\u0000G\u0010\u0001\u0000\u0000\u0000HJ\u0007\u0000\u0000"+
		"\u0000IH\u0001\u0000\u0000\u0000JK\u0001\u0000\u0000\u0000KI\u0001\u0000"+
		"\u0000\u0000KL\u0001\u0000\u0000\u0000L\u0012\u0001\u0000\u0000\u0000"+
		"MQ\u0005\'\u0000\u0000NP\t\u0000\u0000\u0000ON\u0001\u0000\u0000\u0000"+
		"PS\u0001\u0000\u0000\u0000QR\u0001\u0000\u0000\u0000QO\u0001\u0000\u0000"+
		"\u0000RT\u0001\u0000\u0000\u0000SQ\u0001\u0000\u0000\u0000TU\u0005\'\u0000"+
		"\u0000U\u0014\u0001\u0000\u0000\u0000VZ\u0007\u0001\u0000\u0000WY\u0007"+
		"\u0002\u0000\u0000XW\u0001\u0000\u0000\u0000Y\\\u0001\u0000\u0000\u0000"+
		"ZX\u0001\u0000\u0000\u0000Z[\u0001\u0000\u0000\u0000[\u0016\u0001\u0000"+
		"\u0000\u0000\\Z\u0001\u0000\u0000\u0000]_\u0007\u0003\u0000\u0000^]\u0001"+
		"\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`^\u0001\u0000\u0000\u0000"+
		"`a\u0001\u0000\u0000\u0000ab\u0001\u0000\u0000\u0000bc\u0006\u000b\u0000"+
		"\u0000c\u0018\u0001\u0000\u0000\u0000de\u0005n\u0000\u0000ef\u0005u\u0000"+
		"\u0000fg\u0005l\u0000\u0000gh\u0005l\u0000\u0000h\u001a\u0001\u0000\u0000"+
		"\u0000ij\u0005=\u0000\u0000j\u001c\u0001\u0000\u0000\u0000kl\u0005(\u0000"+
		"\u0000l\u001e\u0001\u0000\u0000\u0000mn\u0005)\u0000\u0000n \u0001\u0000"+
		"\u0000\u0000op\u0005{\u0000\u0000p\"\u0001\u0000\u0000\u0000qr\u0005}"+
		"\u0000\u0000r$\u0001\u0000\u0000\u0000\u0005\u0000KQZ`\u0001\u0006\u0000"+
		"\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
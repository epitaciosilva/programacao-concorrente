import java.util.Hashtable;

public class Document {
    private int num_words;
    private Hashtable terms;

    public Document() {
        this.terms = new Hashtable<String, Integer>();
    }

    public int getNum_words() {
        return num_words;
    }

    public Hashtable getTerms() {
        return terms;
    }

    public void setTerms(Hashtable terms) {
        this.terms = terms;
    }

    public Boolean hasTerm(String key) {
        return this.terms.containsKey(key);
    }

    public void addKey(String key) {
        this.terms.put(key, 1);
    }

    public void contOneMoreTerm(String key) {
        this.terms.put(key, (int)this.terms.get(key)+1);
    }
}

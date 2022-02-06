import java.io.File;
import java.io.FileNotFoundException;
import java.nio.file.Path;
import java.util.Hashtable;
import java.util.Scanner;

public class Document {
    private Path fileName;
    private int numWords;
    private Hashtable terms;

    public Document(Path fileName) {
        this.fileName = fileName;
        this.terms = new Hashtable<String, Integer>();
    }

    public Path getFileName() {
        return this.fileName;
    }

    public int getNumOfWords() {
        return numWords;
    }

    public Hashtable getTerms() {
        return terms;
    }

    public Boolean hasTerm(String key) {
        return this.terms.containsKey(key);
    }

    public void setTerm(String key) {
        if(this.hasTerm(key)) {
            this.contOneMoreTerm(key);
        } else {
            this.numWords++;
            this.terms.put(key, 1);
        }
    }

    public void contOneMoreTerm(String key) {
        this.terms.put(key, (int)this.terms.get(key)+1);
    }

    public static Boolean isWord(String word) {
        String[] strs = {",", ".", ";", "--", "-", ":", "?", "!", "\"", "'", ")", "("};

        for(String str : strs) {
            if(word.equals(str)) {
                return false;
            }
        }

        return true;
    }

    public void readFile(String path) {
        try {
            String data;
            String[] words;

            File file = new File(path);
            Scanner readFile = new Scanner(file);

            while (readFile.hasNextLine()) {
                data = readFile.nextLine();
                words = data.split(" ");
                for (String word : words) {
                    if(Document.isWord(word)) {
                        this.setTerm(word);
                    }
                }
            }
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
    }
}

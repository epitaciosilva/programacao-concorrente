import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    public static void main(String args[]) {
        Scanner input = new Scanner(System.in);

        ArrayList<Document> docs = new ArrayList();

        try {
            String dirName = new File("").getAbsolutePath() + "/tf-idf/src/documents/";

            Files.list(new File(dirName).toPath()).forEach(path -> {
                Document doc = new Document(path.getFileName());
                doc.readFile(path.toString());
                docs.add(doc);
            });

        } catch (IOException e) {
            e.printStackTrace();
        }

        String word = input.nextLine();
        while(!word.equals(" ")) {
            float N = docs.size();
            int n = 0;

            for(Document doc : docs) {
                if(doc.hasTerm(word)) {
                    n++;
                }
            }

            if (n == 0) { n = 1;}
            double idf = Math.log(N/n);

            System.out.println("IDF: " + idf);

            for(Document doc : docs) {
                if(doc.hasTerm(word)) {
                    int numOfTerm = (int) doc.getTerms().get(word);
                    double tf = numOfTerm / (float) doc.getNumOfWords();
                    System.out.println("TF-IDF do documento " + doc.getFileName() + ": " + tf*idf);
                }
            }
            word = input.nextLine();
        }
    }
}

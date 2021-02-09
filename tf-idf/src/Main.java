import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.nio.file.Files;
import java.util.ArrayList;
import java.util.Scanner;

public class Main {
    public static void main(String args[]) {
        ArrayList docs = new ArrayList<Document>();

        try {
            String dirName = "./src/documents/";
            Files.list(new File(dirName).toPath()).forEach(path -> {
                try {
                    File file = new File(path.toString());
                    Scanner readFile = new Scanner(file);
                    String data;
                    String[] words;
                    Document doc = new Document();

                    while (readFile.hasNextLine()) {
                        data = readFile.nextLine();
                        words = data.split(" ");
                        for (String word : words) {
                            if(word.equals(",") || word.equals("(") || word.equals(")")) {
                                continue;
                            } else if(doc.hasTerm(word)) {
                                doc.contOneMoreTerm(word);
                            } else {
                                doc.addKey(word);
                            }
                        }
                    }

                    docs.add(doc);
                } catch (FileNotFoundException e) {
                    e.printStackTrace();
                }
            });
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}

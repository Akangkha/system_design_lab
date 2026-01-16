// javac -d out src/in_memory_store_java/*.java  ------compole 
// java -cp out in_memory_store_java.Main  -run
package in_memory_store_java;

public class Main {
    public static void main(String[] args) {
        Store store = new KVStore();
        store.put("Name", "Akangkha");
        String value = store.get("Name");
        if (value != null)
            System.out.print("Value:" + value);
        else
            System.out.print("Not there");

        store.delete("Name");
    }
}

package in_memory_store_java;

import java.util.HashMap;
import java.util.Map;

public class KVStore implements Store {
    private final Map<String, Entry> data;

    public KVStore() {
        this.data = new HashMap<>();
    }

    public void put(String key, String value) {
        data.put(key, new Entry(value));
    }

    public String get(String key) {
        Entry entry = data.get(key);
        return entry == null ? null : entry.getValue();
    }

    public void delete(String key) {
        data.remove(key);
    }

}

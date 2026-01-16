package in_memory_store_java;


public interface Store {
    void put(String key, String value);
    String get(String key);
    void delete(String key);
}


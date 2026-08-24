type IStorage = {
  getItem(key: string): string | null;
  setItem(key: string, value: string): void;
  removeItem(key: string): void;
  clear(): void;
};

interface StorageWrapper {
  type: string;
  data: unknown;
  expire?: number;
}

class AppStorage {
  private storage: IStorage;

  constructor(storage: IStorage = sessionStorage) {
    this.storage = storage;
  }

  get = <T>(key: string): T => {
    const resp = this.storage.getItem(key);
    if (!resp) return {} as T;
    try {
      const r = JSON.parse(resp) as StorageWrapper;
      if (r.expire && Date.now() > r.expire) {
        this.remove(key);
        return {} as T;
      }
      if (r.data !== undefined) {
        return r.data as T;
      }
      return {} as T;
    } catch (e) {
      return {} as T;
    }
  };

  set = <T>(key: string, data: T, expireSecond?: number): void => {
    const payload: StorageWrapper = {
      type: typeof data,
      data,
    };
    if (expireSecond) {
      payload.expire = Date.now() + expireSecond * 1000;
    }
    this.storage.setItem(key, JSON.stringify(payload));
  };

  remove = (key: string): void => {
    this.storage.removeItem(key);
  };

  clearAll = (): void => {
    this.storage.clear();
  };
}

// ✅ 正确命名
export const sessionStore = new AppStorage(sessionStorage);
export const localStore = new AppStorage(localStorage);
const defaultStore = new AppStorage(sessionStorage)
export default defaultStore;

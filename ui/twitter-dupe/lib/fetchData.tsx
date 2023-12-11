import { useState } from "react";

type FetchFunction<T> = (...args: any[]) => Promise<T>;

type FetchDataResult<T> = {
  data: T | null;
  loading: boolean;
  error: Error | null;
};

const useDataApi = <T,>(fetchFunction: FetchFunction<T>) => {
  const [result, setResult] = useState<FetchDataResult<T>>({
    data: null,
    loading: false,
    error: null,
  });

  const fetchData = async (...args: any[]) => {
    setResult({ data: null, loading: true, error: null });

    try {
      const tempResult = await fetchFunction(...args);
      setResult({
        data: tempResult,
        loading: false,
        error: null,
      });
    } catch (error: any) {
      setResult({
        data: null,
        loading: false,
        error: new Error(error.message),
      });
      return;
    }
  };

  return {
    data: result.data,
    loading: result.loading,
    error: result.error,
    fetchData,
  };
};

export default useDataApi;

import { useState, useEffect } from 'react';
interface QueryOptions<T> {
  variables?: T;
}

interface QueryResult<T, V> {
  refetch: (_?: QueryOptions<V>) => Promise<void>;
  data: T | null;
  error: Error | null;
  loading: boolean;
}

interface MutationResult<T, V> extends Omit<QueryResult<T, V>, 'refetch'> {
  run: (_?: QueryOptions<V>) => Promise<T>;
}

function useQuery<Q, V>(uri: string, query: string, { variables }: QueryOptions<V>): QueryResult<Q, V> {
  const [data, setData] = useState<Q | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [loading, setLoading] = useState<boolean >(true);

  const refetch = (options?: QueryOptions<V>) => {
    setLoading(true);
    setData(null);
    setError(null);

    const init = {
      method: 'POST',
      body: JSON.stringify({ query, variables: (options && options.variables) || variables || {} }),
      headers: { 'Content-Type': 'application/json' },
    };

    return fetch(uri, init)
      .then(async (res) => setData((await res.json()).data))
      .catch((err) => setError(err))
      .finally(() => setLoading(false));
  };

  useEffect(() => { refetch(); }, [query, variables]);

  return { error, data, loading, refetch };
}

function useMutation<M, V>(uri: string, query: string, { variables }: QueryOptions<V>): MutationResult<M, V> {
  const [data, setData] = useState<M | null>(null);
  const [error, setError] = useState<Error | null>(null);
  const [loading, setLoading] = useState<boolean >(true);
  const run = (options?: QueryOptions<V>) => {
    const init = {
      method: 'POST',
      body: JSON.stringify({ query, variables: (options && options.variables) || variables || {} }),
      headers: { 'Content-Type': 'application/json' },
    };

    return fetch(uri, init)
      .then(async (res) => {
        const data = (await res.json()).data;
        setData(data);
        return data;
      })
      .catch((err) => setError(err))
      .finally(() => setLoading(false));
  };
  return { error, data, loading, run }
}

export { useQuery, useMutation };

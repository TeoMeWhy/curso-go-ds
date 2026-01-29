# %%
import numpy as np

def gera_pacote(n, stop):
    return np.random.randint(1, stop+1, size=n)

def checa_pacote(album:set, pacote:np.array):
    album.update(pacote)

def completa_album(n, maximo):
    album = set()

    count = 0
    while len(album) < maximo:
        pacote = gera_pacote(n, maximo)
        checa_pacote(album, pacote)
        count +=1
    return count

# %%

n = 5
maximo = 550
N = 10000
values = [completa_album(n, maximo) for i in range(N)]
print(f"Em média, é necessário {sum(values) / len(values)} pacotes para completar o album")




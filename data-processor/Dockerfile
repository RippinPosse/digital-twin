### 1. Get Linux
FROM continuumio/miniconda3

ARG conda_env=myenv

ADD ./environment.yml /src/environment.yml
RUN conda env create -f /src/environment.yml

ENV PATH /opt/conda/envs/$conda_env/bin:$PATH
ENV CONDA_DEFAULT_ENV $conda_env

RUN conda init bash && \
    . ~/.bashrc && \
    conda activate $conda_env

RUN conda install hdfs3 -c conda-forge

COPY . .

ENV HADOOP_CONF_DIR=/etc/hadoop/conf
COPY ./hadoop/conf $HADOOP_CONF_DIR

COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt

CMD [ "python", "-u", "./main.py" ]
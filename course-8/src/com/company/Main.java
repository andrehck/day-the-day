package com.company;

public class Main {

    public static void main(String[] args) {
        String[] vect = new String[]{"teste","teste2"};

        for (int i=0; i<vect.length; i++){

            System.out.println(vect[i]);

        }

        for (String obj : vect) { // Tipo apelido : colecao

            System.out.println(obj);

        }
    }
}

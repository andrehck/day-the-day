package com.company;

import entities.Colaboradores;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {

        Scanner sc = new Scanner(System.in);
        List<Colaboradores> numberFunc = new ArrayList<>();

        Integer numberOfColab;
        Integer id;
        String nameColab;
        Double salaryColab;

        System.out.print("How many employees wiil be registered? ");
        numberOfColab = sc.nextInt();
        for (int n =0; n<numberOfColab;n++) {

            System.out.print("ID: ");
            id = sc.nextInt();

            while(hasID(numberFunc,id)){
                System.out.println("Id already taken! Try aganin");
                System.out.print("ID: ");
                id = sc.nextInt();
            }

            System.out.print("Name: ");
            nameColab = sc.next();

            System.out.print("Salary: ");
            salaryColab = sc.nextDouble();

            Colaboradores employers = new Colaboradores(id, nameColab, salaryColab);
            numberFunc.add(employers);
        }

        for (Colaboradores c : numberFunc){
            System.out.println(c);
        }

        sc.close();
    }

    public static boolean hasID(List<Colaboradores> list, int id){
        Colaboradores emp = list.stream().filter(x -> x.getId() == id).findFirst().orElse(null);
        return emp != null;
    }
}

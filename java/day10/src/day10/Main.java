package day10;

import java.io.FileReader;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.regex.Pattern;

import javafx.animation.KeyFrame;
import javafx.animation.Timeline;
import javafx.application.Application;
import javafx.scene.Scene;
import javafx.scene.layout.Pane;
import javafx.stage.Stage;
import javafx.util.Duration;

public class Main extends Application {

	public static String LIGHT_REGEX = "position=<(-\\d+|\\ \\d+), (-\\d+|\\ \\d+)> velocity=<(-\\d+|\\ \\d+), (-\\d+|\\ \\d+)>";
	
	public static void main(String[] args) {
		launch(args);
	}

	@Override
	public void start(Stage primaryStage) throws Exception {
		var root = new Pane();
		var scene = new Scene(root);
		primaryStage.setScene(scene);
		primaryStage.setTitle("SOS");
		primaryStage.setWidth(2000);
		primaryStage.setHeight(800);
		primaryStage.show();
		var in = new Scanner(new FileReader("input/day10.txt"));
		var lights = new ArrayList<Light>();
		Pattern lightPattern = Pattern.compile(LIGHT_REGEX);
		var maxX = 0;
		var maxY = 0;
		while (in.hasNextLine()) {
			var line = in.nextLine();
			var matcher = lightPattern.matcher(line);
			if (!matcher.find()) {
				System.out.println("BROKEN");
				primaryStage.close();
			}			
			var light = new Light(
					Integer.parseInt(matcher.group(1).replaceAll(" ", "")),
					Integer.parseInt(matcher.group(2).replaceAll(" ", "")),
					Integer.parseInt(matcher.group(3).replaceAll(" ", "")),
					Integer.parseInt(matcher.group(4).replaceAll(" ", "")));
			lights.add(light);
			if (Math.abs(light.getX()) > maxX) {
				maxX = Math.abs(light.getPropertyX());
			}
			if (Math.abs(light.getY()) > maxY) {
				maxY = Math.abs(light.getPropertyY());
			}
		}
		in.close();
		lights.forEach((l) -> l.setPerspective());
		root.getChildren().addAll(lights);
		
		// I'm sick and tired of waiting, just skip to the good part
		// NOTE: For part 2, I just increased 10000 until I got exactly on the image first try, which was 10345
		for (var i = 0; i < 10000; i++) {
			lights.forEach((l) -> l.travel());
		}
		
		var timeline = new Timeline(
			new KeyFrame(Duration.millis(20), e -> {
				lights.forEach((l) -> l.travel());				
			})
		);
		timeline.setCycleCount(345);
		timeline.play();
	}

}
